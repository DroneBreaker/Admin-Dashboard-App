package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/handlers"
	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/repository"
	"github.com/DroneBreaker/Admin-Dashboard-App/backend/internal/services"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	// Db connection
	db, err := sql.Open("mysql", "drone:DroneBreaker55@tcp(localhost:3306)/droners")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Initialize echo
	e := echo.New()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	//user routes
	e.GET("/users", userHandler.GetAll)
	e.POST("/users", userHandler.Create)
	e.GET("/users/:id", userHandler.GetByID)
	e.PUT("/users/:id", userHandler.Update)
	// e.DELETE("/users/:id", userHandler.Delete)

	e.Logger.Fatal(e.Start(":4000"))
	fmt.Println("Welcome to the admin backend, administrators")
}
