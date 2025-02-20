# Echo Store API

A RESTful API built with Go and Echo framework, following clean architecture principles.

## Features

-   User authentication with JWT
-   Clean architecture (Domain-driven design)
-   PostgreSQL database with GORM
-   Docker support
-   Environment-based configuration
-   Middleware for logging, CORS, and authentication

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── config.go              # Configuration management
├── internal/
│   ├── domain/                # Business logic and interfaces
│   ├── handler/               # HTTP handlers
│   ├── repository/            # Data access layer
│   └── usecase/               # Business logic implementation
├── pkg/
│   ├── middleware/            # Custom middleware
│   └── utils/                 # Utility functions
├── deployments/
│   └── Dockerfile            # Docker configuration
├── go.mod                    # Go module file
└── README.md                 # Project documentation
```

## Prerequisites

-   Go 1.21 or higher
-   PostgreSQL
-   Docker (optional)

## Getting Started

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/echo-store-api.git
    cd echo-store-api
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Set up environment variables:

    ```bash
    export PORT=8080
    export JWT_SECRET=your-secret-key
    export DB_HOST=localhost
    export DB_PORT=5432
    export DB_USER=postgres
    export DB_PASS=postgres
    export DB_NAME=echo_store
    ```

4. Run the application:
    ```bash
    go run cmd/main.go
    ```

## Docker Support

Build and run with Docker:

```bash
docker build -t echo-store-api -f deployments/Dockerfile .
docker run -p 8080:8080 echo-store-api
```

## API Endpoints

### Authentication

-   `POST /api/register` - Register a new user
-   `POST /api/login` - Login and get JWT token

### User Management

-   `GET /api/profile` - Get user profile (requires authentication)
-   `PUT /api/profile` - Update user profile (requires authentication)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
