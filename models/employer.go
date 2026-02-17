package models

type Employer struct {
	EmployerID uint   `gorm:"primaryKey;column:employer_id"`
	UserID     uint   `gorm:"unique;not null"`
	FullName   string `gorm:" null"`

	User      User      `gorm:"foreignKey:UserID"`
	Companies []Company `gorm:"-"`
}
