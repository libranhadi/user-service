package middleware

import (
	"net/http"
	"service-user/helpers"
	"service-user/repository"

	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	Authentication(c *fiber.Ctx) error
}

type authImpl struct {
	userRepository repository.IUserRepository
}

func NewAuthImpl(repository repository.IUserRepository) Auth {
	return &authImpl{userRepository: repository}
}

func (auth *authImpl) Authentication(c *fiber.Ctx) error {
	access_token := c.Get("Authorization")
	tokenString := ""
	if len(access_token) > len("Bearer ") {
		tokenString = access_token[len("Bearer "):]
	}

	if len(tokenString) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	checkToken, err := helpers.VerifyToken(tokenString)

	if err != nil {
		return c.Status(401).SendString("Invalid token: Failed to verify token")
	}

	email, ok := checkToken["email"].(string)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Sorry, something went wrong please ",
		})
	}

	user, err := auth.userRepository.FindUserByEmail(email)
	if user == nil {
		return c.Status(http.StatusNotFound).JSON(&helpers.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: err.Error(),
		})
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: err.Error(),
		})
	}

	c.Locals("user", user)

	return c.Next()
}
