package middleware

import (
	"net/http"
	"strings"

	"github.com/erdogancayir/nargileapi/domain"
	"github.com/erdogancayir/nargileapi/internal/tokenutil"
	"github.com/labstack/echo/v4"
)

func JwtAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")

			if len(t) == 2 {
				authToken := t[1]
				authorized, err := tokenutil.IsAuthorized(authToken, secret)

				if authorized {
					userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
					if err != nil {
						return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					}
					c.Set("x-user-id", userID)
					return next(c)
				}

				if err != nil {
					return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
				}
			}

			return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		}
	}
}
