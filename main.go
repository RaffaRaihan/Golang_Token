package main

import (
	"fmt"
	"main/controllers"
	"main/middleware"
	"net"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	// Membaca konfigurasi server
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err.Error())
	}

	// Membuat instance baru controllers
	var db *gorm.DB

	userController := controllers.NewCRUDController(db)

	http.HandleFunc("/api/v1/users", middleware.TokenValidation(userController.ListUsers))
	http.HandleFunc("/api/v1/users/get", middleware.TokenValidation(userController.GetUser))
	http.HandleFunc("/api/v1/users/create", middleware.TokenValidation(userController.CreateUser))
	http.HandleFunc("/api/v1/users/update", middleware.TokenValidation(userController.UpdateUser))
	http.HandleFunc("/api/v1/users/delete", middleware.TokenValidation(userController.DeleteUser))

	// Set address dan port tempat server berjalan
	ln, err := net.Listen("tcp", "localhost:3333")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

	fmt.Println("Server is running on http://localhost:3333")
	if err := http.Serve(ln, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}