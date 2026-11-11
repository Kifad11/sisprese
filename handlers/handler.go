package handlers

import (
	"errors"
	"net/http"
	"time"

	"sisprese/models"
	"sisprese/repositories"
	"sisprese/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	repository *repositories.Repository
	attendance *services.AttendanceService
}

func New(repository *repositories.Repository, attendance *services.AttendanceService) *Handler {
	return &Handler{repository: repository, attendance: attendance}
}

func (h *Handler) GetEmployees(c *gin.Context) {
	employees, err := h.repository.ListEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data karyawan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": len(employees), "data": employees})
}

func (h *Handler) AddEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repository.CreateEmployee(&employee); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"error": "ID Karyawan sudah terdaftar"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data karyawan"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Karyawan berhasil ditambahkan", "data": employee})
}

type attendanceRequest struct {
	EmployeeID string `json:"employee_id" binding:"required"`
	Note       string `json:"note"`
}

func (h *Handler) ClockIn(c *gin.Context) {
	var req attendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee_id wajib diisi"})
		return
	}
	attendance, err := h.attendance.ClockIn(req.EmployeeID, req.Note)
	if err != nil {
		h.writeAttendanceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Clock-In berhasil dicatat", "data": attendance})
}

func (h *Handler) ClockOut(c *gin.Context) {
	var req attendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee_id wajib diisi"})
		return
	}
	attendance, err := h.attendance.ClockOut(req.EmployeeID)
	if err != nil {
		h.writeAttendanceError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Clock-Out berhasil dicatat", "data": attendance})
}

func (h *Handler) GetAttendance(c *gin.Context) {
	date := c.Query("date")
	if date != "" {
		if _, err := time.Parse("2006-01-02", date); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format date harus YYYY-MM-DD"})
			return
		}
	}
	attendances, err := h.attendance.List(date, c.Query("employee_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data presensi"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": len(attendances), "data": attendances})
}

func (h *Handler) writeAttendanceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrEmployeeNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrAlreadyClockedIn), errors.Is(err, services.ErrNotClockedIn):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan presensi"})
	}
}
