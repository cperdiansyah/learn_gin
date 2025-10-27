# Gin REST API Tutorial

This is a RESTful API built using the **Gin** framework in Go. It demonstrates a complete CRUD (Create, Read, Update, Delete) implementation for a `User` entity. The project follows idiomatic Go practices, uses a clean architecture with **Repositories**, **Services**, and **Handlers (Controllers)**, and includes features like graceful shutdown, environment variable support, structured logging with **Zerolog**, and CORS configuration.

This project is for my Medium article. [Medium](https://medium.com/@godusan)

---

## Features

- **PostgreSQL Database**: Uses PostgreSQL as the data source for storing user data.
- **Graceful Shutdown**: Ensures the server shuts down gracefully without interrupting ongoing requests.
- **Environment Variables**: Supports `.env` files for configuration management.
- **Structured Logging**: Uses **Zerolog** for fast and structured logging.
- **CORS Support**: Configured to handle Cross-Origin Resource Sharing (CORS).
- **Clean Architecture**: Follows a clean and idiomatic folder structure with separation of concerns:
  - **Repositories**: Handle database operations.
  - **Services**: Contain business logic.
  - **Handlers (Controllers)**: Manage HTTP requests and responses.
- **Complete CRUD Example**: Full implementation of CRUD operations for a `User` entity.

---

## Prerequisites

Before running the project, ensure you have the following installed:

1. **Go** (version 1.20 or higher)
2. **PostgreSQL** (or a PostgreSQL-compatible database)
3. **Git** (for cloning the repository)

---

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/gin-rest-tutorial.git
cd gin-rest-tutorial
```

### 2. Set Up Environment Variables
Create a `.env` file in the root directory and add the following variables:

```env
# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

# Server configuration
SERVER_PORT=8080
```

### 3. Install Dependencies
Run the following command to install the required Go dependencies:

```bash
go mod tidy
```

### 4. Set Up the Database
Create a PostgreSQL database using the name specified in the `.env` file.

Run the following SQL script to create the `users` table:

```sql
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       first_name VARCHAR(100) NOT NULL,
                       last_name VARCHAR(100) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       phone_number VARCHAR(20)
);

INSERT INTO users (first_name, last_name, email, phone_number)
VALUES
    ('John', 'Doe', 'john.doe@example.com', '1234567890'),
    ('Jane', 'Smith', 'jane.smith@example.com', '0987654321');
```

### 5. Run the Application
Start the application using the following command:

```bash
go run main.go
