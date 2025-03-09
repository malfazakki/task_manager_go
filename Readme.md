# Task Manager API

A simple **Task Manager API** built with **Golang**, using the **Echo framework** and designed following the **SOLID principles**.

## Features
- User authentication (Signup/Login)
- Task management (Create, Read, Update, Delete)
- Middleware for authentication and logging
- Database integration with PostgreSQL using GORM
- Structuring based on SOLID principles

## Tech Stack
- **Golang**
- **Echo Framework**
- **PostgreSQL**
- **GORM (ORM for Golang)**
- **JWT Authentication**

## Project Structure
```
task-manager/
├── cmd/               # Entry point aplikasi
├── config/            # Konfigurasi database, environment
├── internal/          # Folder utama kode aplikasi
│   ├── domain/        # Interface untuk abstraksi
│   ├── models/        # Struct untuk database
│   ├── repository/    # Layer repository (database access)
│   ├── usecase/       # Business logic (use case)
│   ├── handler/       # Controller untuk HTTP request
│   └── middleware/    # Middleware seperti JWT Auth
├── migrations/        # Skrip migrasi database
├── go.mod             # Dependencies
└── main.go            # Entry point utama
```

## Installation
### 1. Clone the repository
```sh
git clone https://github.com/malfazakki/task_manager_go.git
cd task_manager_go
```

### 2. Install dependencies
```sh
go mod tidy
```

### 3. Configure environment variables
Create a `.env` file and set up the required environment variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASS=your_password
DB_NAME=task_manager
JWT_SECRET=your_jwt_secret
```

### 4. Run database migrations (if applicable)
```sh
go run migrations/migrate.go
```

### 5. Start the server
```sh
go run main.go
```

## API Endpoints
### Authentication
| Method | Endpoint      | Description        |
|--------|-------------|--------------------|
| POST   | `/register`    | Register a user   |
| POST   | `/login`     | Authenticate user |

### Task Management
| Method | Endpoint        | Description             |
|--------|---------------|-------------------------|
| POST   | `/tasks`       | Create a new task       |
| GET    | `/tasks`       | Get all tasks           |
| GET    | `/tasks/:id`   | Get task by ID          |
| PUT    | `/tasks/:id`   | Update task by ID       |
| DELETE | `/tasks/:id`   | Delete task by ID       |

## SOLID Principle Implementation
- **S**: **Single Responsibility Principle** → Each package (handlers, services, repository) has a distinct responsibility.
- **O**: **Open/Closed Principle** → The service layer can be extended without modifying existing code.
- **L**: **Liskov Substitution Principle** → The repository interface allows switching database implementations.
- **I**: **Interface Segregation Principle** → Separate interfaces for different functionalities (e.g., UserRepository, TaskRepository).
- **D**: **Dependency Inversion Principle** → The handlers depend on abstractions (interfaces) instead of concrete implementations.

