package models

type Job struct {
	JobID       uint   `gorm:"primaryKey;column:job_id"`
	CompanyID   uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text;not null"`
	JobLocation string
	JobType     string
	SalaryMin   int
	SalaryMax   int
	Status      string `gorm:"default:'open'"`

	Company      Company       `gorm:"-"`
	Applications []Application `gorm:"-"` // ignore for migration
	Skills       []Skill       `gorm:"-"` // ignore for migration
}

func (Job) TableName() string {
	return "jobs"
}
