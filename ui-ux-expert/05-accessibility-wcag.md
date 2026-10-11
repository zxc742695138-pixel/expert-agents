# Accessibility — WCAG 2.2

Nguồn: W3C Web Accessibility Initiative — [WCAG 2.2](https://www.w3.org/TR/WCAG22/)
(hoàn thiện tháng 10/2023). Đây là chuẩn bắt buộc tối thiểu (mức AA) cho bất
kỳ sản phẩm hướng tới người dùng đại chúng.

## 4 nguyên tắc nền (POUR)

1. **Perceivable** — Người dùng phải cảm nhận được nội dung (nhìn, nghe, hoặc
   qua assistive tech).
2. **Operable** — Người dùng phải điều khiển được UI (bàn phím, con trỏ, hỗ
   trợ thao tác).
3. **Understandable** — Nội dung và cách vận hành phải dễ hiểu, dễ đoán.
4. **Robust** — Nội dung phải tương thích với nhiều loại user agent/assistive
   technology khác nhau (semantic HTML, ARIA đúng chuẩn).

## Các success criteria cụ thể cần kiểm tra khi review UI

| Criterion | Mức | Yêu cầu | Khi review design cần check |
|---|---|---|---|
| **1.4.3 Contrast (Minimum)** | AA | Text thường ≥ **4.5:1**; text lớn (≥18pt bold hoặc ≥24pt) ≥ **3:1** | Dùng công cụ đo contrast giữa text và background thực tế, không nhìn mắt áng chừng. |
| **1.4.11 Non-text Contrast** | AA | Icon, border của component tương tác, biểu đồ ≥ **3:1** so với nền liền kề | Icon/border quá nhạt (xám nhạt trên trắng) thường vi phạm. |
| **1.4.1 Use of Color** | A | Không dùng màu là **phương tiện duy nhất** truyền đạt thông tin | Trạng thái lỗi/thành công chỉ khác màu (không kèm icon/text) → vi phạm. |
| **1.4.4 Resize Text** | AA | Text phải resize được đến 200% mà không mất nội dung/chức năng | Layout cứng (fixed height che chữ) khi scale sẽ vi phạm. |
| **2.5.8 Target Size (Minimum)** | AA *(mới trong 2.2)* | Pointer target tối thiểu **24×24 CSS px** (trừ các trường hợp ngoại lệ: inline, equivalent control khác đã đủ lớn, essential) | So khớp với chuẩn 44pt (Apple)/48dp (Material) ở file 04 — đó là best practice cao hơn mức sàn WCAG. |
| **2.4.11 Focus Not Obscured (Minimum)** | AA *(mới trong 2.2)* | Phần tử đang có keyboard focus không được bị che hoàn toàn bởi content khác (sticky header, modal...) | Kiểm tra khi tab qua các phần tử có sticky elements trên màn hình. |
| **2.4.13 Focus Appearance** | AAA | Focus indicator rõ ràng, contrast ≥ 3:1 so với màu xung quanh | Khuyến nghị áp dụng dù là mức AAA — focus state mờ/không rõ là lỗi phổ biến. |
| **3.2.6 Consistent Help** | A *(mới trong 2.2)* | Cơ chế hỗ trợ (chat, hotline, FAQ) phải ở vị trí nhất quán giữa các trang | — |
| **3.3.7 Redundant Entry** | A *(mới trong 2.2)* | Không bắt người dùng nhập lại thông tin đã cung cấp trong cùng flow | Form nhiều bước yêu cầu nhập lại email/địa chỉ đã nhập trước đó → vi phạm. |
| **3.3.8 Accessible Authentication (Minimum)** | AA *(mới trong 2.2)* | Không bắt buộc "cognitive function test" (nhớ password phức tạp, giải captcha khó) làm bước xác thực duy nhất, không có phương án thay thế | Captcha khó/không có alternative (audio, thiết bị hỗ trợ) → vi phạm. |
| **2.3.3 Animation from Interactions** | AAA | Tôn trọng `prefers-reduced-motion`; animation do tương tác không gây khó chịu/chóng mặt | Áp dụng cho parallax, auto-play motion lớn. |

## Cách dùng khi review

- Baseline bắt buộc để coi một design là "pass accessibility" tối thiểu:
  tất cả criterion mức **A và AA** trong bảng trên.
- Criterion mức **AAA** (2.4.13, 2.3.3) nên được đề xuất như "nên có" (best
  practice), không chấm là "lỗi nghiêm trọng" nếu thiếu — nhưng vẫn nên ghi
  nhận trong report.
- Khi không thể đo chính xác contrast từ ảnh tĩnh, ghi rõ trong report là
  "cần xác minh bằng công cụ đo màu thực tế" — không khẳng định pass/fail
  một cách chắc chắn nếu không có giá trị màu chính xác.
