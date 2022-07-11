# gRPC Logger Server 

### You can use this gRPC server with [Evans](https://github.com/ktr0731/evans) or implement a real client.

## Tools & Libraries used
- [Golang 1.16+](https://golang.org/doc/go1.16)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/get-started)
- [gRPC](https://grpc.io/)

## 🗄 Project structure

### /app

**Folder with business logic only**. This directory doesn't care about what database driver you're using.

- `/app/controller` folder for functional controller (used in routes)
- `/app/dto` Data Transfer Objects(DTO) folder for transform data before sent to API clients
- `/app/model` folder for describe business models and methods of your project
- `/app/repository` folder for perform database operations for models of your project

### /pkg

**Folder with project-specific functionality.** This directory contains all the project-specific code tailored only for your business use case.
<!-- 
- `/pkg/config` folder for configuration functions
- `/pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `/pkg/route` folder for describe routes of your project
- `/pkg/validator` folder with validation functions -->

### /platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project,
like setting up the database, logger instance and storing migrations, seeds(demo data).

<!-- - `/platform/database` folder with database setup functions (by default, PostgreSQL)
- `/platform/logger` folder with better logger setup functions (by default, Logrus)
- `/platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)
- `/platform/seeds` folder with demo data for application rapid setup. mostly **sql** scripts -->

### /logger_pb

**Folder with proto file and auto-generated go files**. This directory contains all the go auto-generated files and proto file.

## ⚙️ Configuration


# Server Settings:
SERVER_READ_TIMEOUT=60

# Jwt settings:
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=30
JWT_SECRET_KEY="secret"

# Database settings:
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=password
DOCKER_SERVICE=mongo
MONGO_PORT=27017


``` ini
# .env file
# JWT settings:
JWT_SECRET_KEY="super_secret_here"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=120

# Database settings:
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=password
DOCKER_SERVICE=mongo
MONGO_PORT=27017
```

## 🔨 Docker development

- Install **`docker`**, **`docker-compose`**
- Rename `.env.example` to `.env`
- Start a MongoDB container exposing port 27017
<!-- - Run migrations `make migrate.up`
- Now start api server with hot reloading `make docker.dev`
- Visit **`http://localhost:5000`** or **`http://localhost:5000/swagger/`** -->

### Todo

- [ ] Add new functionalities
- [ ] Implement this gRPC server with a client. 

## ⚠️ License

[MIT](https://opensource.org/licenses/MIT)