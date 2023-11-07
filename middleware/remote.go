package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func HandleRemoteAuthentication(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	authorization, ok := headers["Authorization"]
	if !ok {
		return fiber.NewError(http.StatusUnauthorized, "You are not authenticated.")
	}
	if authorization == "" || !strings.HasPrefix(authorization, "Bearer ") {
		return fiber.NewError(http.StatusUnauthorized, "Authorization was invalid.")
	}

	splitToken := strings.Split(authorization, "Bearer ")
	authorization = splitToken[1]

	if authorization != "ba884eea-ceee-4921-baba-3ae63496d482" {
		return fiber.NewError(http.StatusUnauthorized, "Authorization was incorrect.")
	}

	return c.Next()
}
