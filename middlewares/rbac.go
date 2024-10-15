package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func RoleMiddleware(requiredRole string) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            token := c.Get("userToken").(*jwt.Token) // Ambil token dari konteks
            claims := token.Claims.(jwt.MapClaims)
            role := claims["role"].(string)

            if role != requiredRole {
                return c.JSON(http.StatusForbidden, "Akses Ditolak")
            }
            return next(c)
        }
    }
}


