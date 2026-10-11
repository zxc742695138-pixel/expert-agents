# Clean Code Principles

Nguồn: Robert C. Martin ("Uncle Bob") — *Clean Code: A Handbook of Agile
Software Craftsmanship* (2008); nguyên tắc SOLID gốc do Robert C. Martin đặt
tên/hệ thống hoá dựa trên các khái niệm OOP cổ điển; *Design Patterns:
Elements of Reusable Object-Oriented Software* (Gang of Four — Gamma, Helm,
Johnson, Vlissides, 1994) cho nền tảng pattern. Các nguyên tắc này viết cho
OOP nói chung nhưng áp dụng trực tiếp vào JS/TS/React/Vue component code.

## A. Đặt tên (Naming)

- **Tên phải nói lên ý định** — `elapsedTimeInDays` tốt hơn `d`. Tên biến,
  hàm, class phải trả lời được "nó làm gì/chứa gì" mà không cần đọc comment.
- **Tránh thông tin gây hiểu lầm** — không gọi một biến `accountList` nếu nó
  không phải kiểu List thực sự (dùng `accounts` hay `accountGroup`).
- **Phân biệt rõ ràng, có ý nghĩa** — tránh `data1`, `data2`; tránh các hậu
  tố vô nghĩa như `ProductInfo` vs `ProductData` nếu không có khác biệt thực.
- **Tên phát âm được, tìm kiếm được** — biến 1 ký tự (`e`) khó search trong
  codebase lớn; ưu tiên tên đủ dài để `grep` ra đúng chỗ cần.
- **Tên hàm là động từ/cụm động từ, tên class là danh từ/cụm danh từ** —
  `calculateTotal()`, `UserRepository`, không ngược lại.

## B. Hàm (Functions)

- **Hàm nên nhỏ, và nhỏ hơn thế nữa** — một hàm lý tưởng làm đúng 1 việc.
  Dấu hiệu hàm làm nhiều việc: tên hàm phải dùng "và" để mô tả
  (`validateAndSave`), hoặc có nhiều đoạn code phân tách bằng comment mô tả
  từng "bước".
- **Một mức trừu tượng cho mỗi hàm** — không trộn logic cấp cao (business
  rule) với chi tiết cấp thấp (string formatting, DOM manipulation) trong
  cùng 1 hàm.
- **Số lượng tham số càng ít càng tốt** — 0-2 tham số là lý tưởng; ≥4 tham số
  là dấu hiệu nên gom thành 1 object tham số (áp dụng trực tiếp vào props
  object trong React/Vue thay vì liệt kê nhiều positional args).
- **Không có side-effect ẩn** — tên hàm phải phản ánh đúng mọi effect nó gây
  ra; hàm tên `checkPassword()` mà bên trong còn khởi tạo session là side
  effect ẩn, gây bug khó debug.
- **Command Query Separation** — một hàm nên hoặc làm gì đó (command) hoặc
  trả lời gì đó (query), không nên vừa thay đổi state vừa trả giá trị — dễ
  gây nhầm lẫn khi đọc code gọi hàm đó.

## C. Nguyên tắc SOLID (áp dụng cho module/component/class)

- **S — Single Responsibility Principle**: một class/component chỉ nên có
  1 lý do để thay đổi. Component React vừa fetch data, vừa format, vừa
  render UI phức tạp → nên tách custom hook (fetch/format) khỏi component
  (render).
- **O — Open/Closed Principle**: mở để extend, đóng để modify — thiết kế để
  thêm behavior mới không cần sửa code cũ (ví dụ dùng strategy pattern/props
  injection thay vì if/else dài thêm case mới mỗi lần có yêu cầu).
- **L — Liskov Substitution Principle**: một subtype phải dùng thay được
  cho base type mà không phá vỡ hành vi mong đợi — áp dụng khi extend/override
  component hoặc class.
- **I — Interface Segregation Principle**: không ép một module phụ thuộc
  vào interface nó không dùng — tránh prop interface "god object" nhận 20
  props mà component con chỉ dùng 3.
- **D — Dependency Inversion Principle**: module cấp cao không nên phụ thuộc
  trực tiếp vào module cấp thấp cụ thể — phụ thuộc vào abstraction (ví dụ:
  component gọi qua một service interface/hook, không gọi trực tiếp
  `fetch()` với URL hardcode rải rác khắp codebase).

## D. DRY, KISS, YAGNI

- **DRY (Don't Repeat Yourself)** — tri thức nghiệp vụ chỉ nên có 1 nơi biểu
  diễn. Lưu ý: 2 đoạn code *trông giống* nhau nhưng biểu diễn 2 tri thức
  nghiệp vụ *khác nhau* thì KHÔNG cần trừu tượng hoá chung — trừu tượng hoá
  sai chỗ (DRY quá mức) tạo ra coupling giả, khó sửa hơn là duplicate.
- **KISS (Keep It Simple, Stupid)** — chọn giải pháp đơn giản nhất giải
  quyết đúng vấn đề hiện tại; không thêm abstraction/pattern phức tạp cho
  một vấn đề chưa thực sự phức tạp.
- **YAGNI (You Aren't Gonna Need It)** — không xây tính năng/tham số
  "phòng khi sau này cần" mà hiện tại chưa có yêu cầu cụ thể — tăng diện
  tích bug và bảo trì cho giá trị chưa chắc xảy ra.

## E. Comment

- **Comment tốt nhất là không cần comment** — code tự giải thích qua tên
  biến/hàm rõ ràng luôn tốt hơn comment giải thích code khó hiểu.
- **Chỉ comment khi giải thích "tại sao", không giải thích "làm gì"** — code
  đã nói "làm gì"; comment nên giải thích constraint ẩn, lý do chọn giải pháp
  không hiển nhiên, hoặc cảnh báo tác dụng phụ không rõ từ code.
- **Comment lỗi thời là tệ hơn không có comment** — comment mô tả sai hành vi
  hiện tại của code (do code đã đổi mà comment không update) gây hiểu nhầm
  nguy hiểm hơn việc không có thông tin.
