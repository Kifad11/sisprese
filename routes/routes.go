package routes

import (
	"net/http"
	"sisprese/handlers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, handler *handlers.Handler) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server Sisprese Clean Architecture aktif!"})
	})
	router.GET("/employees", handler.GetEmployees)
	router.POST("/employees", handler.AddEmployee)
	router.POST("/attendance/clock-in", handler.ClockIn)
	router.POST("/attendance/clock-out", handler.ClockOut)
	router.GET("/attendance", handler.GetAttendance)
}
