### Kafka streams
- kafka streams giúp xử lí và phân tích dữ liệu luồng từ kafa
- mục tiêu là xử lí dữ liệu realtime với các phép toán đơn giản như filter, map, join, aggregation, windowing
- bản thân kafka streams là một thư viện java
-> ở đây chỉ đề cập đến ý tưởng cốt lõi của nó

1. các thành phần chính trong kafka streams
- stream: đại diện cho luồng dữ liệu liên tục, không ngừng
- table: lưu trữ trạng thái của các tream
- KStream: một luồng dữ liệu không có trạng thái. nó có thể được filter, map, aggregation
- KTable: Một bảng trạng thái, có thể được sử dụng để cập nhật và duy trì trạng thái của luồng dữ liệu

2. Các api chính của kafka streams
- stream api: được sử dụng để xử lí các luồng dữ liệu
- table api: dùng để xử lí và lưu trữ trạng thái của các luồng dữ liệu

3. Các phép toán cơ bản:
- filter: lọc các message theo điều kiện nhất định
- map và flatmap: chuyển đổi các message
- join: kết hợp các luồng dữ liệu từ hai nguồn khác nhau
- aggregration: tổng hợp dữ liệu (tính tổng, trung bình, đếm số lượng) từ các message trong một khoảng thời gian nhất định
- windowing: xử lí dữ liệu theo time windows (time window là một phạm vi thời gian được xác định để nhóm các sự kiện lại với nhau)
