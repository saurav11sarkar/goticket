package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/saurav11sarkar/ticket/internal/auth"
	"github.com/saurav11sarkar/ticket/internal/httpResponse"
)

func AuthMiddleware(jwtService auth.JwtService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// 1. Authorization header
			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, httpResponse.Error{
					Code:    http.StatusUnauthorized,
					Message: http.StatusText(http.StatusUnauthorized),
				})
			}
			// 2. Bearer token check
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, httpResponse.Error{
					Code:    http.StatusUnauthorized,
					Message: http.StatusText(http.StatusUnauthorized),
				})
			}
			tokenString := parts[1]
			// 3. JWT validate
			token, err := jwtService.ValidateToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, httpResponse.Error{
					Code:    http.StatusUnauthorized,
					Message: http.StatusText(http.StatusUnauthorized),
				})
			}
			// 4. User information context
			c.Set("id", token.ID)
			c.Set("email", token.Email)
			c.Set("name", token.Name)
			// 5. Next handler
			return next(c)
		}
	}
}
