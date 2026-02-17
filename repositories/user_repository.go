package repositories

import (
	"job-board-api/config"
	"job-board-api/models"
)

//Making these functions to seperate the database logic.
func IsEmailExists(email string) bool {
    var user models.User
    if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return false
    }
    return true
}

func CreateUser(email, passwordHash, role, fullName, resumeUrl string) error {
    // 1 First create user
    user := models.User{
        Email: email,
        PasswordHash: passwordHash,
        Role: role,
    }

    if err := config.DB.Create(&user).Error; err != nil {
        return err
    }

    // 2 Now create reference table record based on role
    switch role {
    case "employee":
        employee := models.Employer{ // Agar employee ka table Employer hai
            UserID: user.ID,
            FullName:fullName,
        }
        if err := config.DB.Create(&employee).Error; err != nil {
            return err
        }

    case "job_seeker":
        jobSeeker := models.JobSeeker{
            UserID: user.ID,
			FullName: fullName,
			ResumeURL: resumeUrl,
        }
        if err := config.DB.Create(&jobSeeker).Error; err != nil {
            return err
        }
    }

    return nil
}
