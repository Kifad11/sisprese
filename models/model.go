package models

import "time"

type Employee struct {
	ID         string `json:"id" gorm:"primaryKey" binding:"required"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type Attendance struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	EmployeeID string    `json:"employee_id" gorm:"index"`
	Type       string    `json:"type"` // "IN" atau "OUT"
	Timestamp  time.Time `json:"timestamp" gorm:"index"`
	Note       string    `json:"note" gorm:"type:text"`
}
