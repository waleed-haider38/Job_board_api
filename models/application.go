package models

import "time"

type Application struct {
	ApplicationID uint      `gorm:"primaryKey;column:application_id"`
	JobID         uint      `gorm:"not null;uniqueIndex:idx_job_seeker_job"`
	JobSeekerID   uint      `gorm:"not null;uniqueIndex:idx_job_seeker_job"`
	CoverLetter   string    `gorm:"type:text"`
	Status        string    `gorm:"default:'pending'"`
	AppliedAt     time.Time `gorm:"autoCreateTime"`

	Job       Job       `gorm:"-"`       // FK from Application to Job
	JobSeeker JobSeeker `gorm:"-"` // FK from Application to JobSeeker
}


func (Application) TableName() string {
	return "applications"
}
