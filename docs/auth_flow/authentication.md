└── Authentication
    ├── Client
    │   └── Gửi request đến endpoint bảo vệ (kèm Access Token trong Header)
    └── Server
        ├── Nhận yêu cầu
        ├── Kiểm tra Access Token trong Header
        ├── Xác thực Access Token (chữ ký, hết hạn, claims)
        │   ├── Hợp lệ: Xử lý request và trả về dữ liệu
        │   └── Không hợp lệ: Trả về lỗi: 401 Unauthorized
