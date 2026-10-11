# Framework Patterns & Anti-patterns

Nguồn: [React docs](https://react.dev) chính thức (Rules of Hooks, Thinking
in React, You Might Not Need an Effect); [Vue.js Style Guide](https://vuejs.org/style-guide/)
chính thức; [Angular Style Guide](https://angular.dev/style-guide) chính
thức. Mỗi framework có convention riêng — không áp quy tắc của framework
này cho framework khác.

> Quy ước riêng của codebase đang review (CLAUDE.md, `.eslintrc`, style
> guide nội bộ) luôn được tôn trọng trước style guide chung ở đây nếu có
> xung đột — xem `README.md`, nguyên tắc 4.

## A. React

### Rules of Hooks (react.dev)
- **Chỉ gọi Hook ở top-level của function component hoặc custom hook** —
  không gọi trong loop, condition, hay nested function. Vi phạm này phá vỡ
  thứ tự Hook giữa các lần render, gây bug state sai lệch rất khó debug.
- **Chỉ gọi Hook từ component React hoặc custom hook (tên bắt đầu `use`)** —
  không gọi từ hàm JS thường.
- **`useEffect` không phải nơi để đồng bộ mọi thứ** (react.dev — *You Might
  Not Need an Effect*) — nếu một giá trị có thể tính trực tiếp từ props/state
  hiện có trong lúc render, tính trực tiếp, không đưa vào `useEffect` +
  `useState` riêng (gây render thừa + dễ bug đồng bộ sai thời điểm).
- **Dependency array của `useEffect`/`useMemo`/`useCallback` phải đầy đủ và
  chính xác** — thiếu dependency gây stale closure (dùng giá trị cũ); dư
  dependency không cần thiết gây effect chạy lại quá nhiều.

### State & re-render
- **Không mutate state trực tiếp** — luôn tạo object/array mới khi update
  (xem file `02`, mục C) — React dùng reference equality để quyết định
  re-render.
- **Đưa state xuống đúng component cần nó (state colocation)** — state đặt
  ở component cha quá cao khi chỉ 1 component con cần, gây re-render lan ra
  không cần thiết toàn bộ cây con.
- **Derived state nên tính trong lúc render, không duplicate vào state
  riêng** — state trùng lặp từ 1 nguồn dữ liệu dễ lệch nhau (một nơi update,
  nơi khác quên).
- **Key trong list phải ổn định và duy nhất, không dùng index khi list có
  thể reorder/filter/insert** — dùng index làm key khi list thay đổi thứ tự
  gây React nhận diện sai item, dẫn đến mất state cục bộ hoặc animation sai.

### Component design
- **Tách component theo trách nhiệm (Single Responsibility — xem file
  `01`)** — component vừa fetch data vừa chứa logic nghiệp vụ phức tạp vừa
  render UI chi tiết nên tách thành custom hook (logic) + component thuần
  (render).
- **Tránh prop drilling quá sâu (>3 cấp)** — xem xét Context hoặc state
  management phù hợp, nhưng không dùng Context cho state thay đổi tần suất
  cao toàn cây (gây re-render toàn bộ consumer).

## B. Vue

- **Luôn dùng `key` với `v-for`**, không dùng index khi list có thể thay
  đổi thứ tự (vuejs.org Style Guide — tương tự lý do ở React).
- **Không dùng `v-if` và `v-for` trên cùng 1 element** — thứ tự ưu tiên giữa
  2 directive gây nhầm lẫn và hiệu năng kém hơn tách ra `computed` filter
  trước, hoặc bọc bằng `<template>`.
- **Props nên có validation rõ ràng (type, required, default)** — giúp lỗi
  hiện ra sớm (dev mode warning) khi component cha truyền sai kiểu dữ liệu.
- **Computed property cho giá trị suy ra từ state, không dùng method gọi lại
  mỗi lần render cho cùng 1 phép tính** — `computed` có cache, method thì
  không.

## C. Angular

- **Tuân theo cấu trúc file/folder theo feature module** (angular.dev Style
  Guide) — không tổ chức phẳng toàn bộ component/service trong 1 folder lớn.
- **OnPush change detection cho component không cần theo dõi mutation sâu**
  — giảm số lần Angular check toàn bộ cây component, cải thiện hiệu năng khi
  app lớn.
- **Unsubscribe Observable đúng lifecycle** (`ngOnDestroy`, hoặc dùng
  `takeUntilDestroyed`/`async` pipe) — quên unsubscribe là nguyên nhân memory
  leak phổ biến nhất trong Angular.

## D. Anti-pattern chung (đa framework)

- **"God component"** — 1 component làm quá nhiều việc (fetch + validate +
  format + render + xử lý nhiều sự kiện phức tạp) → khó test, khó tái sử
  dụng, khó review diff.
- **Logic nghiệp vụ nhúng trực tiếp trong JSX/template** — điều kiện phức
  tạp, tính toán nhiều bước ngay trong JSX làm giảm readability; nên trích
  ra biến/hàm có tên rõ nghĩa trước khi render.
- **Inline function/object literal tạo mới mỗi lần render truyền xuống
  component con đã được memo hoá** (`React.memo`, Vue `v-memo`) — vô hiệu
  hoá memoization vì reference luôn khác nhau giữa các lần render.
