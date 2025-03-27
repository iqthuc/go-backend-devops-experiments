```
  Logout
    ├── Client
    │   └── Gửi request logout đến endpoint /api/logout
    ├── Server
    │    ├── Nhận yêu cầu
    │    ├── Tìm và xóa Refresh Token tương ứng trong DB
    │    └── Trả về response thành công cho Client
    └── Client
        └── Xóa Access Token và Refresh Token đã lưu trữ
```
