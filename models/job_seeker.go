package models

type JobSeeker struct {
	JobSeekerID uint   `gorm:"primaryKey;column:job_seeker_id"`
	UserID      uint   `gorm:"unique;not null"`
	FullName    string `gorm:"not null"`
	ResumeURL   string

	User         User          `gorm:"foreignKey:UserID"`
	Applications []Application `gorm:"-"` // ignore for migrations
	Skills       []Skill       `gorm:"many2many:job_seeker_skills"`
}

func (JobSeeker) TableName() string {
	return "job_seekers"
}
