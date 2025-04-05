### Kafka connect

1. Kafka connect là gì trong hệ thống Kafka
- kafka connect là công cụ giúp dễ dàng di chuyển dữ liệu giữa kafka và các hệ thống khác như database, hệ thống lưu trữ file, message queue
- nó là cầu nối giữa kafka và các hệ thống ngoài, giúp không phải viết thêm mã tùy chỉnh cho việc chuyển dữ liệu qua lại giữa kafka và các hệ thống này
- cấu hình:
  - source connectors để lấy dữ liệu từ các hệ thống ngoài vào kafka
  - sink connectors để đẩy dữ liệu từ kafka vào các hệ thống ngoài

2. Tác dụng và mô hình làm việc của kafka connect
- giả sử bạn có một ứng dụng hoặc hệ thống cần giao tiếp với Kafka, ví dụ như:
  - đọc dữ liệu từ một cơ sở dữ liệu và đưa vào Kafka.
  - đưa dữ liệu từ Kafka vào một hệ thống file như HDFS hoặc Amazon S3.
  - kết nối với Elasticsearch để index dữ liệu từ Kafka.
->
  - kafka connect sẽ giúp làm điều đó mà không cần viết mã tùy chỉnh, chỉ cần cấu hình các connectors phù hợp
  - nếu không có connect thì phải viết code thực hiện điêu đó

3. Cách hoạt động
3.1. Source connectors
- souce connectors là các thành phần giúp đọc dữ liệu từ các hệ thống bên ngoài và đưa vào kafka
- ví dụ: đồng bộ dữ liệu từ mysql database vào kafka
- mô tả: MySQL Database ---> [JDBC Source Connector] ---> Kafka Topic
3.2.2 Sink connectors
- sink connectors là các thành phần giúp đưa dữ liệu từ kafka vào các hệ thống bên ngoài
- ví dụ: lưu trữ dữ liệu từ kafka vào Amazone S3
- mô tả: Kafka Topic ---> [S3 Sink Connector] ---> Amazon S3
