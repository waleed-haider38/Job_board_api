package models

type Company struct {
	CompanyID      uint   `gorm:"primaryKey;column:company_id"`
	EmployerID     uint   `gorm:"not null"`
	CompanyName    string `gorm:"not null"`
	CompanyProduct string

	Employer Employer `gorm:"-"`
	Jobs     []Job    `gorm:"-"`
}
