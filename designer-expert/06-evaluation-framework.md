# Evaluation Framework — Quy trình đánh giá UI/UX

File này là quy trình vận hành (không phải một nguyên tắc thiết kế riêng) để
Expert Design Agent dùng các nguồn kiến thức ở file `01`–`05` một cách có hệ
thống khi review một design cụ thể (web, mobile, tablet, desktop...).

## Bước 1 — Xác định context trước khi review

Luôn hỏi/hoặc tự xác định rõ trước khi đánh giá:

- **Category**: web (responsive/desktop), mobile (iOS/Android), tablet,
  desktop app...
- **Loại màn hình**: landing page, dashboard, form/checkout, onboarding,
  empty state, error state...
- **Đối tượng người dùng**: B2C đại chúng, B2B chuyên môn cao, nội bộ...
- **Input có sẵn**: screenshot tĩnh, Figma link (có thể đo spacing/color
  chính xác), hay mô tả bằng lời.

> Nếu input là ảnh tĩnh, không có giá trị màu/spacing chính xác, phải nêu rõ
> giới hạn này trong report (xem lưu ý ở `05-accessibility-wcag.md`).

## Bước 2 — Review theo 5 lớp, đúng thứ tự

Thứ tự này đi từ "có dùng được không" → "có đẹp/mạch lạc không" → "có đúng
chuẩn nền tảng/pháp lý không", để không bị cuốn vào chi tiết thẩm mỹ trước khi
xác nhận usability cơ bản.

| Lớp | Nguồn | Câu hỏi trọng tâm |
|---|---|---|
| **1. Usability heuristics** | `01` | Người dùng có biết đang ở đâu, làm gì, có thể undo không? Lỗi được ngăn/xử lý thế nào? |
| **2. Behavioral/cognitive laws** | `02` | Target có dễ chạm? Số lựa chọn có quá tải? Có tận dụng đúng tâm lý người dùng không? |
| **3. Visual design system** | `03` | Hierarchy có rõ? Spacing/typography/color có nhất quán theo hệ thống? |
| **4. Platform conformance** | `04` | Có đúng convention của iOS/Android/web mà sản phẩm đang nhắm tới? |
| **5. Accessibility** | `05` | Có đạt baseline WCAG 2.2 AA (contrast, target size, không chỉ dùng màu để truyền tin)? |

Với mỗi lớp, chỉ ghi nhận **issue cụ thể gắn với 1 nguyên tắc** — không đưa
nhận xét chung không truy được nguồn.

## Bước 3 — Phân loại mức độ nghiêm trọng

- **Critical** — Chặn người dùng hoàn thành tác vụ chính, hoặc vi phạm WCAG
  mức A/AA.
- **Major** — Gây khó khăn/nhầm lẫn đáng kể, tăng lỗi hoặc thời gian hoàn
  thành tác vụ rõ rệt, nhưng không chặn hoàn toàn.
- **Minor** — Ảnh hưởng nhỏ đến trải nghiệm hoặc nhất quán hệ thống, không
  cản trở tác vụ.
- **Nitpick** — Chi tiết tinh chỉnh (polish), không ảnh hưởng chức năng/
  usability, chỉ nâng cao chất lượng thẩm mỹ.

## Bước 4 — Output report theo template chuẩn

```markdown
# Review: <tên màn hình/flow> — <category/platform>

## Tổng quan
<1-2 câu đánh giá tổng thể, không chấm điểm số trừ khi được yêu cầu>

## Điểm mạnh
- <Điểm mạnh 1> — áp dụng đúng <tên nguyên tắc, nguồn>
- ...

## Vấn đề

### Critical
- **[Vị trí/element]** — <mô tả vấn đề>
  - Vi phạm: **<Tên nguyên tắc>** (<nguồn: file + tên gốc, ví dụ "Fitts's Law — Laws of UX">)
  - Tác động thực tế: <vì sao ảnh hưởng người dùng>
  - Đề xuất: <fix cụ thể, khả thi>

### Major
- ...

### Minor
- ...

### Nitpick
- ...

## Đề xuất ưu tiên
1. <Quick win — effort thấp, impact cao>
2. <Việc cần đầu tư lớn hơn>
```

## Bước 5 — Nguyên tắc ứng xử khi không chắc chắn

- Nếu không đủ dữ liệu để khẳng định (vd: không đo được contrast chính xác từ
  ảnh), nói rõ "cần xác minh thêm" thay vì đoán.
- Nếu một vấn đề nằm ở ranh giới giữa 2 nguyên tắc (vd: vừa là Gestalt
  Proximity vừa là spacing system của Refactoring UI), có thể trích cả hai.
- Không bịa thêm "luật" hoặc "nguyên tắc" không có trong các file `01`–`05`
  hoặc không dẫn được nguồn ngoài rõ ràng. Nếu cần một nguyên tắc mới không có
  sẵn trong kho này, phải nói rõ đó là bổ sung ngoài kho tri thức, kèm nguồn.
