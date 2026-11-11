package handlers

import (
	"net/http"
	"sisprese/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Database sementara (In-Memory)
var Employees = []models.Employee{
	{ID: "EMP001", Name: "Budi Santoso", Department: "Operasional Gas"},
	{ID: "EMP002", Name: "Siti Aminah", Department: "Humas & Administrasi"},
}

var Attendances = []models.Attendance{}

// 1. Get All Employees
func GetEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": len(Employees),
		"data":  Employees,
	})
}

// 2. Add Employee
func AddEmployee(c *gin.Context) {
	var newEmp models.Employee
	if err := c.ShouldBindJSON(&newEmp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Employees = append(Employees, newEmp)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Karyawan berhasil ditambahkan",
		"data":    newEmp,
	})
}

// 3. Clock-In
func ClockIn(c *gin.Context) {
	var req struct {
		EmployeeID string `json:"employee_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee_id wajib diisi"})
		return
	}

	// Validasi karyawan terdaftar
	found := false
	for _, emp := range Employees {
		if emp.ID == req.EmployeeID {
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID Karyawan tidak ditemukan"})
		return
	}

	newAtt := models.Attendance{
		ID:         time.Now().Format("20060102150405"),
		EmployeeID: req.EmployeeID,
		Type:       "IN",
		Timestamp:  time.Now(),
	}
	Attendances = append(Attendances, newAtt)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Clock-In berhasil dicatat",
		"data":    newAtt,
	})
}

// 4. Clock-Out (Fitur Baru Level 1)
func ClockOut(c *gin.Context) {
	var req struct {
		EmployeeID string `json:"employee_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee_id wajib diisi"})
		return
	}

	// Validasi karyawan terdaftar
	found := false
	for _, emp := range Employees {
		if emp.ID == req.EmployeeID {
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID Karyawan tidak ditemukan"})
		return
	}

	newAtt := models.Attendance{
		ID:         time.Now().Format("20060102150405"),
		EmployeeID: req.EmployeeID,
		Type:       "OUT",
		Timestamp:  time.Now(),
	}
	Attendances = append(Attendances, newAtt)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Clock-Out berhasil dicatat",
		"data":    newAtt,
	})
}

// 5. Get All Attendance History
func GetAttendance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": len(Attendances),
		"data":  Attendances,
	})
}