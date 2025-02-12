# Task Manager

Task Manager is a backend service for managing tasks with real-time notifications and reporting functionality.

## Features

- **User Authentication**: Register and login endpoints with JWT-based authentication.
- **Task Management**: Create, update, delete, and retrieve tasks.
- **Real-time Notifications**: WebSocket support for real-time task updates.
- **Report Generation**: Generate and retrieve task reports.
- **Redis Caching**: Cache task data in Redis for faster retrieval.
- **Email Notifications**: Send email notifications for various actions.
- **Docker Support**: Containerized application for easy deployment.
- **CI/CD**: GitHub Actions for continuous integration and deployment.
- **Swagger Documentation**: API documentation using Swagger.
- **Unit and Integration Tests**: Comprehensive testing for all endpoints.

## Endpoints

### User Authentication

- `POST /register`: Register a new user.
- `POST /login`: Authenticate a user and return a

JWT token.

### Task Management

- `POST /tasks`: Create a new task.
- `GET /tasks/:id`: Retrieve a task by ID.
- `PUT /tasks/:id`: Update a task.
- `DELETE /tasks/:id`: Delete a task.

### Real-time Notifications

- `GET /ws`: WebSocket endpoint for real-time notifications.

### Report Generation

- `GET /reports`: Generate and retrieve task reports.

## Setup

1. Clone the repository.
2. Install dependencies.
3. Set up environment variables.
4. Run the application.

## Usage

To run the application, use:

```bash
go run cmd/main.go
```

To run tests, use:

```bash
go test ./...
```

## Docker

To build and run the Docker container, use:

```bash
docker build -t task-manager .
docker run -p 8080:8080 task-manager
```

## Documentation

API documentation is available at `/swagger/index.html` when the application is running.

## License

This project is licensed under the MIT License.
