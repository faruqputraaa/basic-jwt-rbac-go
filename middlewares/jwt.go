package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("secret")

// Middleware untuk verifikasi JWT
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, echo.NewHTTPError(http.StatusUnauthorized, "Token tidak Valid")
            }
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, "Token tidak Valid")
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            return c.JSON(http.StatusUnauthorized, "Token tidak Valid")
        }

        // Simpan token ke dalam konteks Echo
        c.Set("userID", uint(claims["userID"].(float64)))
        c.Set("userToken", token) 

        return next(c)
    }
}


