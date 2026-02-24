package controllers

import (
	"job-board-api/config"
	"job-board-api/models"
	"job-board-api/services"
	"job-board-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Make a user struct to bind the data
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
	FullName string `json:"full_name"`
	ResumeURL string `json:"resume_url"`
}

type LoginRequest struct {
    Email string `json:"email"`
    Password string `json:"password"`
}
//Bind Method:Bind binds path params, query params and the request body into provided type `i`. The default binder binds body based on Content-Type header.
// Our user Register function

func Register(c echo.Context) error {
    var req RegisterRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
    }

    // Call service layer
    err := services.RegisterUser(req.Email, req.Password, req.Role, req.FullName, req.ResumeURL)
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, echo.Map{"success": "User registered successfully"})
}

func Login(c echo.Context) error {
    //login logic
    //Bind the request with our login struct using bind method.
    var req LoginRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest , echo.Map{
            "message":"It is a Bad Request",
        })
    }

    var user models.User
    if err := config.DB.Where("email = ?",req.Email).First(&user).Error; err != nil {
        return c.JSON(http.StatusUnauthorized , echo.Map{
            "message":"Invalide Credentials",
        })
    }
    //verify the password
    if !utils.CheckPasswordHash(req.Password , user.PasswordHash) {
        return c.JSON(http.StatusUnauthorized , echo.Map{
            "message": "Invalide Credentials",
        })
    }
    token , err := utils.GenerateJWT(user.ID, user.Role)
    if err != nil {
        return c.JSON(http.StatusInternalServerError , echo.Map{
            "message":"Could not generate token due to server error",
        })
    }
    //return the token
    return c.JSON(http.StatusOK, echo.Map{
        "token":token,
    })
}