package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	Role         string    `gorm:"type:varchar(20);not null" json:"role"`
	// employer | job_seeker | admin

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Employer   *Employer   `gorm:"constraint:OnDelete:CASCADE;"`
	JobSeeker  *JobSeeker  `gorm:"constraint:OnDelete:CASCADE;"`
}
