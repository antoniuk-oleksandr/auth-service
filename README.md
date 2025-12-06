# Auth Service

A production-grade authentication service built with Go, featuring a fully modular monorepo architecture and clean design principles.

## 🎯 Overview

This project demonstrates a comprehensive implementation of modern backend development practices, including Clean Architecture, custom TCP protocol communication, and full-stack web development with server-side rendering.

## ✨ Key Features

### Architecture & Design
- **Clean Architecture** with strict layer separation (domain → application → infrastructure → presentation)
- **SOLID Principles** applied throughout the codebase
- **DDD-inspired** domain modeling with clear bounded contexts
- **Modular DI System** using Uber FX with named routers, modules, and dependency graphs
- **Repository & Factory Patterns** for data access abstraction
- **Transaction Manager** for database operation consistency

### Communication Layer
- **Custom Binary TCP Protocol** (`custom-protoc`) with:
  - Request/response framing
  - Custom encoders and decoders
  - Reusable client/server layer
  - Type-safe message handling

### Frontend
- **HTMX + Templ** for modern server-side rendering
- Component-based layouts with reusable UI elements
- Server-side templates with type safety
- Interactive forms with progressive enhancement

### Backend Services
- JWT-based authentication with refresh tokens
- User registration and management
- Password hashing with bcrypt
- MongoDB integration with transaction support
- PostgreSQL-ready architecture

### Observability
- **Extensible Logger System** with multiple implementations:
  - Zap Console logger
  - Zap JSON logger
  - Custom CTP logger
  - Factory pattern for logger selection
- **Monitoring Stack**:
  - Promtail for log collection
  - Grafana Loki for log aggregation
  - Request logging middleware

### Infrastructure
- Docker containerization
- Docker Compose orchestration
- Environment-based configuration
- Development and production configurations

## 🏗️ Project Structure

```
.
├── backend/                # Backend service (authentication API)
│   ├── cmd/                # Application entry points
│   ├── internal/           # Internal packages
│   │   ├── application/    # Business logic & use cases
│   │   ├── domain/         # Domain models & interfaces
│   │   ├── infra/          # Infrastructure implementations
│   │   ├── presentation/   # HTTP & CTP handlers
│   │   ├── config/         # Configuration management
│   │   ├── db/             # Database factories & transactions
│   │   └── logger/         # Logging implementations
│   └── pkg/                # Reusable packages
│
├── frontend/               # Frontend service (web UI)
│   ├── cmd/                # Application entry points
│   ├── internal/           # Internal packages
│   │   ├── application/    # Business logic
│   │   ├── domain/         # Domain models
│   │   ├── presentation/   # Fiber handlers
│   │   └── view/           # Templ templates
│   │       ├── component/  # Reusable UI components
│   │       ├── layout/     # Page layouts
│   │       └── page/       # Page templates
│   └── static/             # Static assets (JS, CSS)
│
├── ctp/                    # Custom TCP Protocol library
│   ├── client/             # TCP client implementation
│   ├── server/             # TCP server implementation
│   ├── types/              # Protocol types & interfaces
│   └── internal/           # Protocol internals
│
├── common/                 # Shared utilities
│   ├── config/             # Common configuration
│   ├── env_parser/         # Environment variable parsing
│   └── validator/          # Input validation
│
├── infra/                  # Infrastructure configuration
│   └── dev/                # Development environment
│       ├── docker-compose.yaml
│       ├── loki-config.yaml
│       └── promtail-config.yaml
│
├── justfile                # Command runner for development tasks
└── go.work                 # Go workspace configuration
```

## 🛠️ Tech Stack

- **Language**: Go
- **Web Framework**: Fiber
- **Frontend**: HTMX, Templ
- **Database**: MongoDB (PostgreSQL-ready)
- **DI Container**: Uber FX
- **Logging**: Zap, Promtail, Loki
- **Authentication**: JWT
- **Containerization**: Docker, Docker Compose
- **Task Runner**: just
- **Architecture**: Clean Architecture, SOLID, Design Patterns
- **Networking**: Custom TCP Protocol

## 🚀 Getting Started

### Prerequisites
- Docker & Docker Compose
- [just](https://github.com/casey/just) command runner (optional, but recommended)

### Quick Start

1. Clone the repository:
```bash
git clone <repository-url>
cd auth-service
```

2. Set up environment variables:
```bash
# Copy and configure environment files
cp infra/dev/.env.example infra/dev/.env
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env

# Edit the .env files with your actual values
```

3. Start all services with Docker Compose:
```bash
# Using just (recommended)
just compose-up-build

# Or directly with docker-compose
docker-compose -f infra/dev/docker-compose.yaml up -d --build
```

4. Access the application:
- Frontend: http://localhost:3000
- Backend: http://localhost:8080

### Environment Configuration

#### Infrastructure (`infra/dev/.env`)
```env
# Mongo
MONGO_USER=admin
MONGO_PASSWORD=strong_password
MONGO_PORT=27017
MONGO_CONTAINER_NAME=ctp-mongo
MONGO_DBNAME=ctp

# Frontend
FRONTEND_PORT=3000
BACKEND_ADDR=ctp-backend:8080

# Backend
BACKEND_PORT=8080
BACKEND_CONTAINER_NAME=ctp-backend
JWT_SECRET_KEY=your_super_secret_jwt_key_here
JWT_ACCESS_TOKEN_TTL=3600
JWT_REFRESH_TOKEN_TTL=86400
BACKEND_HTTP_FRAMEWORK=fiber
BACKEND_HASHER_COST=10
BACKEND_LOGGER_TYPE=zap_console
BACKEND_DATABASE_TYPE=mongo
```

#### Backend (`backend/.env`)
```env
# JWT
JWT_SECRET_KEY=your_super_secret_jwt_key_here
JWT_ACCESS_TOKEN_TTL=3600
JWT_REFRESH_TOKEN_TTL=86400

# Server
SERVER_PORT=8080
HTTP_FRAMEWORK=fiber

# Hasher
HASHER_COST=10

# Logger
LOGGER_TYPE=zap_console

# Database
DATABASE_TYPE=mongo

# MongoDB
MONGO_DBNAME=ctp
MONGO_URI=mongodb://admin:strong_password@localhost:27017/?authSource=admin&replicaSet=rs0&directConnection=true
```

#### Frontend (`frontend/.env`)
```env
BACKEND_ADDRESS=localhost:8080
SERVER_PORT=3000
```

## 📦 Modules

### Backend Module
- User authentication and authorization
- JWT token generation and validation
- Password hashing and verification
- MongoDB repository implementation
- Custom TCP protocol server

### Frontend Module
- Login and registration pages
- Server-side rendered templates
- HTMX-powered interactivity
- Custom TCP protocol client

### CTP Module (Custom TCP Protocol)
- Binary protocol for service-to-service communication
- Client and server implementations
- Request/response framing
- Type-safe message encoding/decoding

### Common Module
- Shared configuration utilities
- Environment variable parsing
- Input validation with validator v10

## 🔧 Development Commands

The project uses [just](https://github.com/casey/just) as a command runner. Available commands:

```bash
# Start all services
just compose-up
just cu  # alias

# Stop all services
just compose-down
just cd  # alias

# Rebuild and start all services
just compose-up-build
just cub  # alias

# Start only MongoDB services
just compose-up-mongo
just mongo  # alias
```

### Without just

If you don't have `just` installed, you can use docker-compose directly:

```bash
# Start all services
docker-compose -f infra/dev/docker-compose.yaml up -d

# Stop all services
docker-compose -f infra/dev/docker-compose.yaml down

# Rebuild and start
docker-compose -f infra/dev/docker-compose.yaml up -d --build

# Start only MongoDB
docker-compose -f infra/dev/docker-compose.yaml up -d ctp-mongo ctp-mongo-setup
```

## 📝 Design Patterns Used

- **Repository Pattern**: Data access abstraction
- **Factory Pattern**: Object creation (DB, Logger)
- **Dependency Injection**: Uber FX for IoC
- **Service Layer**: Business logic encapsulation
- **Mapper Pattern**: DTO/Entity conversion
- **Transaction Manager**: Database transaction handling
- **Strategy Pattern**: Multiple logger implementations

## 🔐 Security Features

- Password hashing with bcrypt (configurable cost)
- JWT-based authentication with configurable TTL
- Refresh token rotation
- Environment-based secrets management
- Input validation
- Error handling without information leakage

## 📊 Monitoring

The service includes a complete observability stack:
- Request logging middleware
- Structured JSON logging
- Log aggregation with Loki
- Log shipping with Promtail
- Visualization with Grafana (configurable)

## 🐳 Docker Architecture

The application is fully containerized with Docker Compose handling:
- Automatic service builds
- Network configuration
- Volume management
- MongoDB replica set setup
- Service dependencies
- Environment variable injection

## 🤝 Contributing

This is a personal portfolio project demonstrating modern Go development practices. Feel free to explore the code and use it as a reference for your own projects.

---

**Note**: This is a portfolio project showcasing modern backend development practices with Go. It demonstrates real-world application of Clean Architecture, SOLID principles, and production-ready patterns.