# Usability Heuristics

Nguồn: Jakob Nielsen / Nielsen Norman Group (nngroup.com) — *10 Usability
Heuristics for User Interface Design* (1994, dựa trên phân tích 249 usability
problem); Don Norman — *The Design of Everyday Things*; Steve Krug — *Don't
Make Me Think*.

## A. 10 Usability Heuristics (Jakob Nielsen / NN/g)

1. **Visibility of system status** — Hệ thống luôn cho người dùng biết đang
   xảy ra chuyện gì (loading state, progress indicator, trạng thái nút sau khi
   bấm). Vi phạm: bấm submit mà không có loading/feedback nào.
2. **Match between system and the real world** — Dùng ngôn ngữ, khái niệm,
   icon quen thuộc với người dùng thật, không dùng thuật ngữ nội bộ/hệ thống.
3. **User control and freedom** — Luôn có "exit", "undo", "cancel" rõ ràng.
   Vi phạm: flow một chiều không cho back, không cho huỷ hành động vừa làm.
4. **Consistency and standards** — Cùng một hành động/khái niệm phải được thể
   hiện giống nhau xuyên suốt app; tuân theo convention nền tảng (xem
   `04-platform-guidelines.md`).
5. **Error prevention** — Thiết kế ngăn lỗi xảy ra trước khi cần thông báo lỗi
   (confirm dialog cho hành động phá hủy, disable nút khi input chưa hợp lệ).
6. **Recognition rather than recall** — Không bắt người dùng phải nhớ thông
   tin giữa các màn hình; hiển thị lại option, lịch sử, gợi ý.
7. **Flexibility and efficiency of use** — Có đường tắt cho power user
   (shortcut, bulk action) nhưng vẫn giữ đơn giản cho người mới.
8. **Aesthetic and minimalist design** — Chỉ hiển thị thông tin liên quan;
   mỗi thông tin dư thừa làm loãng thông tin quan trọng (liên hệ trực tiếp với
   *Signal-to-Noise Ratio* trong `03-visual-design-systems.md`).
9. **Help users recognize, diagnose, and recover from errors** — Thông báo lỗi
   bằng ngôn ngữ thường (không mã lỗi khô), chỉ rõ nguyên nhân và cách khắc
   phục.
10. **Help and documentation** — Dù lý tưởng là không cần, nhưng nếu có, phải
    dễ tìm, dễ search, theo đúng context của tác vụ đang làm.

## B. Mô hình nhận thức của Don Norman (*The Design of Everyday Things*)

- **Affordance** — Khả năng một đối tượng "gợi ý" hành động có thể thực hiện
  (một nút nổi khối gợi ý "bấm được").
- **Signifier** — Dấu hiệu tường minh chỉ ra affordance ở đâu, dùng thế nào
  (label, icon, đường viền, đổi màu khi hover).
- **Mapping** — Mối liên hệ giữa control và kết quả phải tự nhiên (ví dụ:
  slider kéo phải = tăng giá trị).
- **Feedback** — Mọi hành động phải có phản hồi tức thời và rõ ràng.
- **Conceptual model** — Người dùng tự xây mô hình tinh thần về cách hệ thống
  hoạt động; UI phải khớp với mô hình đó, không đi ngược trực giác.
- **Constraints** — 4 loại giới hạn giúp ngăn hành động sai: physical
  (nút bị disable không bấm được), cultural (màu đỏ = nguy hiểm/stop),
  semantic (ý nghĩa ngữ cảnh), logical (suy luận hợp lý từ layout).
- **Discoverability** — Tính năng phải được tìm thấy một cách tự nhiên, không
  ẩn sâu không gợi ý.
- **Slip vs. Mistake** — Slip: người dùng biết mục tiêu nhưng tay/thao tác sai
  (bấm nhầm nút) → fix bằng spacing, confirm. Mistake: người dùng hiểu sai mục
  tiêu từ đầu → fix bằng làm rõ thông tin/label.

## C. Nguyên tắc của Steve Krug (*Don't Make Me Think*)

- **Self-evident, không cần giải thích** — Người dùng nhìn vào phải hiểu ngay
  cái gì bấm được, cái gì không, không cần suy nghĩ.
- **Người dùng scan, không đọc** — Thiết kế cho hành vi quét trang (scanning):
  heading rõ, đoạn ngắn, bullet, highlight từ khóa — không phải văn bản dài
  liền mạch.
- **Omit needless words** — Cắt bớt chữ không cần thiết; mỗi chữ thêm vào là
  thêm thứ phải đọc/xử lý.
- **"Trunk test"** — Nhìn vào một màn hình bất kỳ (dù lấy ngẫu nhiên, không có
  context), người dùng phải tự trả lời được: Đây là gì? Tôi có thể làm gì ở
  đây? Vì sao tôi nên quan tâm?
