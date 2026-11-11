package main

import (
	"log"
	"sisprese/config"
	"sisprese/handlers"
	"sisprese/repositories"
	"sisprese/routes"
	"sisprese/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.OpenDatabase()
	if err != nil {
		log.Fatalf("Gagal menginisialisasi database: %v", err)
	}
	repository := repositories.New(db)
	attendanceService := services.NewAttendanceService(repository)
	handler := handlers.New(repository, attendanceService)

	r := gin.Default()
	routes.Register(r, handler)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
