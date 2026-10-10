# Laws of UX

Nguồn: Jon Yablonski — [lawsofux.com](https://lawsofux.com) (nội dung gốc theo
license CC BY-NC-ND 4.0) và sách *Laws of UX: Using Psychology to Design
Better Products & Services*. Đây là tập hợp các nguyên lý tâm lý học hành vi/
nhận thức áp dụng vào thiết kế UX — không phải luật cứng, mà là "rule of
thumb" cần đối chiếu với user research thực tế.

| Law | Định nghĩa ngắn | Áp dụng khi review |
|---|---|---|
| **Jakob's Law** | Người dùng dành đa số thời gian trên các sản phẩm khác, nên họ mong sản phẩm của bạn hoạt động giống pattern quen thuộc đó. | Flag khi UI đi ngược convention phổ biến (vd: icon giỏ hàng không phải ở góc trên phải) mà không có lý do rõ ràng. |
| **Fitts's Law** | Thời gian để chạm tới một target phụ thuộc vào khoảng cách và kích thước của nó. | Target nhỏ/xa (đặc biệt trên mobile, vùng khó với ngón tay cái) → tăng lỗi, giảm tốc độ. |
| **Hick's Law** | Thời gian ra quyết định tăng theo số lượng và độ phức tạp của lựa chọn. | Menu/form quá nhiều option cùng lúc → nên nhóm, phân cấp, progressive disclosure. |
| **Miller's Law** | Trung bình con người chỉ giữ được 7 (±2) đơn vị thông tin trong working memory. | Danh sách dài không phân nhóm (chunking) → quá tải nhận thức. |
| **Postel's Law** | "Liberal trong cái nhận, conservative trong cái gửi" — hệ thống nên khoan dung với input không hoàn hảo của người dùng. | Form validation quá khắt khe (vd: bắt đúng format số điện thoại thủ công) vi phạm law này. |
| **Peak-End Rule** | Người dùng đánh giá một trải nghiệm chủ yếu dựa trên khoảnh khắc đỉnh điểm và khoảnh khắc kết thúc, không phải trung bình toàn bộ. | Trải nghiệm hoàn tất tác vụ (checkout success, onboarding cuối) nên được đầu tư, vì đó là "end" đọng lại. |
| **Aesthetic-Usability Effect** | Thiết kế đẹp thường được cảm nhận là dễ dùng hơn, dù usability thực tế không đổi. | Thẩm mỹ không chỉ là "làm đẹp" — nó ảnh hưởng trực tiếp đến cảm nhận usability, nhưng không thay thế được usability thật. |
| **Von Restorff Effect (Isolation Effect)** | Trong nhiều item giống nhau, item khác biệt sẽ được nhớ/chú ý nhiều hơn. | Dùng để làm nổi bật CTA chính; nhưng lạm dụng (quá nhiều thứ "khác biệt") sẽ mất hiệu quả. |
| **Tesler's Law (Conservation of Complexity)** | Mỗi hệ thống có một độ phức tạp cố hữu không thể loại bỏ, chỉ có thể dịch chuyển (giữa hệ thống và người dùng). | Đơn giản hóa quá mức cho UI đôi khi đẩy độ phức tạp sang phía người dùng phải tự suy luận — cần cân bằng. |
| **Doherty Threshold** | Năng suất tăng vượt bậc khi hệ thống phản hồi trong < 400ms. | Loading > 400ms không có feedback (skeleton, spinner) sẽ cảm giác "chậm" dù thực tế backend nhanh. |
| **Parkinson's Law** | Công việc sẽ giãn ra để lấp đầy thời gian được cấp cho nó. | Form/flow không có ràng buộc thời gian/bước rõ ràng khiến người dùng trì hoãn hoàn thành. |
| **Zeigarnik Effect** | Người dùng nhớ các tác vụ *chưa hoàn thành* tốt hơn tác vụ đã hoàn thành. | Progress bar, checklist onboarding (vd: "70% hoàn thành hồ sơ") tận dụng hiệu ứng này để kéo người dùng quay lại. |
| **Serial Position Effect** | Người dùng nhớ item đầu và cuối trong một danh sách tốt hơn các item giữa. | Đặt action/thông tin quan trọng nhất ở đầu hoặc cuối danh sách/menu, không ở giữa. |
| **Occam's Razor** | Với cùng kết quả, giải pháp đơn giản nhất nên được ưu tiên. | Khi 2 cách thiết kế cùng giải quyết vấn đề, chọn cách ít thành phần/bước hơn. |
| **Pareto Principle (80/20)** | ~80% hiệu ứng đến từ ~20% nguyên nhân. | 20% tính năng/màn hình được dùng nhiều nhất nên được tối ưu kỹ nhất, không rải đều effort. |
| **Goal-Gradient Effect** | Xu hướng hành động tăng tốc khi càng gần đạt mục tiêu. | Progress indicator nên cho thấy rõ "gần xong rồi" để duy trì động lực hoàn tất flow. |
| **Flow** (Csikszentmihalyi) | Trạng thái tập trung cao độ khi thử thách và kỹ năng cân bằng. | UI gây gián đoạn liên tục (popup, thông báo không cần thiết) phá vỡ flow của người dùng. |
| **Cognitive Load** | Tổng lượng "tài nguyên tư duy" cần để hoàn thành một tác vụ. | Mục tiêu thiết kế là giảm cognitive load không cần thiết (extraneous load), không phải giảm nội dung cốt lõi. |
| **Chunking** | Nhóm thông tin thành các cụm nhỏ để dễ xử lý/nhớ hơn. | Số điện thoại/thẻ tín dụng nên hiển thị theo nhóm số, không phải một dãy dài liền. |
| **Selective Attention** | Người dùng lọc bỏ thông tin không liên quan đến mục tiêu hiện tại của họ (liên quan "banner blindness"). | Thông tin quảng cáo/khuyến mãi chen giữa flow tác vụ chính dễ bị bỏ qua hoàn toàn hoặc gây khó chịu. |
| **Law of Proximity** *(Gestalt)* | Các object gần nhau được coi là có liên quan. | Xem chi tiết nhóm Gestalt đầy đủ ở `03-visual-design-systems.md`. |
| **Law of Prägnanz** *(Gestalt)* | Não có xu hướng diễn giải hình ảnh phức tạp theo cách đơn giản nhất có thể. | — |

> Ghi chú: các "Law of Similarity/Common Region/Uniform Connectedness" thuộc
> nhóm Gestalt, được trình bày đầy đủ cùng các nguyên tắc visual design khác
> trong `03-visual-design-systems.md` để tránh trùng lặp.
