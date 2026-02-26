package middleware

import (
	"net/http"
	"strings"

	"job-board-api/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware verifies JWT token before allowing access to protected routes
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// Step 1: Get Authorization header from request
		authHeader := c.Request().Header.Get("Authorization")

		// If header is missing, return 401 Unauthorized
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Authorization header is required",
			})
		}

		// Step 2: Authorization header format should be:
		// "Bearer <token>"
		parts := strings.Split(authHeader, " ")

		// Check if format is correct
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid authorization format",
			})
		}

		// Extract actual token string
		tokenString := parts[1]

		// Step 3: Parse and verify token using secret key
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.JWTSecret, nil
		})

		// If parsing fails or token is invalid
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid or expired token",
			})
		}

		// Step 4: Extract claims (user_id and role) from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid token claims",
			})
		}

		// Step 5: Store user data in context
		// So controllers can access it later
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		// Step 6: Call next handler (actual controller)
		return next(c)
	}
}