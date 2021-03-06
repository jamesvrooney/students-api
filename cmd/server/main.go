package main

import (
	"fmt"
	"net/http"
	"students-api/internal/database"
	internalHttp "students-api/internal/http"
	"students-api/internal/services/student"
)

func Run() error {
	fmt.Println("Running App")

	db, err := database.InitDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	studentService := student.NewService(db)

	handler := internalHttp.NewHandler(studentService)
	handler.InitRoutes()
	if err := http.ListenAndServe(":9000", handler.Router); err != nil {
		return err
	}
	return nil
}

func main() {
	err := Run()
	if err != nil {
		fmt.Println("Error running app")
		fmt.Println(err)
	}
}
