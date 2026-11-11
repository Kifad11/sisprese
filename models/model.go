package models

import "time"

type Employee struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type Attendance struct {
	ID         string    `json:"id"`
	EmployeeID string    `json:"employee_id"`
	Type       string    `json:"type"` // "IN" atau "OUT"
	Timestamp  time.Time `json:"timestamp"`
}