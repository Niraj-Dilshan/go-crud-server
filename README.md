# CRUD Server

A RESTful API server built with Go, Gin, and GORM that provides CRUD operations for managing tasks.

## Project Structure

```
.
├── .env.example        # Example environment variables
├── .env                # Environment variables for development
├── .test.env           # Environment variables for testing
├── .air.toml           # Configuration for Air (live reloading)
├── go.mod              # Go module dependencies
├── go.sum              # Go module checksum
├── Dockerfile          # Docker configuration for building the application
├── docker-compose.yml  # Docker Compose configuration for local development
├── api/                # API handlers and routes
│   └── task/           # Task-related API endpoints
├── cmd/                # Application entrypoints
│   └── main.go         # Main application entry point
├── config/             # Application configuration
├── keys/               # Authorization keys and certificates
├── startup/            # Server initialization and module setup
├── utils/              # Utility functions and helpers
└── tmp/                # Temporary files (generated during development)
```

## Technologies Used

- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM library for database operations
- **PostgreSQL** - Database
- **Docker** - Containerization
- **Air** - Live reloading for development

## Getting Started

### Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose (for local development with database)
- Git

### Environment Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd crud-server
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   ```
   Edit the `.env` file as needed with your configuration.

### Running Locally

#### Option 1: Using Go directly

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run cmd/main.go
   ```

#### Option 2: Using Air (with live reloading)

1. Install Air if you haven't already:
   ```bash
   go install github.com/cosmtrek/air@latest
   ```

2. Run the application with Air:
   ```bash
   air
   ```

#### Option 3: Using Docker Compose

1. Start the application and database:
   ```bash
   docker-compose up
   ```

2. For background execution:
   ```bash
   docker-compose up -d
   ```

## Development

### Adding New Features

1. Create new API handlers in the `api/` directory, following the module pattern
2. Update routes in the appropriate module initialization in `startup/`
3. Add any required models and database migrations
4. Test your changes locally

### Code Structure Guidelines

- Place business logic in appropriate modules
- Follow Go best practices for error handling and documentation
- Write tests for new functionality

### Testing

Run tests using:

```bash
go test ./...
```

## Deployment

### Building the Application

Build the executable binary:

```bash
go build -o app cmd/main.go
```

### Docker Deployment

1. Build the Docker image:
   ```bash
   docker build -t crud-server .
   ```

2. Run the container:
   ```bash
   docker run -p 8888:8888 --env-file .env crud-server
   ```

### Production Deployment

For production deployment, you should:

1. Set up a proper database instance (not the Docker development instance)
2. Configure all environment variables appropriately
3. Use Docker or a container orchestration tool like Kubernetes
4. Set up proper logging and monitoring

## API Endpoints

The server provides the following API endpoints (on the `/api` route):

- Task management (exact endpoints depend on implementation)
  - GET, POST, PUT, DELETE operations