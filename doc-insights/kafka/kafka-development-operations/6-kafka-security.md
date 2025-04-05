### Kafka security
- kafka security bao gồm các phương thức để đảm bảo dữ liệu được bảo vệ, các kết nối là an toàn, và chỉ những người có quyền mới có thể giao tiếp với kafka cluter

1. Authentication
2. Authorization
3. Encryption
4. Autiting
- ghi lại các sự kiện bảo mật cho mục đích kiểm tra, giúp theo dõi các hoạt động như:
  - người dùng nào đã thực hiện hành động nào
  - ai đã gửi/nhận message từ topic nào
  - kiểm tra các hành động bất thường
- công cụ Audit logs trong kafka sẽ hỗ trợ ghi lại các sự kiện và hành động liên quan đến bảo mật
