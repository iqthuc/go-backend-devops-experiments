```
  Login
    ├── Client
    │   └── Gửi thông tin đăng nhập: username/email, password
    └── Server
        ├── Nhận yêu cầu
        ├── Kiểm tra hợp lệ dữ liệu
        └── Truy vấn DB tìm người dùng
            ├── Không tìm thấy: Trả về lỗi: Thông tin đăng nhập không hợp lệ
            └── Tìm thấy:
                ├── So sánh mật khẩu đã nhập với mật khẩu đã hash trong DB
                │   ├── Không khớp: Trả về lỗi: Thông tin đăng nhập không hợp lệ
                │   └── Khớp:
                │       ├── Tạo Access Token và Refresh Token
                │       └── Lưu Refresh Token vào DB (liên kết user, thiết bị)
                └── Trả về Access Token và Refresh Token cho Client
```
# cấu struct token
```json
access token
{
  "user_id": 123,
  "exp": 1679875200,
  "type": "access",
  "username": "example_user",
  "roles": ["user"]
}
```

```json
refresh token
{
  "user_id": 123,
  "exp": 1680480000,
  "type": "refresh",
  "device_id": "unique_device_identifier"
}
```
