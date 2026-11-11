package repositories

import (
	"time"

	"sisprese/models"

	"gorm.io/gorm"
)

type Repository struct{ db *gorm.DB }

func New(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) ListEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := r.db.Order("id").Find(&employees).Error
	return employees, err
}

func (r *Repository) GetEmployee(id string) (models.Employee, error) {
	var employee models.Employee
	err := r.db.First(&employee, "id = ?", id).Error
	return employee, err
}

func (r *Repository) CreateEmployee(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *Repository) CreateAttendance(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *Repository) LatestAttendanceToday(employeeID string, now time.Time) (models.Attendance, error) {
	var attendance models.Attendance
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	end := start.AddDate(0, 0, 1)
	err := r.db.Where("employee_id = ? AND timestamp >= ? AND timestamp < ?", employeeID, start, end).
		Order("timestamp DESC").First(&attendance).Error
	return attendance, err
}

func (r *Repository) ListAttendances(date, employeeID string) ([]models.Attendance, error) {
	query := r.db.Model(&models.Attendance{})
	if date != "" {
		start, _ := time.ParseInLocation("2006-01-02", date, time.Local)
		query = query.Where("timestamp >= ? AND timestamp < ?", start, start.AddDate(0, 0, 1))
	}
	if employeeID != "" {
		query = query.Where("employee_id = ?", employeeID)
	}
	var attendances []models.Attendance
	err := query.Order("timestamp DESC").Find(&attendances).Error
	return attendances, err
}
