sqlc là tool giúp generate code Go từ truy vấn SQL

Các bước thực hiện
- viết sql
- chạy lệnh generate để tạo mã Go
- dùng các mã Go được tạo

### Cấu hình áp dụng vào project
Bước 1: khai báo file sqlc.yml (hoặc sqlc.yaml, sqlc.json)
```yaml
version: "2"
sql:
  - engine: "postgres"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "demo-sqlc"
        out: "demo-sqlc"
        sql_package: "pgx/v5"
```
    - note:
      - engine: engine của database sử sẽ dùng
      - queries: là file/thư mục định nghĩa các query sql
      - schema: là file/thưc mục định nghĩa các schema
      - gen: cấu hình quá trình gen
        - package: tên package của các file go sẽ gen ra
        - out: vị trí của các file sẽ gen
        - sql_package: package muốn dùng (mặc định sẽ dùng lib/pq)

Bước 2: Khai báo schema và queries
- khai báo file schema.sql (hoặc tên bất kì) chứa thông tin về schema của database
```sql
CREATE TABLE authors (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name text    NOT NULL,
  bio  text
);
```
- khai báo các queries
```sql
-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;
```
    - giải thích:
      --name: CreateUser :one -> tên của hàm Go được tạo sẽ là CreateUser và hàm sẽ return một đối tượng User tương ứng (tất nhiên là vẫn return error). các tùy chọn khác là many (trả về một list), exec (không  trả về cái gì), execresult

Bước 3: Generating code
```sh
sqlc generate
```
    - note:
      - mỗi file .sql sẽ tạo ra một file go tương ứng, trong đó sẽ có các hàm tương ứng với sql đã tạo, mỗi hàm sẽ tự thực hiện việc query tới database và trả về dữ liệu (không cần phải scan thủ công)
      - mỗi hàm đều là một query riêng biệt, nếu cần xử lí transaction cần tự kết hợp

Bước 4: sử dụng mã được generated
