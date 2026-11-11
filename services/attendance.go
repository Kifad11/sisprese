package services

import (
	"errors"
	"fmt"
	"time"

	"sisprese/models"
	"sisprese/repositories"

	"gorm.io/gorm"
)

var (
	ErrEmployeeNotFound = errors.New("ID Karyawan tidak ditemukan")
	ErrAlreadyClockedIn = errors.New("Gagal: Karyawan sudah melakukan Clock-In dan belum Clock-Out!")
	ErrNotClockedIn     = errors.New("Gagal: Karyawan belum melakukan Clock-In hari ini!")
)

type AttendanceService struct{ repository *repositories.Repository }

func NewAttendanceService(repository *repositories.Repository) *AttendanceService {
	return &AttendanceService{repository: repository}
}

func (s *AttendanceService) ClockIn(employeeID, note string) (models.Attendance, error) {
	if _, err := s.repository.GetEmployee(employeeID); errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Attendance{}, ErrEmployeeNotFound
	} else if err != nil {
		return models.Attendance{}, err
	}

	now := time.Now()
	last, err := s.repository.LatestAttendanceToday(employeeID, now)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Attendance{}, err
	}
	if err == nil && last.Type == "IN" {
		return models.Attendance{}, ErrAlreadyClockedIn
	}

	attendance := models.Attendance{ID: fmt.Sprintf("%d", now.UnixNano()), EmployeeID: employeeID, Type: "IN", Timestamp: now, Note: note}
	return attendance, s.repository.CreateAttendance(&attendance)
}

func (s *AttendanceService) ClockOut(employeeID string) (models.Attendance, error) {
	if _, err := s.repository.GetEmployee(employeeID); errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Attendance{}, ErrEmployeeNotFound
	} else if err != nil {
		return models.Attendance{}, err
	}

	now := time.Now()
	last, err := s.repository.LatestAttendanceToday(employeeID, now)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Attendance{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) || last.Type == "OUT" {
		return models.Attendance{}, ErrNotClockedIn
	}

	attendance := models.Attendance{ID: fmt.Sprintf("%d", now.UnixNano()), EmployeeID: employeeID, Type: "OUT", Timestamp: now}
	return attendance, s.repository.CreateAttendance(&attendance)
}

func (s *AttendanceService) List(date, employeeID string) ([]models.Attendance, error) {
	return s.repository.ListAttendances(date, employeeID)
}
