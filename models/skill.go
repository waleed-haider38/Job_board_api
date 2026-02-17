package models

type Skill struct {
	SkillID   uint   `gorm:"primaryKey;column:skill_id"`
	SkillName string `gorm:"unique;not null"`
}
