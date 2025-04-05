### Kafka CLuster Management
- việc theo dõi kafka cluster giúp phát hiện và xử lí các sự cố tiềm ẩn

1. Các yếu tố cần theo dõi
- broker health:
  - các chỉ số như: uptime, broker status, under-replicated partitions, disk usage
  - cần theo dõi số lượng các broker trong trạng thái alive đảm bảo rằng các broker bị lỗi được xử lí kịp thời
- tình trạng của các partitions:
  - các partition có thể bị thiểu replica hoặc các relica có thể không đồng bộ với leader
  - kiểm tra ISR (In-Sync Replicas) và under-replicated partitions
- producer và consumer performance:
  - theo dõi số lượng, tốc độ gửi/nhận mesage, thời gian phản hồi của producer và consumer
  - kiểm tra lag của consumer
- thực thi các yêu cầu metada:
  - kafka phải luôn duy trì thông tin metadata về các topics, partitions, broker
- kiểm tra message trong mỗi partition:
  - kiểm tra số lượng message bị bỏ qua, không tiêu thụ, hoặc các phân vùng bị tắc nghẽn

2. Kafka monitoring
- các công cụ giảm sát và cảnh báo, thu thập  chỉ số của kafka như: throughput consumer lag, broker status
- JMX: viết bằng java -> không thích :3
- Prometheus & Grafana

3. các công cụ hỗ trợ quản lí kafka cluster:
- quản lí, giám sát broker, topic, partitions, consumer groups
- kafka manager
- confluent control center
- kafka cruise control
