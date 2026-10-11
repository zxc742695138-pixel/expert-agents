# Evaluation Framework — Quy trình thẩm định code frontend

File này là quy trình vận hành (không phải một nguyên tắc riêng) để Expert
Frontend Code Reviewer dùng các nguồn kiến thức ở file `01`–`05` một cách có
hệ thống khi review code cụ thể (file, diff, PR).

## Bước 0 — Ưu tiên đọc mã nguồn trực tiếp, hạn chế chạy app/chụp ảnh

Nguyên tắc tương tự `designer-expert/06-evaluation-framework.md` (ưu tiên dữ
liệu có cấu trúc hơn ảnh), áp dụng cho code:

1. **Đọc trực tiếp file/diff bằng Read/Grep** — source code tự nó đã là dữ
   liệu text có cấu trúc đầy đủ (không cần "render" ra mới đọc được như
   UI). Đây luôn là nguồn chính, chi phí thấp nhất, chính xác nhất.
2. **Đọc test suite hiện có (nếu có)** trước khi đoán hành vi runtime — test
   pass/fail, test case đã viết sẵn cho biết nhiều về hành vi mong đợi hơn
   là đoán từ đọc logic.
3. **Chạy app/viết test để xác minh hành vi — chỉ khi bắt buộc**, tức khi
   không thể xác định chắc chắn chỉ từ đọc code (ví dụ: race condition phụ
   thuộc thời gian thực tế, hành vi browser-specific, CSS rendering thực
   tế). Khi phải chạy, chỉ chạy đúng phần nghi vấn (1 component/1 flow cụ
   thể), không chạy lại toàn bộ app để xác nhận thứ đã rõ từ đọc code.

> **Cảnh báo quan trọng — "tìm thấy pattern giống nhau" ≠ "mọi nơi có pattern
> đó đều có cùng lỗi/hành vi".** Grep ra N chỗ gọi cùng 1 hàm, hoặc N file có
> tên/cấu trúc tương tự, chỉ xác nhận **pattern tồn tại ở đó** — không xác
> nhận mỗi chỗ đều mắc đúng lỗi đã thấy ở 1 chỗ. Một hàm có thể được gọi với
> input khác nhau ở mỗi nơi, hoặc đã được fix ở một số chỗ nhưng chưa đồng
> bộ ở chỗ khác. Trước khi khẳng định "lỗi X ảnh hưởng N chỗ", phải đọc đủ
> code tại **từng chỗ cụ thể**, không suy ra từ việc Grep ra kết quả giống
> tên. Điều này đúng cho cả 2 chiều: không khẳng định lỗi lan rộng, và cũng
> không khẳng định "pattern Y được áp dụng nhất quán ở N chỗ" (một dạng
> Strength) nếu chưa đọc hết N chỗ đó.

## Bước 1 — Xác định context trước khi review

- **Stack**: JS thuần/TypeScript, framework (React/Vue/Angular/khác), có
  lint/type-check config sẵn không.
- **Phạm vi**: review toàn bộ file, hay chỉ diff/PR (nếu là diff, luôn đọc
  thêm context xung quanh — code trước/sau đoạn thay đổi — để không nhận
  định sai do thiếu ngữ cảnh).
- **Convention riêng của repo**: đọc CLAUDE.md, `.eslintrc`, style guide nội
  bộ nếu có — convention riêng luôn ưu tiên hơn style guide chung ở các file
  `01`–`05`.
- **Có test suite không**: nếu có, chạy test trước khi review để biết trạng
  thái pass/fail hiện tại, tránh báo "bug" cho thứ test đã cover và pass.

## Bước 2 — Review theo 5 lớp, đúng thứ tự

Thứ tự đi từ "có đúng không" → "có sạch/bảo trì được không" → "có đúng
convention framework không" → "có nhanh không" → "có an toàn/tiếp cận được
không", để không bị cuốn vào style trước khi xác nhận code đúng.

| Lớp | Nguồn | Câu hỏi trọng tâm |
|---|---|---|
| **1. Correctness & bugs** | Đọc logic trực tiếp, test suite | Code có làm đúng điều nó định làm không? Edge case (null/undefined/empty/race condition) được xử lý chưa? |
| **2. Code quality & maintainability** | `01` | Naming rõ không? Hàm/component có làm đúng 1 việc? Có vi phạm SOLID/DRY/KISS rõ ràng? |
| **3. Framework/stack idioms** | `02`, `03` | Có dùng đúng idiom JS/TS hiện đại? Có vi phạm Rules of Hooks / convention framework? |
| **4. Performance** | `04` | Có pattern gây re-render/reflow thừa, bundle nặng không cần thiết, chặn main thread? |
| **5. Security & accessibility** | `05` | Có lỗ hổng XSS/CSRF/leak secret rõ ràng? Có dùng đúng semantic HTML/quản lý focus? |

Với mỗi lớp, chỉ ghi nhận **issue cụ thể gắn với 1 nguyên tắc + vị trí
file:line** — không đưa nhận xét chung không truy được nguồn hoặc không chỉ
ra được chỗ cụ thể trong code.

## Bước 3 — Phân loại mức độ nghiêm trọng

- **Critical** — Bug gây sai kết quả/crash ở luồng chính, hoặc lỗ hổng bảo
  mật khai thác được thực tế (XSS, leak secret, v.v.).
- **Major** — Vi phạm rõ ràng gây khó bảo trì/dễ sinh bug tương lai, hoặc
  vấn đề hiệu năng có tác động đo được đến Core Web Vitals, nhưng không
  crash/không khai thác được ngay.
- **Minor** — Vi phạm style/convention nhất quán, không ảnh hưởng hành vi
  runtime hiện tại.
- **Nitpick** — Chi tiết tinh chỉnh (polish), có thể để lại cho reviewer
  khác quyết định, không block merge.

## Bước 4 — Output report theo template chuẩn

```markdown
# Review: <tên file/PR> — <stack/framework>

## Tổng quan
<1-2 câu đánh giá tổng thể>

## Điểm mạnh
- <Điểm mạnh 1> — áp dụng đúng <tên nguyên tắc, nguồn>
- ...

## Vấn đề

### Critical
- **[file:line]** — <mô tả vấn đề>
  - Vi phạm: **<Tên nguyên tắc>** (<nguồn: file + tên gốc, ví dụ "Rules of Hooks — react.dev">)
  - Kịch bản thực tế: <input/hành động nào gây ra hậu quả gì>
  - Đề xuất: <fix cụ thể, khả thi, kèm snippet nếu cần>

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

- Nếu không đủ ngữ cảnh để khẳng định (vd: hành vi phụ thuộc code ở file
  khác chưa đọc, hoặc cần chạy thực tế để xác nhận race condition), nói rõ
  "cần xác minh thêm" thay vì đoán.
- Nếu một vấn đề nằm ở ranh giới giữa 2 nguyên tắc (vd: vừa là performance
  vừa là code quality), có thể trích cả hai.
- Không bịa thêm "rule" không có trong các file `01`–`05` hoặc không dẫn
  được nguồn ngoài rõ ràng. Nếu cần một nguyên tắc mới không có sẵn trong
  kho này, phải nói rõ đó là bổ sung ngoài kho tri thức, kèm nguồn.
- Luôn tôn trọng convention riêng của repo đang review (CLAUDE.md,
  `.eslintrc`) khi nó khác với style guide chung ở đây — convention nội bộ
  của 1 codebase cụ thể là nguồn có độ ưu tiên cao hơn style guide tổng quát.
