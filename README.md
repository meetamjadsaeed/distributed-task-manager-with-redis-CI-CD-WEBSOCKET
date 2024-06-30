# Task Manager

Task Manager is a distributed task management system built with Go, PostgreSQL, and Google Cloud Infrastructure. It features real-time notifications, caching with Redis, and a robust CI/CD pipeline using GitHub Actions. The application is designed for high availability and scalability, leveraging modern backend technologies and best practices.

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

### 5. Continuous Integration and Deployment

- **GitHub Actions:** Automated CI/CD pipeline.
- **Features:**
  - Runs tests on every push to the `main` branch.
  - Builds Docker images and pushes to Docker Hub.
  - Deploys the application to Google Cloud Run.

### 6. Testing

- **Unit and Integration Tests:** Ensure code quality with automated tests.
- **Implementation:**
  - Uses the `testing` package for writing tests.
  - Includes tests for task creation and retrieval.

### 7. API Documentation

- **Swagger:** Comprehensive API documentation.
- **Implementation:**
  - Uses `swaggo/gin-swagger` for generating Swagger documentation.
  - Accessible at `/swagger/index.html`.

## Getting Started

### Prerequisites

- Go 1.18 or later
- Docker
- PostgreSQL
- Redis
- Google Cloud SDK
- GitHub Account

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/task-manager.git
   cd task-manager
   ```

2. **Set up environment variables:**
   Create a `.env` file and add the following variables:

   ```plaintext
   POSTGRES_USER=yourusername
   POSTGRES_PASSWORD=yourpassword
   POSTGRES_DB=task_manager
   POSTGRES_HOST=localhost
   REDIS_ADDR=localhost:6379
   JWT_SECRET=yourjwtsecret
   ```

3. **Install dependencies:**

   ```bash
   go mod download
   ```

4. **Run the application:**
   ```bash
   go run cmd/server/main.go
   ```

### Running Tests

To run the unit and integration tests:

```bash
go test ./...
```

### Using Docker

1. **Build the Docker image:**

   ```bash
   docker build -t yourusername/task-manager:latest .
   ```

2. **Run the Docker container:**
   ```bash
   docker run -p 8080:8080 --env-file .env yourusername/task-manager:latest
   ```

### CI/CD Pipeline

The CI/CD pipeline is configured using GitHub Actions. It runs tests, builds Docker images, and deploys the application to Google Cloud Run.

To set up the pipeline:

1. **Store secrets in GitHub:**

   - `DOCKER_USERNAME`: Your Docker Hub username.
   - `DOCKER_PASSWORD`: Your Docker Hub password.
   - `GCP_PROJECT`: Your Google Cloud project ID.
   - `GCP_KEY_FILE`: Your base64-encoded Google Cloud service account key.

2. **Create the GitHub Actions workflow file:**
   `.github/workflows/ci-cd.yml` (already included in the repository).

### API Documentation

API documentation is available at `/swagger/index.html`. It provides details about all available endpoints, request/response formats, and example payloads.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact
