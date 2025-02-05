package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	//db connection
	db, err := sql.Open("mysql", "drone:DroneBreaker55@tcp(localhost:3306)/droners")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initializing echo instance
	e := echo.New()

	e.Logger.Fatal(e.Start(":4000"))
	fmt.Println("Hi, welcome to my backend, administrator")
}
