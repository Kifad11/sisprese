package main

import (
	"sisprese/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success", "message": "Server Sisprese Clean Architecture aktif!"})
	})

	// Routes Karyawan
	r.GET("/employees", handlers.GetEmployees)
	r.POST("/employees", handlers.AddEmployee)

	// Routes Presensi (Clock-In & Clock-Out)
	r.POST("/attendance/clock-in", handlers.ClockIn)
	r.POST("/attendance/clock-out", handlers.ClockOut)
	r.GET("/attendance", handlers.GetAttendance)

	// Jalankan Server
	r.Run(":8080")
}