package models

type JobSkill struct {
	JobID   uint `gorm:"primaryKey"`
	SkillID uint `gorm:"primaryKey"`
}
