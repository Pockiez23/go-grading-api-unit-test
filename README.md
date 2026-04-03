
``` test
go-grading-api/
├── cmd/
│   └── server/
│       └── main.go              # Entry point / Dependency Injection / Router setup
│
├── internal/
│   ├── auth/                    # Authentication Domain
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   │
│   ├── student/                 # Student Domain
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   │
│   ├── grade/                   # Grade Domain
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   │
│   └── middleware/              # Auth Middleware (JWT)
│       └── auth_middleware.go
│
├── pkg/
│   ├── logger/
│   │   └── logger.go
│   ├── jwt/
│   │   └── jwt.go
│   └── response/
│       └── response.go
│
├── api/
│   └── swagger.yaml             # API documentation
│
├── tests/
│   ├── grade_test.go
│   ├── auth_test.go
│   └── student_test.go
│
├── postman/
│   └── university-api.postman_collection.json
│
├── Makefile
├── go.mod
├── go.sum
└── README.md
```