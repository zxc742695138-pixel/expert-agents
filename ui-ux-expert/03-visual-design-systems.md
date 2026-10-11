# Visual Design & Design Systems

Nguồn: Tâm lý học Gestalt (Max Wertheimer, Kurt Koffka, Wolfgang Köhler — đầu
thế kỷ 20, được áp dụng rộng rãi vào UI design); Adam Wathan & Steve Schoger —
*Refactoring UI* (sách, blog, và các nguyên tắc công khai mà hai tác giả chia
sẻ qua Tailwind CSS — một design system do chính họ xây dựng dựa trên các
nguyên tắc này).

## A. Gestalt Principles

- **Proximity** — Các phần tử ở gần nhau được nhận thức là có liên quan với
  nhau. Spacing là công cụ phân nhóm thông tin, không chỉ là "khoảng trống".
- **Similarity** — Các phần tử giống nhau về hình dạng/màu/kích thước được
  coi là cùng nhóm/cùng chức năng. Vi phạm: 2 nút cùng hành động nhưng style
  khác nhau → gây hiểu nhầm là 2 hành động khác nhau.
- **Closure** — Não tự "lấp đầy" phần thiếu để nhận ra một hình hoàn chỉnh.
  Dùng trong icon tối giản, loading skeleton.
- **Continuity** — Mắt có xu hướng theo một đường/hướng liên tục. Dùng để dẫn
  hướng mắt người dùng qua flow (alignment theo một trục).
- **Figure/Ground** — Người dùng phân biệt vật thể chính (figure) với nền
  (ground). Contrast không đủ giữa content và background làm mất phân biệt
  này (liên quan trực tiếp đến WCAG contrast — xem `05-accessibility-wcag.md`).
- **Common Fate / Common Region** — Các phần tử di chuyển cùng nhau hoặc nằm
  trong cùng một đường viền/khung được coi là một nhóm (card, container).
- **Prägnanz** — Não diễn giải hình ảnh phức tạp theo cách đơn giản nhất có
  thể → ưu tiên hình dạng/layout đơn giản, dễ nhận diện.

## B. Nguyên tắc từ Refactoring UI (Wathan & Schoger)

### Visual hierarchy
- Không phải mọi thứ đều quan trọng ngang nhau — chủ động **de-emphasize**
  (giảm trọng số) các phần tử phụ (vd: label phụ, timestamp) bằng cách giảm
  size/weight/contrast, thay vì chỉ tăng trọng số phần tử chính.
- Dùng **kết hợp** size + weight + color + contrast để tạo hierarchy, không
  chỉ dựa vào 1 thuộc tính (vd: chỉ phóng to chữ mà không đổi weight/color sẽ
  cho hierarchy yếu).
- **Labels là giải pháp cuối cùng** — nếu có thể truyền đạt ý nghĩa qua
  formatting/vị trí/icon, không cần thêm label chữ.

### Layout & spacing
- Dùng **hệ thống spacing nhất quán** dựa trên một scale cố định (phổ biến:
  bội số của 4px hoặc 8px — "8-point grid system"), tránh giá trị spacing tùy
  tiện (7px, 13px, 22px...).
- **Whitespace không phải không gian lãng phí** — whitespace chủ động là
  công cụ tạo nhóm (Proximity) và giảm cognitive load.
- Không phải mọi phân chia cần border — ưu tiên dùng spacing hoặc nền màu
  khác biệt trước khi thêm border; lạm dụng border làm UI rối.

### Typography
- Dùng **type scale** cố định (modular scale), không tự do chọn font-size
  tùy hứng. Scale phổ biến (dựa trên hệ spacing tương tự Tailwind):
  `12 / 14 / 16 / 18 / 20 / 24 / 30 / 36 / 48 / 60 / 72px`.
- Giới hạn số lượng font-size và font-weight dùng trong một màn hình để giữ
  tính nhất quán (constrain the design space).
- Dùng font được thiết kế cho UI (có nhiều weight, số liệu rõ ràng, dễ đọc ở
  size nhỏ: vd Inter, SF Pro, Roboto...).
- Line-height và line-length (độ dài dòng văn bản) ảnh hưởng trực tiếp đến
  khả năng đọc — dòng quá dài (> ~75-80 ký tự) hoặc line-height quá chật đều
  giảm readability.

### Color
- Xây **bảng màu có scale tint/shade** (thường 50→900, giống Tailwind) cho
  mỗi màu chính, không chỉ 1 giá trị hex cố định — giúp có đủ sắc độ cho các
  trạng thái (hover, disabled, background nhẹ...).
- Áp dụng tỉ lệ phân bổ màu kiểu **60-30-10** (màu chủ đạo – màu phụ – màu
  nhấn) để tránh UI "loạn màu" không có điểm nhấn rõ ràng.
- Không dùng chữ màu xám nhạt trên nền có màu (dễ vi phạm contrast ratio —
  xem WCAG ở file 05).

### Depth & finishing touches
- Dùng **shadow để biểu đạt elevation** một cách nhất quán (elevation thấp =
  shadow nhẹ, elevation cao = shadow rõ hơn), theo một hệ thống chứ không
  random giữa các component.
- Có thể dùng **overlap** (phần tử chồng lên nhau nhẹ) để tạo cảm giác
  depth/liên kết giữa các block.
- Border-radius nên nhất quán theo 1-2 giá trị trong toàn bộ hệ thống, không
  trộn nhiều mức bo góc khác nhau giữa các component cùng cấp.
