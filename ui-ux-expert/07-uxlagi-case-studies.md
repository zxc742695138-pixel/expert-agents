# Case Studies — uxlagi.com

Nguồn: [uxlagi.com](https://uxlagi.com) — blog cá nhân của Nguyễn Đặng Việt Anh
(Senior UX Engineer, nhóm Material Design tại Google; sau đó làm UX Engineer
cho hệ thống thiết kế Base của Uber). Nội dung © Nguyễn Đặng Việt Anh,
2018–2021, chia sẻ theo license Attribution 3.0 Unported. 28 bài, viết bằng
tiếng Việt, kể lại kinh nghiệm thực chiến tại Google Play, Google Store,
Material Design và Uber Base.

> File này khác các file `01`–`05`: nó không phải một bộ nguyên tắc lý thuyết
> riêng, mà là **case study thực tế** minh hoạ cho các nguyên tắc đã có. Khi
> review, có thể trích dẫn song song — ví dụ: "màu sắc loè loẹt làm giảm độ
> tin cậy, giống case Dropbox Pricing (uxlagi.com — Sử dụng màu sắc)" — để
> nhận định cụ thể và thuyết phục hơn là chỉ nêu tên law chung.

## A. Hiểu người dùng trước khi thiết kế

- **Rập khuôn** (Stereotype) — Dùng khuôn mẫu quen thuộc (vai thiện/ác, tông
  truyện phiêu lưu/lãng mạn) giúp truyền đạt ý tưởng nhanh, dễ nhớ (case:
  Trader Joe's dùng hình cướp biển cho thông điệp "hàng từ phương xa"). Rủi
  ro: dễ gây động chạm văn hoá nếu không cẩn trọng (case: đội Washington
  Redskins dùng biểu tượng bộ lạc da đỏ). Chỉ nên dùng khi mục tiêu truyền đạt
  rõ ràng và đã cân nhắc đánh đổi.
- **Hướng cây** (Biophilia) — Con người có nhu cầu bản năng kết nối với thiên
  nhiên (yên tâm hơn khi nhìn cây cỏ, nước). Google dành cả năm cải tạo
  khoảng không giữa các tòa nhà ở Mountain View để trồng cây — ứng dụng được
  cả vào thiết kế không gian vật lý, không chỉ UI số.
- **Cách hiểu** (Mental model) — Người dùng luôn tự xây một mô hình tinh thần
  về cách hệ thống hoạt động, dù mô hình đó có thể hoàn toàn sai nhưng vẫn
  *tự nhất quán* với họ (ví dụ vui: tin thang máy chạy bằng bóng bay, hay bằng
  người tí hon). Vị trí của người thiết kế UX là đứng giữa người dùng và sản
  phẩm, dịch chuyển sự hiểu biết giữa hai bên qua hình ảnh, màu sắc, âm thanh.
  *Không có việc người dùng "làm đúng hay làm sai", chỉ có cách hiểu khác
  nhau.*
- **UX cho bố mẹ** — Case thực tế: bố không hiểu "link" nghĩa là gì khi được
  hướng dẫn cài bàn phím tiếng Việt (gần 10 bước qua hệ thống Android); icon
  mic/camera thay đổi hình dạng qua nhiều thế hệ làm người lớn tuổi không
  nhận ra. Kết luận: *"khi người dùng không cần hiểu hệ thống, cách thiết kế
  tốt nhất là không bắt người dùng phải hiểu."* Áp dụng khi review sản phẩm
  có đối tượng người dùng lớn tuổi/ít rành công nghệ: luôn hỏi "liệu bước này
  có đòi hỏi hiểu biết hệ thống mà người dùng không có?"
- **Cách hỏi** (User research) — Quy trình 6 bước phỏng vấn người dùng của
  Google Ventures: (1) Xin chào — làm người đối diện thoải mái; (2) Giới
  thiệu — nêu rõ mục đích, cam kết bảo mật, quyền từ chối; (3) Khám phá —
  dùng câu hỏi mở ("Bạn thường làm gì khi...") thay vì câu hỏi đóng; (4) Dùng
  thử — quan sát hành vi thật, khuyến khích người dùng nói ra suy nghĩ; (5)
  Tổng kết; (6) Tạm biệt. Nguyên tắc cốt lõi: nghe nhiều hơn nói, mục đích
  phỏng vấn là *hiểu* người dùng, không phải hỏi họ giải pháp. Trích Steve
  Krug: "Phỏng vấn 1 người ở giai đoạn đầu làm sản phẩm giá trị bằng phỏng
  vấn 50 người ở giai đoạn cuối."
- **Nghề UX** — 8 vai trò trong một đội hình UX hoàn chỉnh: Interaction
  Designer (dễ dùng hơn), Visual Designer (hấp dẫn hơn), UX Writer/Content
  Designer (dễ hiểu hơn), UX Engineer (hiện thực hoá), UX Researcher (đúng
  đối tượng), Industrial Designer (sờ nắn được), Motion Designer (hệ thống dễ
  hiểu hơn qua chuyển động), Program Manager (tránh xung đột quy trình). Hữu
  ích khi review để xác định loại vấn đề thuộc trách nhiệm vai trò nào.

## B. Dễ dùng & môi trường

- **Tính dễ dùng** (Affordance) — 3 tiêu chí cụ thể để đánh giá affordance
  của một chi tiết UI: (1) hình dạng diễn tả đúng cách hoạt động, (2) không
  dễ bị lẫn lộn, (3) tín hiệu phản hồi tốt. Case kinh điển "Cửa Norman" (Don
  Norman, *The Design of Everyday Things*) — nắm cửa trông như phải kéo
  nhưng hoá ra là cửa đẩy.
- **Nhà trần cao** — Không gian trần cao kích thích sáng tạo và tư duy mở
  (bảo tàng, sảnh khách sạn); trần thấp/không gian nhỏ giúp tập trung (góc
  đọc sách, bàn làm việc riêng). Có thể ẩn dụ sang bố cục màn hình: layout
  "thoáng" phù hợp giai đoạn khám phá/sáng tạo, layout "đặc, gọn" phù hợp tác
  vụ cần tập trung hoàn thành nhanh.
- **Ít hay nhiều?** — Không có đáp án đúng tuyệt đối giữa tối giản (phương
  Tây, ít chi tiết, nhiều khoảng trống) và dày đặc thông tin (phương Đông,
  nhiều chi tiết cùng lúc) — tác giả tự thử nghiệm khảo sát bạn đọc và nhận
  ra cả hai gu đều có người thích ngang nhau. Bài học quan trọng: **không áp
  đặt gu thẩm mỹ cá nhân của người thiết kế làm chuẩn**, phải kiểm chứng bằng
  dữ liệu/khảo sát thật với đúng đối tượng mục tiêu.
- **Khoảng trống** (Whitespace) — Phân hai loại: *Micro whitespace* (khoảng
  cách giữa chữ, dòng, chi tiết nhỏ — ảnh hưởng trực tiếp đến readability) và
  *Macro whitespace* (khoảng không giữa các block bố cục lớn — dùng để dẫn
  hướng hành vi, tạo thiết kế sạch, giúp người dùng nhớ nội dung lâu hơn, case
  minh hoạ: trang chủ Google). Khoảng trống không phải không gian lãng phí mà
  là một lựa chọn thiết kế chủ động.

## C. Kể chuyện & dẫn dắt mắt

- **Bôi đậm** (Highlighting) — 5 cách bôi đậm thông tin: qua **ý nghĩa** (ví
  dụ nhấn từ khác nhau trong câu "Tôi thích ăn cơm" tạo ra nghĩa khác nhau),
  qua **màu sắc**, qua **bố cục** (hình to/nhỏ trong App Store làm tăng tải
  app hơn 1747%), qua **chuyển động** (icon nhảy trên MacOS khi có thông báo
  quan trọng), qua **âm thanh** (tiếng "b-líp" của Facebook — nhẹ, không gấp —
  so với "bóp-đing" của Messenger — nổi hơn, cần trả lời ngay). Bài học: mỗi
  khi bôi đậm, phải tự hỏi "tôi đang nhấn mạnh điều gì, và mức độ bôi đậm có
  khớp với mức độ ưu tiên thật của thông tin không?" Lạm dụng bôi đậm nhiều
  chỗ cùng lúc sẽ làm mất hiệu quả hoàn toàn (tường quảng cáo dán đầy, không
  biết nhìn vào đâu).
- **Phân cấp thị giác** (Visual hierarchy) — Case "Google bị bắt nạt": trang
  bán điện thoại Nexus/Pixel trên BestBuy.com bị chen giữa quá nhiều thông
  điệp cạnh tranh (giá, bảo hiểm, nhà tài trợ, nút mua) nên không kể được câu
  chuyện sản phẩm; so với trang tự thiết kế trên Google Store chỉ tập trung 1
  câu chuyện chính duy nhất mỗi màn hình. Nguyên tắc rút ra: **mọi giao diện
  đều đang kể một câu chuyện** — nếu không chủ động thiết kế phân cấp (to/nhỏ,
  đậm/nhạt theo đúng tầm quan trọng), câu chuyện sẽ bị kể sai hoặc không ai
  hiểu được kể gì (case đối chiếu: RetroArch liệt kê hết thuật ngữ kỹ thuật mà
  quên làm nổi nút chơi game chính).
- **Thẳng hàng** (Alignment) — Lỗi rất phổ biến: gióng các thành phần theo
  *đường biên/viền* của khung chứa, thay vì gióng theo *nội dung bên trong*
  (ví dụ: ô tìm kiếm nên gióng chữ "visual..." với các chữ "All", "Optical..."
  bên dưới, không phải theo viền ô vuông). Ở mức cao hơn, cần gióng theo
  *trọng lượng thị giác* (visual weight) — mỗi hình có tâm điểm riêng không
  nhất thiết ở giữa hình học, nên không thể chỉ dùng thước đo máy móc mà phải
  tinh chỉnh bằng mắt ("đẩy điểm ảnh" / pixel pushing) đến khi cân.
- **Lưới** (Grid) — Lưới là hệ thống đường kẻ dọc/ngang vô hình giúp sắp xếp
  nội dung thẳng hàng với nhau; kể cả sau khi ẩn lưới đi, mắt người xem vẫn
  cảm nhận được sự liên kết. Dùng được ở cả bố cục digital (web, app) và vật
  lý (sắp sách lên tủ, thiết kế bao bì sản phẩm).
- **Bộ luật Gestalt** — 6 luật với ví dụ ứng dụng cụ thể trong UI thật:
  **Similarity** (các tin tức cùng độ giật gân có cùng kích cỡ hiển thị),
  **Proximity** (mỗi kết quả tìm kiếm Google nhóm link/tóm tắt/ảnh gần nhau),
  **Closure** (thiết kế tối giản "KISS" dựa vào não tự lấp đầy phần thiếu),
  **Symmetry** (đặt nút/thông báo quan trọng ở chính giữa màn hình để thu hút
  chú ý), **Continuation** (chữ bị cắt ở lề dưới màn hình báo hiệu còn nội
  dung để cuộn tiếp; hàng hoá bị cắt ở lề để báo còn để lướt ngang),
  **Common Fate** (các yếu tố chuyển động cùng hướng được xem là liên quan,
  dùng cho motion design mượt mà; ngược lại quảng cáo nhảy lệch hướng để "ăn
  cướp" sự chú ý).

## D. Màu sắc

- **Ý nghĩa màu sắc** — Màu có ảnh hưởng sinh lý qua bước sóng: màu ấm
  (đỏ, vàng) bước sóng dài, kích thích thần kinh, tăng cảm giác đói (dùng
  trong logo đồ ăn); màu lạnh (xanh, tím) bước sóng ngắn, gây dễ chịu (dùng
  cho nước uống, trang y tế). Quan trọng hơn: **ý nghĩa văn hoá của màu
  không cố định toàn cầu** — ví dụ vàng là "hạnh phúc" ở phương Tây nhưng là
  "phản bội" hay "tang tóc" ở nơi khác; trắng là "tinh khiết" ở phương Tây
  nhưng là "tang tóc" ở nhiều nước Á. Case thực: bộ màu mặc định của Material
  Design (ảnh hưởng triết lý Bauhaus phương Tây — trắng/đen/xám/xanh) bị các
  công ty Hàn Quốc/Nhật Bản phản hồi là không phản ánh được bản sắc thương
  hiệu của họ (thị trường Đông Á thường ưa màu ấm, rực rỡ hơn).
- **Sử dụng màu sắc** — Case Dropbox: redesign trang Pricing dùng bảng màu
  táo bạo (xanh dương đậm + hồng san hô) khiến mọi chỉ số chuyển đổi sụt
  giảm nghiêm trọng, phải revert về thiết kế cũ đơn giản hơn — nguyên nhân là
  màu mè làm giảm cảm giác *rõ ràng và tin cậy* cần có ở một trang giá cả.
  Nguyên tắc rút ra: **dùng càng ít màu càng tốt** — chọn 1 màu chủ đạo rồi
  dùng scale tint/shade (như Material 50→900) để tạo phân cấp, thay vì nhiều
  màu tuỳ hứng. Khi cần dùng nhiều màu (ví dụ phân biệt trạng thái Gmail vs
  Inbox), phải kiểm tra độ hài hoà (công cụ Adobe Color) và luôn đảm bảo độ
  tương phản đạt chuẩn đọc được (Material Color Tool / WebAIM).
- **Chế độ Tối** (Dark Mode) — 3 nhóm lỗi điển hình khi triển khai ở quy mô
  lớn (case Uber): **lỗi từ ngữ** (team đặt tên màu không thống nhất — "Gray"
  vs "Grey" vs mã hex trực tiếp — dẫn đến không ai biết sửa chỗ nào khi có
  yêu cầu đổi); **lỗi lười** (designer dùng màu theo cảm tính thay vì theo
  công thức hệ thống, nên giải pháp hiệu quả nhất là làm tool tự động convert
  thay vì chỉ ra quy định bằng văn bản); **lỗi "võ công"** (trộn lẫn 2 trường
  phái xử lý độ sâu không tương thích — *Phân Tầng*: dùng độ sáng tăng dần
  theo elevation kiểu Google/Facebook, vs *Tương Phản*: ưu tiên contrast rõ
  ràng để đọc kiểu Twitter/Medium, đảm bảo luôn ≥4.5:1 theo WCAG). Câu chốt
  đáng nhớ: *"Mục tiêu thiết kế Dark Mode không phải là 'màu gì?' mà là 'giúp
  ai?'"* — chọn trường phái dựa trên nhu cầu thực của người dùng (tài xế Uber
  cần đọc nhanh trong lúc di chuyển), không dựa trên gu thẩm mỹ nội bộ team.

## E. Chữ & biểu tượng

- **Xuất xứ của chữ** — Lược sử chữ viết: chữ tượng hình Ai Cập → chữ tượng
  thanh của người Phoenicia (mượn hình, chuyển thành âm, tiện cho buôn bán) →
  chữ có chân La Mã (phát sinh từ kỹ thuật khắc đá, giải quyết ảo ảnh thị
  giác "béo bụng giữa") → chữ viết thường thời Alcuin (thêm khoảng cách giữa
  chữ, viết nhanh hơn chữ hoa). Bài học xuyên suốt: **mọi hình dạng chữ trong
  lịch sử đều tiến hoá để giải quyết một vấn đề thực tế cụ thể** (tốc độ viết,
  vật liệu, độ dễ đọc) — không phải chỉ vì "đẹp".
- **Chọn mặt chữ** (Typeface) — Quy trình 4 bước theo đúng thứ tự ưu tiên:
  (1) chọn theo **loại thông tin** cần truyền tải (biển giao thông cần mặt
  chữ rõ khi lái xe tốc độ cao: DIN, Highway Gothic); (2) chọn theo **phương
  tiện hiển thị** (in giấy vs màn hình nhỏ cần mặt chữ khác nhau); (3) chọn
  theo **thẩm mỹ** (hiện đại hay cổ kính, có hợp với các mặt chữ khác đang
  dùng không); (4) chọn theo **ngân sách** (mặt chữ nhái thường thiếu kerning
  chuẩn, dễ gây đọc nhầm — ví dụ "Nha Khoa" đọc nhầm thành "Nhà Kho A"). Trích
  dẫn đáng nhớ: *"95% hầu hết các giao diện là chữ"* (Oliver Reichenstein) —
  chọn mặt chữ cẩn thận có thể cải thiện phần lớn sản phẩm ngay lập tức, mà
  không cần thay đổi gì khác.
- **Hình biểu tượng** (Icon) — Nguyên tắc cốt lõi: *icon cần rõ ràng hơn là
  đẹp*, đặc biệt khi người dùng cần tìm gấp (case: icon nhà vệ sinh nam/nữ chỉ
  khác nhau một chi tiết tam giác nhỏ, gây nhầm lẫn thực tế). Quy trình vẽ
  icon hệ thống ở Uber: nghĩ cấu tạo vật thể ngoài đời → quy về các hình học
  cơ bản (hình chữ nhật, tam giác). 3 tuyệt chiêu kỹ thuật: lưới 24×24 với các
  đỉnh ở toạ độ nguyên (để nét không bị mờ/răng cưa), càng ít nét càng tốt
  (khung sườn phải tự nhận ra được ý nghĩa), nét đậm rộng 3px (để đọc được cả
  khi đang di chuyển/chạy). Icon cũng cần lặp lại kiểm tra với người dùng thật
  (case: icon "trẻ em" hình cái nôi ban đầu bị hiểu lầm là "mỳ tôm", phải sửa
  lại).

## F. Vận hành ở quy mô tổ chức

- **Hệ thống thiết kế** (Design System) — Thông qua hội thoại giả định tại
  Uber, bài viết bác bỏ 3 quan niệm sai phổ biến về design system: (1) chỉ là
  **tài liệu quy chuẩn** viết sẵn — không đủ, vì không thể dự đoán hết mọi
  trường hợp sử dụng thực tế (case: biển tàu điện New York mất 4 thập kỷ mới
  đồng bộ dù có hướng dẫn chi tiết từ 1970); (2) chỉ là **ép buộc tuân lệnh**
  — gây tác dụng ngược, team sợ bị phạt nên không còn chia sẻ lý do hợp lý
  khi cần đi khác chuẩn (ví dụ: team taxi dùng xanh đậm thay đen vì tài xế
  lái đêm đỡ mỏi mắt — một ngoại lệ có cơ sở); (3) chỉ là **thư viện tái sử
  dụng component** — không đủ nhân lực của 1 nhóm nhỏ để cover hết nhu cầu của
  hàng trăm người dùng nội bộ. Bản chất đúng theo tác giả: **hệ thống thiết kế
  là một quá trình liên tục thu thập kinh nghiệm nên/không nên từ các nhóm
  đang trực tiếp đối mặt với người dùng thật, rồi viết kinh nghiệm đó thành
  tài liệu sống** — để tránh tình trạng toàn bộ tri thức chỉ nằm trong đầu một
  người (case cảnh báo: thiết kế trưởng duy nhất nắm hết "chuẩn đúng" của công
  ty bị tai nạn, nghỉ 3 tháng, quay lại thấy sản phẩm công ty thành "mớ hỗn
  độn" vì ai cũng tự xưng hiểu đúng ý chị).

## Cách dùng khi review

- Dùng các case trong file này để **minh hoạ cụ thể** cho nguyên tắc trừu
  tượng ở file `01`–`05`, không dùng để thay thế nguồn nguyên tắc gốc.
- Khi trích dẫn, nêu rõ: tên bài trên uxlagi.com + case cụ thể trong bài (ví
  dụ: "giống lỗi phân cấp thị giác trong case BestBuy — uxlagi.com, bài
  *Phân cấp thị giác*"), không chỉ nói chung "theo uxlagi".
- File này viết bằng tiếng Việt với nhiều ví dụ đặc thù văn hoá Việt Nam/Đông
  Á (ý nghĩa màu sắc theo văn hoá Á Đông, case bàn phím tiếng Việt, biển phố
  Hà Nội...) — ưu tiên dùng khi review sản phẩm nhắm thị trường Việt Nam/Đông
  Nam Á, nơi các case study gốc tiếng Anh (NN/g, Refactoring UI) có thể thiếu
  góc nhìn văn hoá phù hợp.
