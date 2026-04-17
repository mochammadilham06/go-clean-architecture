# Go Clean Architecture Template

A Golang boilerplate template implementing **Clean Architecture** principles for building scalable, testable, and maintainable RESTful APIs.

## 🚀 Key Features

- **Separation of Concerns**: Clear boundary and separation between Handler, Service, and Repository layers.
- **Dependency Injection**: Automated dependency management using [Google Wire](https://github.com/google/wire).
- **API Documentation**: Auto-generated API documentation with Swagger.
- **Database Migration**: Structured database schema management located in the `db/migration` folder.
- **Containerization**: Deployment-ready with Docker and Docker Compose.

## 📁 Folder Structure

```text
├── cmd/
│   └── main.go           # Application entry point
├── db/
│   └── migration/        # Database migration files (SQL up/down)
├── docs/                 # Auto-generated Swagger documentation
├── server/
│   ├── api/
│   │   ├── contract/     # Interfaces for Services and Repositories
│   │   ├── handler/      # Delivery Layer (Gin Handlers & Routes)
│   │   ├── models/       # Domain Entities / Database Models
│   │   ├── repository/   # Data Access Layer (Database Queries)
│   │   └── service/      # Business Logic Layer
│   ├── lib/              # Internal libraries (Logger, DB Config, Middleware)
│   └── wire.go           # Dependency Injection configuration
├── .env.example          # Template for environment configurations
├── Dockerfile            # Docker image configuration
├── Makefile              # Terminal command shortcuts
└── go.mod                # Dependency manager
```
