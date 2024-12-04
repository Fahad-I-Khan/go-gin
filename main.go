package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// @title User API
// @version 1.0
// @description This is a simple API for managing users in a PostgreSQL database.
// @host localhost:8000
// BasePath /api/v1 Because of this in url "/api/v1" was repeating and causing the error.
// @contact.name API Support
// @contact.url http://localhost:8000/support   // Local URL for your development environment
// @contact.email support@localhost.com
func main() {
	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Add swagger
	// Serve Swagger UI
	r := gin.Default()
	r.Use(cors.Default())
	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define other routes here...
	r.GET("/api/v1/users", getUsers)
	r.GET("/api/v1/users/:id", getUser)
	r.POST("/api/v1/users", createUser)
	r.PUT("/api/v1/users/:id", updateUser)
	r.DELETE("/api/v1/users/:id", deleteUser)

	// // Add your API routes here
	// r.GET("/api/v1/users", getUsers, getUser, createUser, updateUser, deleteUser)

	r.Run(":8000")

}

// get all users

// @Summary Get all users
// @Description Retrieve a list of all users in the database
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} User
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users [get]
func getUsers(c *gin.Context) {
	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query to get all users from the database
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error fetching users"})
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"message": "Getting all users",
	})

	// Return the list of users as JSON
	c.JSON(http.StatusOK, users)
}

// get user by id

// @Summary Get user by ID
// @Description Retrieve a single user's details by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID" // The ID of the user to retrieve
// @Success 200 {object} User // The user object returned in the response
// @Failure 400 {object} ErrorResponse // Bad request if the ID is invalid
// @Failure 404 {object} ErrorResponse // User not found
// @Failure 500 {object} ErrorResponse // Internal server error
// @Router /api/v1/users/{id} [get]
func getUser(c *gin.Context) {
	id := c.Param("id")

	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Internal server error"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Getting a user by ID",
	})

	c.JSON(http.StatusOK, u)
}

// create user

// @Summary Create a new user
// @Description Create a new user by providing a name and email
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body User true "New user information"
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users [post]
func createUser(c *gin.Context) {
	var u User

	// Parse the incoming request body into the User struct
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input"})
		return
	}

	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert the new user into the database
	err = db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create user"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Creating a new user",
	})

	c.JSON(http.StatusCreated, u) // Return the newly created user
}

// update user

// @Summary Update an existing user
// @Description Update a user's name and email by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID" // This is the ID parameter from the URL path
// @Param user body User true "Updated user information" // The request body (updated user data)
// @Success 200 {object} User // The updated user object returned in the response
// @Failure 400 {object} ErrorResponse // Bad request if the input is invalid
// @Failure 404 {object} ErrorResponse // If the user is not found
// @Failure 500 {object} ErrorResponse // Internal server error
// @Router /api/v1/users/{id} [put]
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var u User

	// Parse the incoming request body into the User struct
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input"})
		return
	}

	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update the user in the database
	_, err = db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to update user"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Updating user",
	})

	c.JSON(http.StatusOK, u) // Return the updated user
}

// delete user

// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID" // ID of the user to delete
// @Success 200 {string} string "User deleted" // Success message
// @Failure 404 {object} ErrorResponse // If the user is not found
// @Failure 500 {object} ErrorResponse // Internal server error
// @Router /api/v1/users/{id} [delete]
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// Connect to the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the user exists
	var u User
	err = db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Error fetching user"})
		return
	}

	// Delete the user
	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to delete user"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Deleting user",
	})

	c.JSON(http.StatusOK, "User deleted")
}
