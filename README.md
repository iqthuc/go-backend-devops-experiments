# Project này là để trải nghiệm Go-Backend-DevOps

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
