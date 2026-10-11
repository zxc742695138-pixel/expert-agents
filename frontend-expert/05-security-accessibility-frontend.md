# Security & Accessibility ở tầng code frontend

Nguồn: [OWASP Top 10](https://owasp.org/www-project-top-ten/) và
[OWASP Cheat Sheet Series](https://cheatsheetseries.owasp.org) — chuẩn công
nghiệp cho lỗ hổng bảo mật web; W3C [WCAG 2.2](https://www.w3.org/TR/WCAG22/)
áp dụng cụ thể vào cách viết HTML/JSX/template (bổ sung, không thay thế, cho
`designer-expert/05-accessibility-wcag.md` ở tầng thiết kế).

## A. Lỗ hổng bảo mật frontend thường gặp (OWASP)

- **XSS (Cross-Site Scripting)** — Không bao giờ render HTML từ dữ liệu
  người dùng mà không sanitize. Cảnh báo cụ thể theo framework:
  - React: `dangerouslySetInnerHTML` là điểm rủi ro XSS trực tiếp — chỉ
    dùng khi nội dung đã được sanitize qua thư viện tin cậy (ví dụ DOMPurify),
    không dùng trực tiếp với dữ liệu thô từ API/user input.
  - Vue: `v-html` có rủi ro tương đương `dangerouslySetInnerHTML`.
  - Vanilla JS: gán trực tiếp vào `innerHTML` từ input người dùng là lỗi
    XSS cổ điển — dùng `textContent` khi chỉ cần hiển thị text thuần.
- **Lưu token/secret không an toàn ở client** — không lưu JWT/access token
  nhạy cảm vào `localStorage` nếu có thể tránh (dễ bị đánh cắp qua XSS vì
  JS nào cũng đọc được); ưu tiên `httpOnly` cookie khi backend hỗ trợ.
- **CSRF (Cross-Site Request Forgery)** — request thay đổi state (POST/PUT/
  DELETE) nên có CSRF token hoặc dựa vào `SameSite` cookie, không chỉ dựa
  vào việc "request được gọi từ frontend của mình".
- **Dependency có lỗ hổng biết trước** — kiểm tra định kỳ (`npm audit` hoặc
  tương đương) các dependency có CVE công khai; một thư viện lâu không cập
  nhật là rủi ro tích lũy, không phải vấn đề lý thuyết.
- **Expose thông tin nhạy cảm qua bundle JS** — API key bí mật, secret
  backend không nên nằm trong code chạy ở client (mọi thứ trong bundle JS
  gửi tới browser đều có thể bị đọc ngược) — chỉ các key/token được thiết kế
  để public (ví dụ public API key có giới hạn domain/scope) mới an toàn ở
  client.
- **Open redirect** — redirect dựa trên tham số URL do người dùng kiểm soát
  mà không validate domain đích có thể bị lợi dụng cho phishing.

## B. Accessibility ở tầng code (bổ sung cho WCAG ở tầng thiết kế)

- **Dùng đúng semantic HTML trước khi dùng ARIA** ("No ARIA is better than
  bad ARIA" — nguyên tắc nền của ARIA Authoring Practices, W3C) — `<button>`
  cho hành động click được, không phải `<div onClick>` + thêm ARIA role để
  "giả" thành button; thẻ heading (`<h1>`-`<h6>`) đúng cấp bậc, không chọn
  theo size chữ muốn hiển thị.
- **Mọi phần tử tương tác phải dùng được bằng bàn phím** — `tabIndex` hợp
  lý, không set `tabIndex="-1"` cho phần tử người dùng cần focus tới; custom
  component (dropdown, modal, custom checkbox) phải tự implement đúng hành
  vi bàn phím tương đương native element (Enter/Space để activate, Escape để
  đóng modal, Arrow key để di chuyển trong menu).
- **Ảnh có `alt` text có ý nghĩa** (rỗng `alt=""` chỉ khi ảnh thuần trang
  trí, không mang thông tin) — tương ứng WCAG 1.1.1.
- **Form input luôn có `<label>` liên kết đúng** (`htmlFor`/`for` khớp
  `id`, hoặc label bọc input) — không chỉ dùng `placeholder` làm label (mất
  khi người dùng bắt đầu nhập, và không được screen reader đọc như label).
- **Quản lý focus khi UI thay đổi động** — mở modal phải chuyển focus vào
  trong modal, đóng modal phải trả focus về phần tử đã mở nó; route change
  trong SPA nên thông báo cho screen reader (ví dụ focus vào heading trang
  mới) vì không có full page reload tự nhiên báo hiệu.
- **Tôn trọng `prefers-reduced-motion`** trong code animation (CSS media
  query hoặc check trong JS) — tương ứng WCAG 2.3.3, quan trọng hơn ở tầng
  code vì đây là nơi quyết định animation có thực sự bị tắt/giảm hay không.

## Cách dùng khi review

- Lỗ hổng bảo mật tìm được phải nêu rõ **kịch bản khai thác cụ thể** (ai có
  thể lợi dụng, bằng cách nào, hậu quả gì) — không chỉ dán nhãn "không an
  toàn" mà không giải thích được đường tấn công thực tế.
- Nếu một pattern *có thể* là lỗ hổng nhưng cần biết thêm ngữ cảnh (ví dụ:
  `dangerouslySetInnerHTML` với nội dung đã sanitize ở nơi khác mà reviewer
  chưa thấy) — hỏi rõ hoặc nêu "cần xác minh nguồn dữ liệu đầu vào đã được
  sanitize ở đâu" thay vì khẳng định chắc chắn là lỗi.
