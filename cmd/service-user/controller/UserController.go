package controller

import (
	"net/http"
	"service-user/helpers"
	"service-user/model"
	"service-user/service"

	"github.com/gofiber/fiber/v2"
)

type IUserController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Auth(c *fiber.Ctx) error
}

type userController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &userController{userService: userService}
}

func (uc *userController) Register(c *fiber.Ctx) error {
	requestBody := new(model.User)
	if err := c.BodyParser(requestBody); err != nil {
		return c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}

	err := uc.userService.Register(requestBody)
	if err != nil {
		webResponse, ok := err.(*helpers.WebResponse)
		if ok {
			return c.Status(webResponse.Code).JSON(webResponse)
		} else {
			return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
				Code:    http.StatusInternalServerError,
				Status:  "Internal server error",
				Message: "Sorry, something went wrong please try again later!",
			})
		}
	} else {
		return c.JSON(&helpers.WebResponse{
			Code:    201,
			Status:  "OK",
			Data:    requestBody.Email,
			Message: "Thank you for signing up! You can now log in to your account!",
		})
	}
}

func (uc *userController) Login(c *fiber.Ctx) error {
	requestBody := new(model.User)
	if err := c.BodyParser(requestBody); err != nil {
		return c.JSON(&helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid data request",
		})
	}
	resp, err := uc.userService.Login(requestBody)

	if err != nil {
		webResponse, ok := err.(*helpers.WebResponse)
		if ok {
			return c.Status(webResponse.Code).JSON(webResponse)
		} else {
			return c.Status(http.StatusInternalServerError).JSON(&helpers.WebResponse{
				Code:    http.StatusInternalServerError,
				Status:  "Internal server error",
				Message: "Sorry, something went wrong please try again later!",
			})
		}
	}

	access_token := helpers.SignToken(requestBody.Email)

	return c.JSON(struct {
		Code        int
		Status      string
		AccessToken string
		Message     string
		Data        struct {
			ID       uint   `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"Data"`
	}{
		Code:        200,
		Status:      "OK",
		AccessToken: access_token,
		Message:     "Successfully login !",
		Data: struct {
			ID       uint   `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			ID:       resp.Id,
			Email:    resp.Email,
			Username: resp.Username,
		},
	})
}

func (uc *userController) Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}
