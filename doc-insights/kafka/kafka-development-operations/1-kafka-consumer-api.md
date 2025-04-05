### Kafka Producer API
- kafka consumer api cho phép client nhận các message từ Kafka broker
- mỗi consumer có thể đọc từ một hoặc nhiều topic và partition, và có thể tổ chức thành một consumer group
- trong một consumer group, mõi consumer chỉ đọc một phần của các partition, giúp phân chia công việc giữa các consumer để đạt hiệu quả cao hơn

- Các phần cần tìm hiểu:
  - các thành phần của Kafka consumer API
  - các cấu hình quan trọng của Consumer
  - offset và commit
  - consumer groups

1. Các thành phần của Kafka consumer api
- Consumer Record: đây là một object chứa thông tin về message mà consumer nhận được, bao gồm thông tin về topic, partition, offset, key, value, ...
- Consumer Config: các tùy chọn dùng để cấu hình consumer, ví dụ như group.id, auto.offset.reset
- Consumer: đây là đối tượng chính mà bạn dùng để nhận các message từ Kafka

2. Các cấu hình quan trọng
- group.id: id của consumer group, nếu nhiều consumer có cùng group.id, chúng sẽ chia sẽ việc đọc từ các partition của topic
- auto.offset.reset: hành vi của consumer khi không tìm thấy offset đã lưu:
  - earliest: bắt đầu từ đầu (offset 0)
  - latest: bắt đầu từ cuối (offset mới nhất)
  - none: nếu không có offset, sẽ có lỗi
- enable.auto.commit: consumer sẽ tự động commit offset sau khi nhận message

3. Offset và Commit
- kafka sử dụng offset để theo dõi vị trí của các message trong một partition
- mỗi message trong một partition đều có một offset duy nhất, giúp consumer biết được đã tiêu thụ đến đâu
- consumer có thể tự động hoặc thủ công commit offset để kafka biết message đã được thiêu thụ và không gửi lại cho consumer

4. Nhận message đồng bộ và bất đồng bộ
- synchronous: conumser sẽ xử lí xong message trước khi tiếp tục nhận message mới
- asynchronous: consumer sẽ nhận và xử lí nhiều message cùng lúc mà không cần xử lí xong message hiện tại

5. Consumer groups
- consumer groups là để phân tán việc tiêu thụ message từ các partition của Kafka
- khi có nhiều consumer trong một group, kafka đảm bảo rằng mỗi partition chỉ được một consumer trong group tiêu thụ tại một thời điểm
  ->
    - các consumer trong cùng một group sẽ chia sẻ các partition của topic, và mỗi consumer chỉ đọc từ một partition duy nhất
    - nếu có nhiều consumer hơn partition, một số consumer sẽ không nhận được message
    - nếu có ít consumer hơn partition, một số consumer sẽ đọc từ nhiều partition

6. Các tình huống sử dụng consumer trong thực tế
- xử lí dữ liệu realtime: nhận dữ liệu và xử lý realtime trong các ứng dụng phân tán. ví dụ ứng dụng e-commerce có thể dùng consumer để nhận thông báo về các giao dịch mới và xử lí chúng (cập nhật kho hàng, gửi email)
- phân tích log: nhận log từ các service và phân tích chúng để giám sát hoặc trích xuất thông tin quan trọng
