# Evaluation Framework — Quy trình đánh giá UI/UX

File này là quy trình vận hành (không phải một nguyên tắc thiết kế riêng) để
Expert Design Agent dùng các nguồn kiến thức ở file `01`–`05` một cách có hệ
thống khi review một design cụ thể (web, mobile, tablet, desktop...).

## Bước 0 — Thứ tự ưu tiên đọc dữ liệu theo loại nguồn input (quy tắc tiết kiệm token)

Đây là rule **bắt buộc, ưu tiên cao nhất**, áp dụng trước mọi bước khác, với
**mọi loại nguồn input** — không riêng Figma. Nguyên tắc chung: luôn ưu tiên
đọc dữ liệu layout dưới dạng **text/số có cấu trúc**; đọc **ảnh UI (raster
screenshot)** luôn là phương án **tệ nhất, chỉ dùng khi bị ép buộc** — vì tốn
token nhiều hơn và kém chính xác hơn dữ liệu có cấu trúc cho hầu hết loại lỗi
(lệch alignment, spacing sai scale, màu không nhất quán, contrast không đạt).

### Tầng 1 — Nguồn có dữ liệu cấu trúc đầy đủ (luôn ưu tiên, đọc trước tiên)

Tuỳ loại nguồn, dùng đúng kênh text/số tương ứng — không có "một tool cho tất
cả", nhưng cùng một nguyên tắc:

| Loại nguồn | Đọc bằng gì | Lấy được |
|---|---|---|
| Figma | `get_metadata` rồi `get_design_context` | toạ độ, spacing, màu hex, type, cấu trúc layer |
| Web (site đang chạy/URL) | đọc DOM + computed CSS (devtools/inspect, không chụp màn hình) | spacing, màu, font thực tế đã render |
| Code (React/Vue/HTML/CSS/Tailwind...) | đọc trực tiếp source file | class, token spacing/color, breakpoint |
| Sketch/Adobe XD/Zeplin hoặc design tool khác có export spec | đọc JSON/CSS/spec export nếu có | tương đương Figma design context |
| SVG tĩnh (không phải raster) | đọc trực tiếp mã SVG (text/XML) | toạ độ, path, màu — không cần "nhìn" |

Với mọi nguồn ở tầng này, chạy hết các check đo được bằng số trước: alignment
(so toạ độ x/y), spacing (so với scale 4px/8px), contrast (tính trực tiếp từ
2 giá trị màu so với ngưỡng WCAG ở file `05`), type scale, màu lặp lại giữa
component cùng chức năng.

> **Cảnh báo quan trọng — "cấu trúc tồn tại" ≠ "giá trị thực giống nhau".**
> `get_metadata` (và tương đương: cây layer, DOM structure không kèm computed
> style) chỉ cho biết **một pattern/component tồn tại ở đâu** (id, tên, toạ
> độ, kích thước) — **không** cho biết giá trị màu/token đã resolve thực tế
> của từng instance. Tìm thấy cùng 1 tên frame (ví dụ "Pagination") xuất hiện
> ở N màn hình **không đồng nghĩa** với việc lỗi màu/contrast phát hiện ở 1
> instance cũng tồn tại ở N-1 instance còn lại — mỗi instance có thể đã được
> bind vào token khác nhau (kể cả khi style cũ từng giống nhau). Trước khi
> khẳng định một lỗi màu/contrast "lan rộng qua nhiều instance", phải gọi
> `get_design_context` (hoặc tối thiểu đọc được giá trị fill/variable đã
> resolve) **trên từng instance cụ thể** — không suy ra từ việc cấu trúc
> hoặc tên layer giống nhau. Nếu không đủ thời gian/budget để kiểm hết từng
> instance, phải nêu rõ trong report: "phát hiện tại 1 instance cụ thể (dẫn
> node-id), chưa xác minh các instance cùng tên khác" — không quy kết mức độ
> nghiêm trọng (vd: "Critical, ảnh hưởng N màn hình") dựa trên suy luận cấu
> trúc.
>
> **Quy tắc này áp dụng cho CẢ hai chiều, không chỉ khi báo lỗi.** Một nhận
> định "pattern X được tái sử dụng nhất quán ở N nơi" (ghi vào mục **Điểm
> mạnh**) mắc đúng lỗi suy luận tương tự nếu chỉ dựa vào việc N frame có
> **cùng tên** ở top-level (qua `get_metadata`), mà không mở children của
> *từng* instance ra để diff. "Có 1 frame tên 'Pagination' ở 6 nơi" và "cả 6
> nơi đều implement pagination theo cùng 1 cách" là hai claim khác nhau — chỉ
> claim đầu được miễn phí check thêm; claim thứ hai (nhất quán/tái sử dụng)
> bắt buộc phải diff children từng instance (ví dụ: instance này có các frame
> "Page N" con, instance khác chỉ có 1 text node tĩnh "Page 1 of 71" — khác
> hẳn nhau dù top-level đều tên "Pagination"). Không ghi "X nhất quán ở N màn
> hình" vào Điểm mạnh nếu chưa diff đủ N instance; nếu chỉ kiểm được một phần,
> ghi rõ "nhất quán ở các instance đã kiểm: [liệt kê], chưa kiểm hết toàn bộ."

### Tầng 2 — Nguồn chỉ có dữ liệu cấu trúc một phần

Dùng tối đa phần cấu trúc đọc được (ví dụ: code có class nhưng thiếu ảnh
preview, hoặc Figma file bị khoá quyền chỉ cho screenshot công khai). Chỉ bù
phần **không thể suy ra từ text** bằng ảnh, và chỉ crop đúng vùng thiếu dữ
liệu đó — không chụp toàn bộ lại những gì đã đọc được bằng text.

### Tầng 3 — Nguồn không có dữ liệu cấu trúc nào (ảnh UI thuần, PDF render,
ảnh chụp mockup in giấy...) — phương án ép buộc cuối cùng

Chỉ đọc ảnh khi thực sự không còn lựa chọn nào khác ở Tầng 1–2. Khi buộc phải
dùng ảnh:

- Crop đúng phạm vi đang nghi vấn, không đọc nguyên trang/nguyên flow nhiều
  màn hình trong 1 lần.
- Luôn ghi rõ trong report rằng giá trị đo được (màu, spacing, contrast...)
  chỉ là **ước lượng từ ảnh**, cần xác minh bằng dữ liệu thực nếu có thể lấy
  được sau đó (xem Bước 5).
- Không dùng ảnh để xác nhận lại một issue đã xác nhận được bằng số ở Tầng
  1–2 — đó là lãng phí token không cần thiết.

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
