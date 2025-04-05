### Tạo, xóa, gửi tin nhắn đến topic
- tạo một topic:
```sh
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 3 --topic my-topic
```
      --replication-factor 1: mỗi partition chỉ có một replication (bao gồm chính nó)
      --partitions 3: tạo 3 partitions cho topic này
- thay đổi cấu hình một topic:
```sh
kafka-topics --bootstrap-server localhost:9092 --alter --topic <topic-name> --partitions <new-total-partition-count>
```

- xóa một topic:
```sh
kafka-topics --bootstrap-server localhost:9092 --delete --topic my-topic
```
    note: cần bật cấu hình cho phép xóa trong file 'server.properties': delete.topic.enable=true

- liệt kê tất cả các topic đang có:
```sh
kafka-topics --bootstrap-server localhost:9092 --list
```

- Xem thông tin một topic:
```sh
kafka-topics --bootstrap-server localhost:9092 --describe --topic my-topic
```
    - kết quả mẫu:
    Topic: my-topic  TopicId: ... PartitionCount: 3 ReplicationFactor: 1 Configs: ...
    Topic: test-topic  Partition: 0  Leader: 1  Replicas: 1,2  Isr: 1,2
    Topic: test-topic  Partition: 1  Leader: 2  Replicas: 2,3  Isr: 2,3
    Topic: test-topic  Partition: 2  Leader: 3  Replicas: 3,1  Isr: 3,1
    - chú thích
        - PartitionCount: 3 -> có 3 partitions
        - Leader -> Broker đang làm leader cho partition đó
        - Replicas -> Danh sách broker có replication của partition tương ứng (bao gồm cả partiton leader)
        - Isr (In-sync replicas) -> các replica đang đồng bộ với leader

=> Tóm tắt lệnh:
- List topic:       kafka-topics --bootstrap-server ... --list
- Describe topic:   kafka-topics --bootstrap-server ... --describe --topic ...
- Create topic:     kafka-topics --bootstrap-server ... --create ...
- Delete topic:     kafka-topics --bootstrap-server ... --delete --topic ...
- Alter partitions:	kafka-topics --bootstrap-server ... --alter --topic ...

### Gửi message đến topic
  - gửi không kèm key:
```sh
kafka-console-producer --bootstrap-server localhost:9092 --topic my-topic
```
  - gửi kèm key:
```sh
kafka-console-producer --bootstrap-server localhost:9092 --topic my-topic --property "parse.key=true" --property "key.separator=:"
```
    - message chứa key-value được phân tách bởi dấu 2 chấm, ví dụ "user1: hello"
    - note:
        -> với các message không có key:
            - trong thực tế các message sẽ luôn phiên phân phối vào các partition
            - tuy nhiên với lệnh kafka-console-producer, kafka sẽ tạo một producer tạm thời chỉ dùng trong phiên làm việc này, và chọn ngẫu nhiên một partition mặc định, sau đó gửi tất cả message trong phiên làm việc này vào partitiion đó
        -> với các message có key:
            - partition = hash(key) % number_of_partitions
- đọc message
```sh
kafka-console-consumer --bootstrap-server localhost:9092 --topic my-topic --partition 0 --from-beginning
```
    --partition 0: đọc ở partition 0
    --from-beginning: đọc từ offset đầu tiên của partition, không chỉ định thì chỉ đọc message mới


### Consumer Group & Message Delivery Semantics
- mục tiêu:
    - hiểu cơ chế Consumer group
    - phân biệt consumer độc lập và consumer trong nhóm
    - hiểu các khái niệm: at-most-once, at-least-one, exactly-one

- Consumer group:
  - là tập hợp các consumer cùng đọc message từ một topic
  - mỗi partition của topic chỉ được xử lí bởi một một consumer trong group tại một thời điểm
  - kafka tự động phân chia partition cho consumer, gọi là rebalancing
  - => mục tiêu: scale việc tiêu thụ message và tránh trùng lặp xử lí

  - tạo consumer với group:
```sh
kafka-console-consumer --bootstrap-server localhost:9092 --topic my-topic --group my-Group
kafka-console-consumer --bootstrap-server localhost:9092 --topic my-topic --group my-Group
```
  - kiểm tra group:
```sh
kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group my-group
```
  - xem danh sác group đang tồn tại:
```sh
kafka-consumer-groups --bootstrap-server localhost:9092 --list
```
  - xem chi tiết group nào đang đọc topic nào, ở partitio nào:
```sh
kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group my-group
```
  - Xóa offset của một group cho một topic:
```sh
kafka-consumer-groups --bootstrap-server localhost:9092 --group <group-name> --topic <topic-name> --reset-offsets --to-earliest --execute
```
    - chú thích:
      - có thể thay --to-earliest bằng:
        --to-latest: nhảy tới offset mới nhất
        --shift-by-10: lùi lại 10 offset
        --to-offset 123: tới offset 123
      - dùng --dry-run dể xem kết quả mà không thực thi

=> Tóm tắt lệnh:
- List all groups:  kafka-consumer-groups --bootstrap-server ... --list
- Describe group:   kafka-consumer-groups --bootstrap-server ... --describe --group my-group
- Delete group:     kafka-consumer-groups --bootstrap-server ... --delete --group my-group
- Reset offset:     ... --reset-offsets --to-earliest --execute


### Offset
  - offset là chỉ số đại diện cho vị trí của một message trong một partition
  - kafka chia topic ra thành cac partition. trong mỗi partition, message được sắp xếp theo thứ tự ghi vào, mà mỗi message dược đánh một số số tăng dần, gọi là offset
  - Offset dùng làm gì:
    - consumer dùng offset để:
      - biết nên đọc message nào kế tiếp
      - kafka không xóa message sau khi gửi, mà sẽ lưu lại một khoảng thời gian được chỉ định
      - consumer có thể quyết định offset nào cần đọc tiếp

### Lag
  - lag là độ trễ giữa vị trí mới nhất của message (latest offset) trên partition và vị trí offset hiện tại của consumer đang đọc đến
  -> lag = partition có nhiều message mới, nhưng consumer của bạn vẫn chưa đọc tới (số message chưa đọc đó là giá trị lag)
  - ý nghĩa:
    - để theo dõi hiệu suất xử lí của consumer:
    - nếu lag tăng liên tục, có thể là:
      - consumer xử lí message chậm
      - partition có quá nhiều message gửi tới
      - consumer bị treo hoặc mất kết nối
  - lưu ý:
    - lag = 0 là lí tưởng, nhưng vài chục/vài trăm là bình thường
    - lag tăng dần là dấu hiệu cần scale thêm consumer hoặc xử lí message nhanh hơn
  - kiểm tra lag thực tế:
```sh
kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group my-group
```
