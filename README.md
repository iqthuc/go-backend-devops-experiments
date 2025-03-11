# Project này là để trải nghiệm Go-backend-devops

### cấu trúc project:
<pre><code>sport-shop/
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
│   │       ├── repository.go    //Database access methods for user data
│   │       └── user_case.go     //Core business logic for user operations
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
</code></pre>
