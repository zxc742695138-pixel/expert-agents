# Platform Guidelines chính thức

Nguồn: Apple — [Human Interface Guidelines](https://developer.apple.com/design/human-interface-guidelines)
(developer.apple.com/design); Google — [Material Design 3](https://m3.material.io);
Microsoft — [Fluent 2 Design System](https://fluent2.microsoft.design).

> Nguyên tắc nền tảng chỉ áp dụng khi review đúng nền tảng tương ứng. Không
> chấm điểm một app Android theo chuẩn Apple HIG, và ngược lại. Với web
> thuần, dùng chủ yếu `01`, `02`, `03`, `05`; platform guideline ở đây chỉ áp
> khi web đó cố tình theo native pattern của một hệ điều hành cụ thể (vd.
> web app đóng gói thành iOS app).

## A. Apple Human Interface Guidelines (iOS / iPadOS / macOS)

3 nguyên tắc thiết kế cốt lõi của Apple:

- **Clarity** — Text dễ đọc ở mọi size, icon chính xác dễ hiểu, visual
  hierarchy rõ ràng; mọi yếu tố UI phải phục vụ nội dung, không gây xao nhãng.
- **Deference** — UI nên "lùi lại" để nội dung là trung tâm: dùng
  translucency, blur, bezel/border tối giản để UI không lấn át content.
- **Depth** — Dùng layer thị giác và motion để truyền đạt thứ bậc/trạng thái
  (vd: một item được chọn "nổi" lên trên các item khác).

Quy chuẩn kỹ thuật cần kiểm tra khi review:
- **Touch target tối thiểu 44×44pt**.
- **Dynamic Type** — chữ phải scale theo cài đặt accessibility của hệ thống,
  không hardcode size tuyệt đối.
- **SF Symbols** — ưu tiên icon hệ thống để đồng bộ với platform, thay vì
  custom icon không nhất quán style.
- **Safe area** — layout không được bị che bởi notch/home indicator/status
  bar.
- Tuân thủ **navigation pattern chuẩn**: tab bar ở đáy cho điều hướng chính,
  navigation bar ở trên cho back/action theo context, modal cho tác vụ biệt
  lập ngắn.
- Từ 2025, Apple bổ sung vật liệu **Liquid Glass** (translucent material) vào
  hệ thống thiết kế — khi review app mới trên iOS, kiểm tra việc dùng
  translucency có nhất quán với cập nhật này không, thay vì dùng flat
  opaque layer cũ một cách máy móc.

## B. Material Design 3 (Android / Google / Web theo Material)

- **Dynamic color** — theme màu có thể được sinh tự động từ wallpaper/nội
  dung người dùng (Material You); khi review, kiểm tra hệ màu có dùng đúng
  role token (primary/secondary/tertiary/surface...) thay vì hardcode hex.
- **Tonal elevation** — MD3 dùng lớp màu tonal (surface tint) làm tín hiệu
  elevation chính, không chỉ dựa vào shadow như Material cũ.
- **Adaptive layout** — UI phải thích ứng theo window size class (compact/
  medium/expanded), không chỉ thiết kế cố định cho 1 kích thước màn hình.
- **Touch target tối thiểu 48×48dp** (lớn hơn chuẩn Apple).
- **Motion có mục đích** — chuyển động phải mang tính thông tin (informative),
  có trọng tâm (focused), không chỉ trang trí.

## C. Microsoft Fluent 2 Design System

5 nguyên tắc cốt lõi: **Light, Depth, Motion, Material, Scale**.

- **Light** — dùng ánh sáng để hướng sự chú ý.
- **Depth** — phân lớp UI rõ ràng qua elevation.
- **Motion** — chuyển động mượt, có mục đích, nối tiếp ngữ cảnh.
- **Material** — vật liệu như Acrylic (mờ có blur) và Mica (mờ dựa theme hệ
  thống) tạo chiều sâu và ngữ cảnh desktop.
- **Scale** — UI thích ứng nhất quán trên nhiều kích thước màn hình/thiết bị
  (desktop, web, Xbox...).
