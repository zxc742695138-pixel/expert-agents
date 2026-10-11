# JavaScript/TypeScript Best Practices

Nguồn: [Airbnb JavaScript Style Guide](https://github.com/airbnb/javascript)
và Airbnb TypeScript addendum (được dùng rộng rãi trong công nghiệp, convention
rõ ràng có lý do kèm theo cho từng rule); [MDN Web Docs](https://developer.mozilla.org)
cho ngữ nghĩa ngôn ngữ chuẩn; [TypeScript Handbook](https://www.typescriptlang.org/docs/handbook/)
(Microsoft) cho type system.

## A. Khai báo biến & scope

- **Luôn dùng `const`, chỉ dùng `let` khi chắc chắn cần reassign, không bao
  giờ dùng `var`** (Airbnb JS — `no-var`) — `var` có function-scope và bị
  hoisting gây bug khó lường; `const`/`let` có block-scope rõ ràng.
- **Khai báo biến gần nơi dùng nhất, mỗi `const`/`let` một dòng khai báo
  riêng** — dễ đọc, dễ diff khi review.
- **Tránh biến global ngầm định** — mọi thứ nên được import/export rõ ràng
  qua module system (ESM `import`/`export`), không gán vào `window`/global
  scope trừ khi có lý do kỹ thuật rõ ràng (ví dụ polyfill).

## B. Hàm & arrow function

- **Dùng arrow function cho callback để giữ đúng `this`** (Airbnb JS) —
  tránh phải `.bind(this)` hoặc dùng biến tạm `self`/`that`.
- **Luôn dùng default parameter thay vì kiểm tra `undefined` thủ công trong
  body hàm** — `function f(x = 10)` rõ ràng hơn `if (x === undefined) x = 10`.
- **Không optimize sớm bằng cách lạm dụng tham chiếu/con trỏ thủ công** —
  JS engine hiện đại tối ưu tốt các pattern rõ ràng; code khó đọc để "nhanh
  hơn" một chút thường không đáng, trừ khi đã profile và chứng minh là
  bottleneck thật.

## C. Bất biến (Immutability) & cấu trúc dữ liệu

- **Ưu tiên không mutate object/array trực tiếp** — dùng spread
  (`{...obj}`, `[...arr]`), `Array.prototype.map/filter/reduce` thay vì
  `push`/`splice` trực tiếp lên state đang được tham chiếu nơi khác. Đặc biệt
  quan trọng trong React: mutate trực tiếp state có thể khiến component
  không re-render đúng (React so sánh reference, không deep-compare).
- **Dùng `Object.freeze`/`readonly` (TS) cho dữ liệu không nên đổi** — biến
  constant cấu hình, enum-like object.
- **Tránh deep nesting dữ liệu quá sâu** — object/state lồng >3 cấp khó
  update immutably và khó đọc; nên normalize (tách thành map theo id) khi dữ
  liệu quan hệ nhiều.

## D. Async/Promise

- **Ưu tiên `async`/`await` hơn `.then()` chain** khi logic tuần tự, trừ khi
  cần chạy song song (thì dùng `Promise.all`/`Promise.allSettled` rõ ràng) —
  dễ đọc hơn, stack trace lỗi rõ hơn.
- **Luôn xử lý rejected promise** — không để `async` function throw mà không
  có `try/catch` ở nơi gọi, hoặc không gắn `.catch()` — unhandled rejection
  có thể crash Node process hoặc để lỗi UI ở trạng thái "treo" không rõ
  nguyên nhân.
- **Không `await` tuần tự những việc độc lập nhau** — gây chậm không cần
  thiết; nếu 2 request không phụ thuộc nhau, chạy song song bằng
  `Promise.all`.
- **Race condition khi có nhiều request bất đồng bộ cập nhật cùng 1 state**
  (ví dụ search-as-you-type) — phải có cơ chế huỷ/bỏ qua response cũ (abort
  controller, hoặc so sánh request id) nếu không response chậm hơn có thể
  ghi đè lên response mới hơn.

## E. TypeScript — type safety

- **Tránh `any`** — mất hoàn toàn lợi ích của type checking; nếu type chưa
  rõ, dùng `unknown` và narrow trước khi dùng, thay vì `any` để tắt type
  checker.
- **Ưu tiên type suy luận (inference) khi rõ ràng, chỉ annotate khi cần làm
  rõ ý định hoặc tại ranh giới public API** (function signature, exported
  type) — annotate thừa mọi nơi làm code rối, không thêm giá trị.
- **Dùng union/discriminated union thay vì optional field tràn lan** — ví
  dụ trạng thái loading/success/error nên là 1 discriminated union
  (`{status: 'loading'} | {status: 'success', data: T} | {status: 'error',
  error: E}`) thay vì 1 object có `data?`, `error?`, `isLoading?` độc lập có
  thể rơi vào trạng thái vô nghĩa (vừa loading vừa có error).
- **Không dùng non-null assertion (`!`) để "im lặng" type checker** trừ khi
  có bằng chứng runtime chắc chắn (ví dụ ngay sau một check `if`) — lạm dụng
  `!` triệt tiêu lợi ích chính của TypeScript là bắt lỗi null/undefined lúc
  compile.

## F. So sánh & coercion

- **Luôn dùng `===`/`!==`, không dùng `==`/`!=`** (Airbnb JS) — tránh type
  coercion ngầm gây bug khó lường (`0 == false` là `true`, `'' == false` là
  `true`...).
- **Kiểm tra giá trị rõ ràng thay vì dựa vào truthy/falsy mơ hồ** khi giá
  trị có thể là `0`, `''`, `NaN` hợp lệ — `if (count === 0)` rõ ràng hơn
  `if (!count)` khi `0` là giá trị hợp lệ cần phân biệt với
  `null`/`undefined`.
