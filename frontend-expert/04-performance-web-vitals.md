# Performance & Core Web Vitals

Nguồn: [web.dev](https://web.dev) (Google) — Core Web Vitals, tài liệu chính
thức về performance web; Addy Osmani — các bài viết/sách về bundle
optimization, *Essential Image Optimization*; MDN Web Docs cho API
performance (IntersectionObserver, requestIdleCallback...).

## A. Core Web Vitals (web.dev)

3 chỉ số chính Google dùng để đo chất lượng trải nghiệm thực tế:

- **LCP (Largest Contentful Paint)** — thời gian render xong phần tử lớn
  nhất trong viewport. Ngưỡng tốt: **≤ 2.5s**. Nguyên nhân LCP chậm thường
  gặp trong code: ảnh/video lớn không tối ưu, font chặn render (không có
  `font-display`), JS chặn main thread trước khi render nội dung chính,
  không preload resource quan trọng nhất.
- **INP (Interaction to Next Paint)** — độ trễ từ lúc người dùng tương tác
  (click, tap, gõ phím) đến khi UI phản hồi lại trên màn hình. Ngưỡng tốt:
  **≤ 200ms**. Nguyên nhân thường gặp: handler sự kiện chạy tác vụ nặng
  đồng bộ trên main thread (tính toán lớn, parse JSON lớn), re-render toàn
  bộ cây component không cần thiết sau mỗi tương tác nhỏ.
- **CLS (Cumulative Layout Shift)** — tổng độ "nhảy" bố cục ngoài ý muốn
  trong lúc trang tải. Ngưỡng tốt: **≤ 0.1**. Nguyên nhân thường gặp: ảnh/
  iframe không set `width`/`height` (hoặc `aspect-ratio`) trước khi load
  xong, font thay thế (fallback → web font) làm đổi kích thước text, nội
  dung/banner được chèn động vào phía trên nội dung đã render.

## B. Bundle size & code splitting

- **Code splitting theo route/theo tính năng lớn** — không bundle toàn bộ
  app vào 1 file JS; dùng dynamic `import()` cho phần code không cần ngay
  khi tải trang (modal ít dùng, trang phụ, library nặng chỉ 1 tính năng cần).
- **Tree-shaking chỉ hoạt động với ESM, không hoạt động tốt với CommonJS**
  — import named export cụ thể (`import { debounce } from 'lodash-es'`)
  thay vì import cả thư viện (`import _ from 'lodash'`) để bundler loại bỏ
  code không dùng.
- **Kiểm tra kích thước dependency mới trước khi thêm** — một thư viện nhỏ
  chức năng tương đương có thể thay thế một thư viện lớn chỉ dùng 1 hàm nhỏ
  trong đó.
- **Lazy-load thứ không cần ngay khi First Paint** — ảnh dưới fold
  (`loading="lazy"` hoặc IntersectionObserver), component không hiển thị
  ngay (modal, tab ẩn).

## C. Rendering performance

- **Tránh forced synchronous layout (layout thrashing)** — đọc property gây
  reflow (`offsetHeight`, `getBoundingClientRect`...) ngay sau khi ghi style
  trong 1 loop buộc browser phải tính layout lại nhiều lần liên tiếp; nên
  đọc hết trước, ghi hết sau (batch read/write riêng).
- **Debounce/throttle các handler chạy liên tục** (scroll, resize, input
  gõ nhanh) — tránh gọi hàm nặng (network request, tính toán lớn) ở tần số
  cao hơn cần thiết.
- **Virtualize list dài** (react-window, vue-virtual-scroller hoặc tương
  đương) — render toàn bộ DOM node cho list hàng nghìn item ngay cả khi chỉ
  hiển thị 10-20 item trong viewport là lãng phí rendering lớn.
- **`requestAnimationFrame` cho animation điều khiển bằng JS**, không dùng
  `setTimeout`/`setInterval` — đồng bộ với refresh rate của browser, tránh
  giật/lag.

## D. Network & asset loading

- **Preload resource quan trọng nhất cho LCP** (`<link rel="preload">` cho
  hero image/font chính) — không để browser tự phát hiện muộn.
- **Chọn đúng format ảnh** — AVIF/WebP cho ảnh thường (nhẹ hơn JPEG/PNG đáng
  kể ở cùng chất lượng), SVG cho icon/illustration vector.
- **`font-display: swap` (hoặc `optional`)** để tránh chữ "vô hình" (FOIT)
  trong lúc web font tải — đổi lấy khả năng xảy ra CLS nhẹ khi font đổi,
  nhưng vẫn tốt hơn chặn hiển thị nội dung hoàn toàn.

## Cách dùng khi review

- Không đoán hiệu năng từ "nhìn code có vẻ nặng" — ưu tiên chỉ ra pattern cụ
  thể đã biết gây vấn đề (ví dụ: liệt kê đúng dòng code gây re-render thừa,
  đúng chỗ thiếu `width`/`height` trên ảnh) theo đúng nguyên tắc "trích
  nguồn khi nhận xét" ở README.
- Nếu không đủ dữ liệu để khẳng định tác động thực tế (ví dụ cần đo bằng
  Lighthouse/profiler thật), nói rõ "cần đo thực tế để xác nhận mức độ ảnh
  hưởng" thay vì khẳng định chắc chắn mức độ nghiêm trọng.
