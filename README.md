# Task Manager

Task Manager is a distributed task management system built with Go, PostgreSQL, Redis, and Google Cloud Infrastructure. It features user authentication, task management, real-time notifications, caching with Redis, report generation, and a robust CI/CD pipeline using GitHub Actions. The application is designed for high availability and scalability, leveraging modern backend technologies and best practices.

## Features

### 1. User Authentication

- **JWT-Based Authentication:** Secure user authentication using JSON Web Tokens.
- **Endpoints:**
  - `POST /register`: Register a new user.
  - `POST /login`: Authenticate a user and return a JWT token.

### 2. Task Management

- **CRUD Operations:** Create, read, update, and delete tasks.
- **Endpoints:**
  - `POST /tasks`: Create a new task.
  - `GET /tasks/:id`: Retrieve a task by ID.
  - `PUT /tasks/:id`: Update a task by ID.
  - `DELETE /tasks/:id`: Delete a task by ID.

### 3. Real-time Notifications

- **WebSocket Notifications:** Real-time updates using WebSockets for tasks.
- **Implementation:**
  - Uses `github.com/gorilla/websocket` for WebSocket communication.
  - Clients can connect to `/ws` to receive notifications about task updates.

### 4. Caching with Redis

- **Redis Caching:** Improve performance by caching task details.
- **Implementation:**
  - Uses `github.com/go-redis/redis/v8` for Redis integration.
  - Cached data is stored for 10 minutes to reduce database load.

### 5. Report Generation

- **Generate Reports:** Create reports based on task data.
- **Endpoints:**
  - `GET /report`: Generate a report summarizing task data.

### 6. Continuous Integration and Deployment

- **GitHub Actions:** Automated CI/CD pipeline.
- **Features:**
  - Runs tests on every push to the `main` branch.
  - Builds Docker images and pushes to Docker Hub.
  - Deploys the application to Google Cloud Run.

### 7. Docker

- **Dockerized Application:** Containerized application for easy deployment and scalability.
- **Docker Compose:** Multi-container Docker applications for managing PostgreSQL, Redis, and the Go application.

### 8. Testing

- **Unit and Integration Tests:** Ensure code quality with automated tests.
- **Implementation:**
  - Uses the `testing` package for writing tests.
  - Includes tests for task creation and retrieval.

### 9. API Documentation

- **Swagger:** Comprehensive API documentation.
- **Implementation:**
  - Uses `swaggo/gin-swagger` for generating Swagger documentation.
  - Accessible at `/swagger/index.html`.

## Getting Started

### Prerequisites

- Go 1.18 or later
- Docker
- Docker Compose
- PostgreSQL
- Redis
- Google Cloud SDK
- GitHub Account

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/meetamjadsaeed/distributed-task-manager-with-redis-CI-CD-WEBSOCKET.git
   cd task-manager
   ```

2. **Set up environment variables:**
   Create a `.env` file and add the following variables:
   ```plaintext
   POSTGRES_USER=yourusername
   POSTGRES_PASSWORD=yourpassword
   POSTGRES_DB=task_manager
   POSTGRES_HOST=db
   REDIS_ADDR=redis:6379
   JWT_SECRET=yourjwtsecret
   ```

### Running the Application

To run the application using Docker Compose:

```bash
docker-compose up --build
```

### Running Tests

To run the unit and integration tests:

```bash
go test ./...
```

### CI/CD Pipeline

The CI/CD pipeline is configured using GitHub Actions. It runs tests, builds Docker
