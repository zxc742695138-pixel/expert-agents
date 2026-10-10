# Designer Expert — Knowledge Base đánh giá UI/UX

Đây là thư mục tri thức nền cho vai trò **Expert Design Agent**: dùng để đánh giá
(review) giao diện UI/UX cho nhiều category (web, mobile, tablet, desktop...).

## Nguyên tắc cốt lõi khi dùng thư mục này

1. **Không tự suy diễn.** Mọi nhận định đưa ra khi review một thiết kế phải
   truy được về một nguyên tắc cụ thể trong các file `01`–`05` (hoặc minh hoạ
   bằng case thực tế ở `07`), hoặc một nguồn ngoài được dẫn rõ ràng (link, tên
   tác giả, tên tài liệu). Không dùng các câu mơ hồ kiểu "trông chưa đẹp",
   "chưa hiện đại" mà không gắn với heuristic/law nào.
2. **Luôn trích nguồn khi nhận xét.** Ví dụ: "Nút CTA đặt cách xa ngón tay cái
   trên mobile → vi phạm **Fitts's Law** (Laws of UX)" — không chỉ nói "nút khó
   bấm".
3. **Phân lớp đánh giá** theo đúng thứ tự trong `06-evaluation-framework.md`,
   không nhảy cóc để tránh bỏ sót nhóm vấn đề.
4. **Tôn trọng ngữ cảnh nền tảng.** Nguyên tắc của Apple HIG không áp cứng cho
   Android, nguyên tắc Material không áp cứng cho iOS — xem `04-platform-guidelines.md`.

## Cấu trúc thư mục

| File | Nội dung | Nguồn gốc chính |
|---|---|---|
| `01-usability-heuristics.md` | 10 heuristics của Nielsen, mô hình affordance/signifier của Don Norman, nguyên tắc "Don't Make Me Think" của Steve Krug | Nielsen Norman Group, Don Norman, Steve Krug |
| `02-laws-of-ux.md` | Các "law" tâm lý học hành vi áp dụng cho UX (Fitts, Hick, Jakob, Miller...) | Jon Yablonski — lawsofux.com / sách *Laws of UX* |
| `03-visual-design-systems.md` | Gestalt principles, hệ thống spacing/typography/color, các nguyên tắc visual hierarchy | Tâm lý học Gestalt (Wertheimer, Koffka, Köhler); Adam Wathan & Steve Schoger — *Refactoring UI* |
| `04-platform-guidelines.md` | Quy chuẩn thiết kế chính thức theo nền tảng | Apple HIG, Google Material Design 3, Microsoft Fluent 2 |
| `05-accessibility-wcag.md` | Chuẩn accessibility bắt buộc/khuyến nghị | W3C WCAG 2.2 |
| `06-evaluation-framework.md` | Quy trình review thực tế + template output | Tổng hợp quy trình từ các nguồn trên — đây là phần "quy trình" do agent tổng hợp, không phải một nguyên tắc thiết kế riêng |
| `07-uxlagi-case-studies.md` | Case study thực tế minh hoạ các nguyên tắc trên (rập khuôn, affordance, màu sắc, dark mode, chọn mặt chữ, icon, design system...) | Nguyễn Đặng Việt Anh — uxlagi.com (Senior UX Engineer, Google Material Design & Uber Base) |

## Cách dùng khi review một design

1. Xác định category + platform (web responsive, iOS, Android, desktop...).
2. Thu thập input: screenshot, link Figma, hoặc mô tả luồng.
3. Chạy qua từng lớp theo `06-evaluation-framework.md`.
4. Output theo template chuẩn trong file đó: Strengths → Issues (theo mức độ
   nghiêm trọng, mỗi issue gắn nguyên tắc nguồn) → Đề xuất ưu tiên.
