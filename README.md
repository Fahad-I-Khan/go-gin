# User API - Go with Gin, Swagger, PostgreSQL, and Docker

This is a simple API for managing users using the **Gin web framework**, **PostgreSQL**, and **Swagger** for API documentation. The project uses **Docker** for containerization, so it can easily be deployed and run in isolated environments.

## Table of Contents

1. [Project Overview](#project-overview)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Setup Instructions](#setup-instructions)
5. [API Endpoints](#api-endpoints)
6. [Dependencies](#dependencies)
7. [Swagger UI](#swagger-ui)

## Project Overview

- **Gin**: A web framework for Go (Golang) for building APIs.
- **Swagger**: For API documentation.
- **PostgreSQL**: A relational database used to store user data.
- **Docker**: For containerizing the application and the PostgreSQL database.
- **Docker Compose**: For managing multi-container applications.

## Features

- Create a new user
- Retrieve all users
- Retrieve a single user by ID
- Update an existing user
- Delete a user
- Auto-generating API documentation using Swagger

## Prerequisites

Before running the project, ensure you have the following installed:

1. **Docker** and **Docker Compose**:
   - Download and install [Docker](https://www.docker.com/get-started).
   - Docker Compose should be included in Docker installation (check [Docker Compose installation](https://docs.docker.com/compose/install/)).

2. **Go** (optional, for local development):
   - If you plan to work on the Go code locally (without Docker), you need Go installed on your machine. Follow the installation guide on the [Go website](https://golang.org/doc/install).

3. **PostgreSQL** (handled through Docker):
   - The application uses PostgreSQL, and it is configured to run in a Docker container, so no need to install it manually.

   ### main.go
This is the entry point for the application, which defines API routes and interacts with the PostgreSQL database. It also sets up Swagger for API documentation.

### Dockerfile
The Dockerfile is used to build the Docker image for the Go application. It includes the steps to copy the source code, install dependencies, and build the application.

### docker-compose.yml
The `docker-compose.yml` file defines two services:
1. **go-app**: The Go application that interacts with the PostgreSQL database.
2. **go_db**: A PostgreSQL database container.
   
It ensures that both services are linked together and can communicate.

## Setup Instructions

### 1. Clone the repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/Fahad-I-Khan/go-gin-gorm.git
cd go-gin-gorm
```

### 2. Start the PostgreSQL container
Start the `go_db` container first to ensure the PostgreSQL database is up and running before starting the Go application:

```bash
docker-compose up -d go_db
```

This starts the `go_db` container in detached mode (`-d`).
It ensures that the PostgreSQL database is running before you attempt to connect to it.

### 3. Build Docker Containers
Once the `go_db` container is up and running, you can now build the Go application container:

```bash
docker-compose build
```
- This builds the `go-app` Docker image based on the instructions in the `Dockerfile`.

### 4. Start the Go application container
Now that the PostgreSQL container is running and the application image is built, you can start the Go application container:

```bash
docker-compose up go-app
```
- This starts the `go-app` container in detached mode (`-d`).
The Go application will now be able to connect to the PostgreSQL database because the `go_db` container is already running.

### 5. Access the Application
Swagger UI for API documentation can be accessed at:

```bash
http://localhost:8000/swagger
```

### 6. Stopping the Application
To stop the application, use the following command:

```bash
docker-compose down
```
This will stop and remove the containers, but the data in the PostgreSQL container persists due to the `pgdata` volume.

## API Endpoints
The following endpoints are available in the application:

1. **GET** `/api/v1/users`
Retrieve all users.

**Response:**
`200 OK`: Returns a list of all users.

2.  **GET** `/api/v1/users/:id`
Retrieve a user by ID.

**Parameters:**
`id` (path): The ID of the user.

**Response:**
- `200 OK:` Returns the user object.
- `404 Not Found`: If the user with the given ID does not exist.

3. **POST** `/api/v1/users`
Create a new user.

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

**Response:**
- `201 Created`: Returns the created user object.

4. **PUT** `/api/v1/users/:id`
Update an existing user by ID.

**Parameters:**
`id` (path): The ID of the user to update.

**Request Body:**

```json
{
  "name": "Updated Name",
  "email": "updated.email@example.com"
}
```

**Response:**
- `200 OK:` Returns the updated user object.
- `404 Not` Found: If the user with the given ID does not exist.

5. **DELETE** `/api/v1/users/:id`
Delete a user by ID.

**Parameters:**
- `id` (path): The ID of the user to delete.
**Response:**
- `200 OK`: Returns a success message indicating the user was deleted.
- `404 Not Found`: If the user with the given ID does not exist.

## Dependencies
The Go application relies on the following dependencies:

- **Gin**: Web framework for Go.
- **Gin Swagger**: Swagger UI integration for Gin.
- **PostgreSQL**: Database for storing user information.
- **pq**: PostgreSQL driver for Go.

These dependencies are listed in the `go.mod` file. They are automatically installed when you build the Docker container or run `go mod tidy`.

To install dependencies locally (if running outside Docker):

```bash
go mod tidy
```

## Swagger UI
After starting the application, you can access the Swagger UI at:

```bash
http://localhost:8000/swagger
```
Swagger will display all the available API endpoints, their descriptions, and request/response details.