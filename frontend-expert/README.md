# Frontend Expert — Knowledge Base thẩm định code frontend

Đây là thư mục tri thức nền cho vai trò **Expert Frontend Code Reviewer**:
dùng để thẩm định (review) code frontend — JavaScript/TypeScript, React/Vue/
Angular, HTML/CSS — về chất lượng, hiệu năng, bảo mật và khả năng bảo trì.
Cùng triết lý và cấu trúc với `designer-expert/` trong repo này, áp dụng cho
code thay vì UI/UX.

## Nguyên tắc cốt lõi khi dùng thư mục này

1. **Không tự suy diễn.** Mọi nhận định khi review code phải truy được về
   một nguyên tắc cụ thể trong các file `01`–`05`, hoặc một nguồn ngoài được
   dẫn rõ ràng (link, tên tác giả, tên tài liệu/spec). Không dùng câu mơ hồ
   kiểu "code chưa clean", "chưa tối ưu" mà không gắn với nguyên tắc nào.
2. **Luôn trích nguồn khi nhận xét.** Ví dụ: "Component gọi `setState` trong
   render body không có điều kiện dừng → vi phạm **Rules of Hooks** (React
   docs)" — không chỉ nói "code này có bug".
3. **Phân lớp đánh giá** theo đúng thứ tự trong `06-evaluation-framework.md`
   (đúng trước, sạch sau, nhanh sau đó, an toàn/tiếp cận được xuyên suốt),
   không nhảy cóc để tránh bỏ sót nhóm vấn đề.
4. **Tôn trọng ngữ cảnh framework/stack.** Quy ước của React không áp cứng
   cho Vue, convention của một codebase cụ thể (nếu có CLAUDE.md/style guide
   riêng của repo đang review) luôn được tôn trọng trước style guide chung ở
   đây — xem `03-framework-patterns.md`.
5. **Ưu tiên đọc mã nguồn trực tiếp (text) hơn là chạy app/chụp ảnh UI khi
   review code.** Source code tự nó đã là dữ liệu có cấu trúc — review bằng
   cách đọc trực tiếp file (Read/Grep), không cần "nhìn" app chạy trừ khi
   đang xác minh hành vi runtime cụ thể không thể suy ra từ đọc code (xem
   Bước 0 trong `06-evaluation-framework.md`).

## Cấu trúc thư mục

| File | Nội dung | Nguồn gốc chính |
|---|---|---|
| `01-clean-code-principles.md` | Nguyên tắc đặt tên, hàm, SOLID, DRY/KISS/YAGNI | Robert C. Martin — *Clean Code*; *Design Patterns* (Gang of Four) cho SOLID gốc |
| `02-javascript-typescript-best-practices.md` | Idioms JS/TS hiện đại, async, type safety | Airbnb JavaScript/TypeScript Style Guide; MDN Web Docs; TypeScript Handbook (microsoft) |
| `03-framework-patterns.md` | Quy ước/anti-pattern React, Vue, Angular | React docs (react.dev) — Rules of Hooks, Thinking in React; Vue.js Style Guide (vuejs.org); Angular Style Guide (angular.dev) |
| `04-performance-web-vitals.md` | Core Web Vitals, bundle size, rendering performance | web.dev (Google) — Core Web Vitals; Addy Osmani — *Essential Image Optimization*, bài viết về bundle/tree-shaking |
| `05-security-accessibility-frontend.md` | Lỗ hổng bảo mật frontend thường gặp, accessibility ở tầng code | OWASP Top 10 & OWASP Cheat Sheet Series; W3C WCAG 2.2 (semantic HTML/ARIA) |
| `06-evaluation-framework.md` | Quy trình review thực tế + template output | Tổng hợp quy trình từ các nguồn trên — phần "quy trình" do agent tổng hợp, không phải nguyên tắc riêng |

## Cách dùng khi review code

1. Xác định stack (JS thuần/TS, framework, có test suite/lint config có sẵn
   không) và convention riêng của repo (CLAUDE.md, .eslintrc, style guide nội
   bộ nếu có — luôn ưu tiên hơn style guide chung ở đây khi có xung đột).
2. Đọc trực tiếp diff/file bằng Read/Grep — không chạy app trừ khi cần xác
   minh hành vi runtime.
3. Chạy qua từng lớp theo `06-evaluation-framework.md`.
4. Output theo template chuẩn trong file đó: Strengths → Issues (theo mức độ
   nghiêm trọng, mỗi issue gắn nguyên tắc nguồn + file:line cụ thể) → Đề xuất
   ưu tiên.
