1. Kafka Producer API
2. Kafka Consumer API
3. Kafka Streams
- Giới thiệu về Kafka Streams API.
- Cách xử lý dữ liệu theo thời gian thực bằng Kafka Streams.
- Các phép toán như map, filter, reduce trên Kafka Streams.
4. Kafka Connect
- Cách sử dụng Kafka Connect để kết nối với các hệ thống bên ngoài (database, file systems, v.v.).
- Cài đặt và cấu hình các connectors.
5. Quản lý Kafka Cluster
- Cách theo dõi và quản lý Kafka cluster.
- Kafka monitoring (JMX, Prometheus, v.v.).
- Các công cụ hỗ trợ quản lý cluster (Kafka Manager, Confluent Control Center).
6. Kafka security
- các tính năng bảo mật trong kafka: authentication, authorization, encrytion
- cấu hfinh bảo mật trong kafka:
  - cấu hình SASL và SSL cho authentication và encryption
  - cấu hình ACLs để phần quyền cho các client và consumer
- các kịch bản bảo mật thực tế:
  - cấu hình Kafka để chỉ cho phép một nhóm client xác thực qua SSL.
  - cấu hình quyền truy cập cho các producer và consumer dựa trên các ACLs.
7. Kafka for Big Data Applications
- Kafka và Hadoop
- Kafka và Apache Spark
- Kafka và Apache Flink:
