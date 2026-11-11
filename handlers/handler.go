package handlers

import (
	"net/http"
	"sisprese/models"
	"time"

	"github.com/gin-gonic/gin"
)

var Employees = []models.Employee{
	{ID: "EMP001", Name: "Budi Santoso", Department: "Operasional Gas"},
	{ID: "EMP002", Name: "Siti Aminah", Department: "Humas & Administrasi"},
}

var Attendances = []models.Attendance{}

func GetEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": len(Employees),
		"data":  Employees,
	})
}

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

// 3. Clock-In dengan Validasi Ganda (Tidak boleh Clock-In dua kali tanpa Clock-Out)
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

	// Cek status presensi terakhir karyawan tersebut hari ini
	var lastStatus string = ""
	for _, att := range Attendances {
		if att.EmployeeID == req.EmployeeID {
			// Cek apakah presensi dicatat pada tanggal yang sama hari ini
			if att.Timestamp.Format("2006-01-02") == time.Now().Format("2006-01-02") {
				lastStatus = att.Type
			}
		}
	}

	// Jika status terakhir adalah "IN", artinya belum Clock-Out
	if lastStatus == "IN" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal: Karyawan sudah melakukan Clock-In dan belum Clock-Out!"})
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

// 4. Clock-Out dengan Validasi (Harus sudah Clock-In sebelumnya)
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

	// Cek status presensi terakhir hari ini
	var lastStatus string = ""
	for _, att := range Attendances {
		if att.EmployeeID == req.EmployeeID {
			if att.Timestamp.Format("2006-01-02") == time.Now().Format("2006-01-02") {
				lastStatus = att.Type
			}
		}
	}

	// Jika belum pernah Clock-In atau sudah Clock-Out sebelumnya
	if lastStatus == "" || lastStatus == "OUT" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal: Karyawan belum melakukan Clock-In hari ini!"})
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

func GetAttendance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"total": len(Attendances),
		"data":  Attendances,
	})
}