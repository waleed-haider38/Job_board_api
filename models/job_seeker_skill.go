package models

type JobSeekerSkill struct {
	JobSeekerID uint `gorm:"primaryKey"`
	SkillID     uint `gorm:"primaryKey"`
}
