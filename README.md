# Task API - REST API Learning Project

A RESTful API for managing tasks, built with Go. This is my first Go project, created to learn REST API fundamentals, clean architecture, and Go best practices.

## 🎯 Learning Goals Accomplished

- ✅ Building REST APIs with Go's standard library
- ✅ Clean architecture with separation of concerns
- ✅ Repository pattern for data access
- ✅ HTTP middleware (CORS, logging, recovery, validation)
- ✅ Error handling and input validation
- ✅ Thread-safe concurrent operations
- ✅ Environment-based configuration

## 🏗️ Project Structure

```
task-api/
├── cmd/server/          # Application entry point
├── internal/            # Private application code
│   ├── models/         # Domain models & validation
│   ├── repository/     # Data access layer (in-memory)
│   ├── handlers/       # HTTP request handlers
│   └── middleware/     # HTTP middleware chain
└── .env                # Environment configuration
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.5 or higher

### Installation

```bash
git clone https://github.com/danyaljs/task-api.git
cd task-api
go mod download
```

### Configuration

Copy the example environment file and configure as needed:

```bash
cp .env.example .env
```

Environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT`   | Server port | `8080`  |

### Running the Server

```bash
# Using default port (8080)
go run cmd/server/main.go

# Or with custom port
PORT=3000 go run cmd/server/main.go
```

Server will start on `http://localhost:8080`

## 📡 API Endpoints

| Method | Endpoint      | Description      |
|--------|---------------|------------------|
| GET    | `/tasks`      | Get all tasks    |
| POST   | `/tasks`      | Create new task  |
| GET    | `/tasks/{id}` | Get task by ID   |
| PUT    | `/tasks/{id}` | Update task      |
| DELETE | `/tasks/{id}` | Delete task      |

## 💡 Example Usage

### Create a task

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Build a REST API"
  }'
```

**Response:**
```json
{
  "message": "task created successfully",
  "task": {
    "id": 1,
    "title": "Learn Go",
    "description": "Build a REST API",
    "completed": false,
    "created_at": "2024-12-06T10:00:00Z",
    "updated_at": "2024-12-06T10:00:00Z"
  }
}
```

### Get all tasks

```bash
curl http://localhost:8080/tasks
```

**Response:**
```json
{
  "message": "Tasks retrieved successfully",
  "tasks": [
    {
      "id": 1,
      "title": "Learn Go",
      "description": "Build a REST API",
      "completed": false,
      "created_at": "2024-12-06T10:00:00Z",
      "updated_at": "2024-12-06T10:00:00Z"
    }
  ],
  "count": 1
}
```

### Get a specific task

```bash
curl http://localhost:8080/tasks/1
```

### Update a task

```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go - Updated",
    "completed": true
  }'
```

### Delete a task

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## 🛠️ Technical Features

### Middleware Chain
- **Logger**: Request logging with method, path, status code, duration
- **CORS**: Cross-origin resource sharing support
- **Recovery**: Panic recovery to prevent crashes
- **Content-Type Validation**: Ensures JSON for POST/PUT requests

### Data Validation
- Title: Required, max 200 characters
- Description: Optional, max 1000 characters
- Automatic whitespace trimming

### Thread Safety
- Concurrent-safe in-memory storage using `sync.RWMutex`
- Safe for multiple simultaneous requests

## 🧠 What I Learned

Through building this project, I gained hands-on experience with:

- **Go fundamentals**: Structs, interfaces, error handling, pointers
- **HTTP server**: Routing, handlers, middleware patterns
- **Concurrent programming**: Mutexes, race condition prevention
- **Clean architecture**: Separation of concerns, dependency injection
- **API design**: RESTful principles, proper HTTP status codes
- **Request/Response handling**: JSON encoding/decoding, validation
- **Environment configuration**: Using .env files and environment variables

## 🔮 Next Steps (Future Projects)

This was intentionally kept simple to focus on fundamentals. For my next iteration, I plan to build a more advanced project with:

- [ ] Database integration (PostgreSQL)
- [ ] Unit and integration tests
- [ ] JWT authentication & authorization
- [ ] Structured logging
- [ ] Graceful shutdown handling
- [ ] Docker containerization
- [ ] API versioning (`/api/v1`)
- [ ] OpenAPI/Swagger documentation
- [ ] CI/CD pipeline (GitHub Actions)
- [ ] Request rate limiting
- [ ] Pagination and filtering

## ⚠️ Important Note

This is a **learning project** built to understand Go and REST API fundamentals. It uses in-memory storage and is **not production-ready**. Data is lost when the server restarts. 

For production use, you would need:
- Persistent database storage
- Authentication and authorization
- Comprehensive test coverage
- Monitoring and observability
- Security hardening
- And much more...

Check out my other projects to see my progression! 🚀

## 📄 License

MIT

## 👤 Author

**Danyal** - [@danyaljs](https://github.com/danyaljs)

---

*Built with ❤️ while learning Go*
