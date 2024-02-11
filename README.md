
# Go DDD CQRS layout with Echo Framework

This is a Go application template utilizing Domain-Driven Design (DDD) and Command Query Responsibility Segregation (CQRS) principles with the Echo framework.

## Project Structure

```
.
├── cmd
│   └── myapp
│       ├── cli
│       │   ├── cli.go
│       │   ├── http_server.go
│       │   └── sql_migrate.go
│       └── main.go        # Entry point of the application.         
├── internal
│   ├── application
│   │   ├── app.go         # Application configuration and initialization.
│   │   ├── command
│   │   │   └── command.go # Command handling (CQRS).
│   │   └── query
│   │       └── query.go   # Query handling (CQRS).
│   ├── domain
│   │   ├── aggregate      # Domain aggregates.
│   │   ├── entity         # Domain entities.
│   │   ├── repository     # Domain repository interfaces.
│   │   └── service        # Domain services.
│   ├── genmocks.go        # Mock generation for testing.
│   ├── infrastructure
│   │   ├── api
│   │   │   └── http
│   │   │       ├── handler     # HTTP handlers.
│   │   │       ├── middleware  # HTTP middleware.
│   │   │       ├── routes.go   # HTTP route configuration.
│   │   │       ├── server.go   # HTTP server configuration.
│   │   │       └── validator   # HTTP request validation.
│   │   ├── decorator
│   │   │   ├── decorator.go # Decorators (e.g., for logging).
│   │   │   └── logging.go   # Logging.
│   │   └── pgsql
│   │       ├── model        # Data models for PostgreSQL.
│   │       └── repository   # Domain repository implementation for PostgreSQL.
│   └── mocks                # Mocks for testing.
└── pkg
    └── logger
        └── logger.go        # Logging service.
```

## Component Descriptions

### `cmd`
Contains the application's entry point (`main.go`), where the server is started and all necessary components are initialized.

### `internal`
The core of the application, divided into layers according to DDD and CQRS.

#### `application`
The application layer, responsible for coordinating actions and delegating tasks between the domain layer and the infrastructure layer. Includes command and query processing.

#### `domain`
The heart of the domain model, containing business logic, entities, aggregates, repository interfaces, and domain services.

#### `infrastructure`
The layer that facilitates communication between the application layer and the external world (databases, web services, etc.). Includes API implementation, database access, etc.

### `pkg`
General packages for the entire application, e.g., the logging service.

## Getting Started

(Here you can add instructions on how to install, configure, and run your project.)

