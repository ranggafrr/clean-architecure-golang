<div align="center">

# 🚀 Clean Architecture Golang Template

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=flat&logo=mysql&logoColor=white)](https://www.mysql.com)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**Production-ready REST API template dengan Clean Architecture pattern**

[Features](#-features) • [Quick Start](#-quick-start) • [API Documentation](#-api-documentation) • [Architecture](#-architecture)

</div>

---

## ✨ Features

- 🏗️ **Clean Architecture** - Separation of concerns dengan layer yang jelas
- 🔒 **Type Safe** - Strongly typed dengan Go
- 🗄️ **MySQL Integration** - Database connection pooling & optimized queries
- ✅ **Input Validation** - Request validation di setiap layer
- 🎯 **Error Handling** - Custom error types & consistent error responses
- 📝 **Structured Logging** - Easy debugging & monitoring
- 🔧 **Environment Config** - Flexible configuration management
- 🧪 **Testable** - Easy to mock & unit test
- 📦 **Zero Dependencies** - Minimal external packages (hanya MySQL driver)

## 📁 Project Structure

```
clean-architecture/
├── cmd/
│   └── main.go                 # 🚪 Application entry point
├── config/
│   └── config.go               # ⚙️ Configuration management
├── internal/
│   ├── domain/                 # 🎯 Business entities & interfaces
│   │   ├── user.go            # User entity & contracts
│   │   └── errors.go          # Domain errors
│   ├── usecase/                # 💼 Business logic layer
│   │   └── user_usecase.go
│   ├── repository/             # 💾 Data access layer
│   │   └── user_repository.go
│   └── delivery/http/          # 🌐 HTTP handlers & routing
│       ├── user_handler.go
│       └── router.go
├── pkg/
│   ├── database/               # 🔌 Database connection
│   │   └── mysql.go
│   └── response/               # 📤 Response helpers
│       └── response.go
├── docs/
│   └── database.sql            # 📊 Database schema
├── .env.example                # 📝 Environment template
├── .gitignore
├── Makefile                    # 🛠️ Build commands
└── go.mod
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- MySQL 8.0 or higher
- Make (optional)

### Installation

1️⃣ **Clone repository**
```bash
git clone https://github.com/ranggafrr/clean-architecture.git
cd clean-architecture
```

2️⃣ **Setup database**
```bash
mysql -u root -p < docs/database.sql
```

3️⃣ **Configure environment**
```bash
cp .env.example .env
# Edit .env dengan konfigurasi database Anda
```

4️⃣ **Install dependencies**
```bash
go mod download
```

5️⃣ **Run application**
```bash
go run cmd/main.go
```

✅ Server running at `http://localhost:8080`

### Using Makefile

```bash
make deps      # Install dependencies
make migrate   # Run database migration
make run       # Run application
make build     # Build binary
```

## 📚 API Documentation

### Endpoints

| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| `GET` | `/api/users` | Get all users | ✅ |
| `GET` | `/api/users/detail?id=1` | Get user by ID | ✅ |
| `POST` | `/api/users` | Create new user | ✅ |
| `PUT` | `/api/users/detail?id=1` | Update user | ✅ |
| `DELETE` | `/api/users/detail?id=1` | Delete user | ✅ |

### Request Examples

#### Create User
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

**Response (201 Created)**
```json
{
  "success": true,
  "message": "user created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

#### Get All Users
```bash
curl http://localhost:8080/api/users
```

**Response (200 OK)**
```json
{
  "success": true,
  "message": "users retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ]
}
```

#### Get User by ID
```bash
curl http://localhost:8080/api/users/detail?id=1
```

#### Update User
```bash
curl -X PUT http://localhost:8080/api/users/detail?id=1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com"
  }'
```

#### Delete User
```bash
curl -X DELETE http://localhost:8080/api/users/detail?id=1
```

### Error Responses

**400 Bad Request**
```json
{
  "success": false,
  "message": "name is required"
}
```

**404 Not Found**
```json
{
  "success": false,
  "message": "user not found"
}
```

**500 Internal Server Error**
```json
{
  "success": false,
  "message": "internal server error"
}
```

## 🏗️ Architecture

### Clean Architecture Layers

```
┌─────────────────────────────────────────┐
│          HTTP Handler Layer             │  ← Delivery
│  (Routing, Request/Response handling)   │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│         Use Case Layer                  │  ← Business Logic
│    (Business rules & validation)        │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│        Repository Layer                 │  ← Data Access
│     (Database operations)               │
└──────────────────┬──────────────────────┘
                   │
┌──────────────────▼──────────────────────┐
│           Domain Layer                  │  ← Core
│   (Entities, Interfaces, Errors)        │
└─────────────────────────────────────────┘
```

### Dependency Rule

- **Domain** tidak bergantung pada layer manapun
- **Use Case** hanya bergantung pada Domain
- **Repository** hanya bergantung pada Domain
- **Handler** bergantung pada Domain & Use Case

### Benefits

✅ **Independent of Frameworks** - Business logic tidak terikat framework
✅ **Testable** - Mudah di-mock dan di-test
✅ **Independent of UI** - Bisa ganti HTTP ke gRPC tanpa ubah logic
✅ **Independent of Database** - Ganti MySQL ke PostgreSQL hanya ubah repository
✅ **Independent of External Services** - Business rules tidak tahu external services

## 🔧 Configuration

Edit `.env` file:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=your_db_name

# Server Configuration
SERVER_PORT=8080
```
## 📦 Build

```bash
# Development
go run cmd/main.go

# Production
go build -o bin/app cmd/main.go
./bin/app
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by Uncle Bob
- Go community for amazing tools and libraries

---

<div align="center">

**⭐ Star this repo if you find it helpful!**

Made with ❤️ by Rangga Firmansyah (https://github.com/ranggafrr)

</div>
