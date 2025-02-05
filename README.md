# Go Application

## Project Structure

```
.
├── cmd
│   └── app
│       └── main.go                     # Entry point of the application
├── internal
│   └── app
│       ├── config                      # Configuration setup
│       │   └── config.go
│       ├── controller
│       │   ├── common_controllers.go   # Common controllers
│       │   ├── users_controllers.go    # User-related controllers
│       ├── model
│       │   └── users_models.go   # User data models
│       ├── routes
│       │   ├── common_routes.go  # Common API routes
│       │   ├── users_routes.go   # User API routes
│       ├── utils
│       │   ├── database
│       │   │   ├── mongodb
│       │   │   │   └── conn.go   # MongoDB connection setup
│       │   │   ├── postgresql    # PostgreSQL setup (placeholder)
│       │   │   ├── redis         # Redis setup (placeholder)
│       │   ├── middleware
│       │   │   ├── basic_auth.go # Basic authentication middleware
│       │   │   ├── common.go     # Common middleware utilities
│       │   │   ├── wrapper.go    # Request wrapper utilities
├── tests                        # Test cases
├── tmp
│   ├── .air.toml                # Air configuration for live reloading
│   ├── .env                     # Environment variables file
│   ├── .gitignore                # Git ignore rules
│   ├── go.mod                    # Go module file
│   ├── go.sum                    # Dependencies checksum
│   ├── private.pem                # Private key (for authentication, encryption, etc.)
│   ├── public.pem                 # Public key
├── README.md                     # Project documentation
```

## Setup Instructions

1. **Clone the repository**

   ```sh
   git clone <repository-url>
   cd <project-directory>
   ```

2. **Install dependencies**

   ```sh
   go mod tidy
   ```

3. **Setup environment variables**

   - Create a `.env` file based on `.env.example`
   - Configure database and API credentials

4. **Run the application**

   ```sh
   go run cmd/app/main.go
   ```

5. **Run tests**

   ```sh
   go test ./...
   ```

## Features

- Modular architecture using `internal/` package
- RESTful API with structured controllers, models, and routes
- Middleware support for authentication and request handling
- Database integrations for MongoDB, PostgreSQL, and Redis (WIP)
- Environment configuration using `.env` file
- Live reloading support via [Air](https://github.com/cosmtrek/air)

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
