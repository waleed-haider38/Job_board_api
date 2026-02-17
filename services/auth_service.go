package services

import (
	"fmt"
	"job-board-api/repositories"
	"job-board-api/utils"
)

func RegisterUser(email, password, role, fullName, resumeUrl string) error {
    // Check if user exists
    exists := repositories.IsEmailExists(email)
    if exists {
        return fmt.Errorf("Email already exists")
    }

    // Hash password
    hashed, err := utils.HashPassword(password)
    if err != nil {
        return err
    }

    // Create user
    return repositories.CreateUser(email, hashed, role,fullName, resumeUrl)
}
