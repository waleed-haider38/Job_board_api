package controllers

import (
	"job-board-api/services"
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

