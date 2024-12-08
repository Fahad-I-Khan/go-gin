# User API - Go with Gin, Swagger, PostgreSQL, and Docker

This is a simple API for managing users using the **Gin web framework**, **PostgreSQL**, and **Swagger** for API documentation. The project uses **Docker** for containerization, so it can easily be deployed and run in isolated environments.

## Table of Contents

1. [Technologies Used](#technologies-used)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Setup Instructions](#setup-instructions)
- [Step 1: Clone the Repository](#step-1-clone-the-repository)
- [Step 2: Install Go Dependencies](#step-2-install-go-dependencies)
5. [How to Run the Application](#how-to-run-the-application)
6. [API Endpoints](#api-endpoints)
7. [Dependencies](#dependencies)
8. [Swagger Documentation](#swagger-documentation)

## Technologies Used

![Go](https://img.shields.io/badge/Language-Go-blue) ![Gin](https://img.shields.io/badge/Framework-Gin-brightgreen) ![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-blue) ![Swagger](https://img.shields.io/badge/API-Swagger-orange) ![Docker](https://img.shields.io/badge/Docker-Enabled-blue) ![Docker Compose](https://img.shields.io/badge/Docker%20Compose-Used-blueviolet)

## Description

This is a simple RESTful API built using the **Go** programming language with the **Gin** framework, connected to a **PostgreSQL** database. The API documentation is powered by **Swagger**, and the project is containerized using **Docker** and **Docker Compose**.


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

### Step 1: Clone the repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/Fahad-I-Khan/go-gin-gorm.git
cd go-gin-gorm
```

## Step 2: Install Go Dependencies
Run the following commands to install the necessary Go dependencies for the project:

1. **Install Gin** - the web framework used in this project:

```bash
go get github.com/gin-gonic/gin
```
2. **Install PostgreSQL driver** for Go:

```bash
go get github.com/lib/pq
```
3. **Install Swagger dependencies** for API documentation generation:

- **Install Swag CLI**: This is a command-line tool used to generate Swagger documentation from Go comments.

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

4. **Install Swagger UI for Gin**:

```bash
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```
These packages allow you to serve the Swagger UI within your Gin application.

5. **Install CORS middleware for Gin** (to handle cross-origin requests):

```bash
go get github.com/gin-contrib/cors
```
6. **Install Swagger UI for Gin**:

```bash
go get github.com/swaggo/gin-swagger
```
7. **Optional**: Run `go mod tidy` to clean up the `go.mod` file and download any missing dependencies:

```bash
go mod tidy
```
8. **Run Swag Init**: Generate the Swagger documentation in the project. This creates a `docs` directory that contains the API docs. **Important**: You should run `swag init` whenever you change your API code or Swagger comments to regenerate the Swagger documentation.

```bash
swag init
```
This will scan your Go code and generate the Swagger documentation in the `docs` folder. **Make sure to run** `swag init` each time you modify API routes or add/modify Swagger comments in your code to keep the documentation up-to-date.
### 6. Stopping the Application
To stop the application, use the following command:

```bash
docker-compose down
```
This will stop and remove the containers, but the data in the PostgreSQL container persists due to the `pgdata` volume.

## How to Run the Application
### Step 1: Build and Start Docker Containers
You can start the Go application and PostgreSQL database using Docker Compose.

1. Build the Docker images and start the containers in the background:

```bash
docker-compose -d go_db
docker-compose build
docker-compose up -d go-app
```
Here's what each command does:

- `docker-compose -d go_db`: Starts the PostgreSQL container in detached mode.
- `docker-compose build`: Builds the Go application Docker image.
- `docker-compose up -d go-app`: Starts the Go application container in detached mode.

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

## Swagger Documentation
To view the API documentation with Swagger, navigate to the following URL once the application is running:
```bash
http://localhost:8000/swagger/index.html
```
This will open the Swagger UI, where you can interact with all the available API endpoints and see the API documentation.