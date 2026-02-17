package migrations

import (
	"log"

	"job-board-api/config"
	"job-board-api/models"
)

func RunMigrations() {

	err := config.DB.AutoMigrate(
		&models.User{},           // 1
		&models.Employer{},       // 2 → depends on User
		&models.Company{},        // 3 → depends on Employer
		&models.Job{},            // 4 → depends on Company
		&models.Skill{},          // 5
		&models.JobSeeker{},      // 6 → depends on User
		&models.Application{},    // 7 → depends on Job + JobSeeker
		&models.JobSkill{},       // 8 → many2many Job ↔ Skill
		&models.JobSeekerSkill{}, // 9 → many2many JobSeeker ↔ Skill
	)






	if err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	log.Println("Database migrated successfully")
}
