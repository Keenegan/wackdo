# Wackdo

A REST API for managing a food ordering system built with Go and Gin.

## Documentation

- [Class Diagram](docs/CLASS-DIAGRAM.md) - System architecture and data models
- [API documentation](https://keenegan.github.io/wackdo/)

## Setup

1. Copy the environment file and fill in the values:
   ```bash
   cp .env.dist .env
   ```

2. Edit `.env` with your database credentials and JWT secret

3. Start the database:
   ```bash
   make up
   ```

4. Run the application:
   ```bash
   make run
   ```

The API will be available at `http://localhost:8080`

Database admin interface (Adminer) is available at `http://localhost:8081`

## Development Commands

- `make run` - Start the application
- `make test` - Run tests
- `make check` - Run staticcheck linter
- `make up` - Start Docker containers
- `make down` - Stop Docker containers
- `make restart` - Restart Docker containers

## Roles

The application has four user roles:

- **Manager** - Full access to manage menus, products, users, and orders
- **Cashier** - Can create orders
- **Prep** - Can update order status
- **Admin** - Superuser with all permissions