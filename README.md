### Mục đích
  - Project này là để trải nghiệm Go-Backend-DevOps.
  - Mục đích là để học lí thuyết về các công nghệ
    và tự triển khai chúng trước khi dùng library trong dự án thực tế
    vì vậy các core business logic trông hơi đần

### Các mục chính
- Go
- sh script
- database: postgres, mongoDB, Redis, Elasticsearch
- api: rest, graphql, gRPC
- authentication: jwt, paseto, OAuth
- message brokers: kafka
- ci/cd: github actions, jenkins
- devops: docker, kubernetes

### cấu trúc project:
```sport-shop/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── delivery/
│   │   ├── middleware/
│   │   │   └── demo.go
│   │   └── router/
│   │       └── demo.go
│   ├── features/
│   │   └── demo/
│   │       ├── handler.go
│   │       ├── models.go
│   │       ├── repository.go    //Database access methods
│   │       └── user_case.go     //Core business logic
│   └── initializer/
│       ├── init_demo.go
│       └── run.go
│
├── pkg/
│   ├── database/
│   │   └── postgres.go
│   ├── logger/
│   │   └── logger.go
│   └── validator/
│       └── validator.go
├── config/
│   └── config.go
├── scripts/
├── migrations/ # Database migrations
├── docs/
├── test/
├── .env
├── go.mod
└── go.sum
```
