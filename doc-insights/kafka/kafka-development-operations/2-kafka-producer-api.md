### Kafka Producer API
- Producer trong kafka có nhiệm vụ gửi message đến topic của kafka cluster
-> kafka producer API giúp gửi message từ ứng dụng đến một topic và có thể cấu hình các thuộc tính để kiểm soát cách gửi message

- Các phần cần tìm hiểu:
  - cách tạo một kafka producer
  - các cấu hình quan trọng của producer
    - acks
    - retries
    - batch.size
    - ligner.ms
    - key.serializer và value.serializer
  - gửi message đồng bộ và bất đồng bộ
  - các tính huống sử dụng producer trong thực tế

1. Cách tạo một kafka producer
- sử dụng thư viện của ngôn ngữ tương ứng để tạo

2. Các cấu hình quan trọng của producer
  - acks: xác định cách kafka xử lí xác nhận khi gửi message
    - giá trị của acks:
      - 0: producer không chờ xác nhận từ broker
      - 1: chờ broker leader xác nhận
      - all: chờ tất cả replica xác nhận
  - retries: xác định số lần thử lại khi gửi message thất bại, mặc định là 0
  - batch.size: xác định kích thước gói tin của batch mà producer gửi đi, mặc định là 16KB
  - linger.ms: cho phép producer chờ một thời gian ngắn trước khi gửi batch đi (nếu batch chưa đầy). giúp giảm số lần gửi message đến broker nhưng có thể gây độ trễ nhỏ
  - key.serializer và value.serializer: xác định cách kafka serialzies key và value của message. ví dụ StringSerializer nếu muốn gửi dữ liệu kiểu string

3. Gửi message đồng bộ và bất đồng bộ
- synchronous: producer sẽ đợi đến khi kafka xác nhận message đã được ghi vào topic
- asynchronous: producer gửi message và không chờ xác nhận

4. Các tình huống sử dụng Producer trong ứng dụng thực tế
- ứng dụng gửi log: producer gửi log vào kafka, sau đó các consumer sẽ xử lí log
- ứng dụng xử lí sự kiện: kafka gửi các event đến topic, các dịch vụ khác sẽ sử dụng chúng
- ứng dụng giao dịch tài chính: kafka được sử dụng để ghi lại các giao dịch của người dùng để xử lí và phân tích sau này
