package service

import (
	"net/http"
	"service-user/helpers"
	"service-user/model"
	"service-user/repository"
)

type IUserService interface {
	Register(user *model.User) error
	Login(user *model.User) (*model.User, error)
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{userRepository: repository}
}

func (service *UserService) Register(user *model.User) error {

	if err := user.Validate(); err != nil {
		return &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    user,
		}
	}

	userEmail, err := service.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		if err.Error() != "user not found" {
			return &helpers.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: err.Error(),
			}
		}
	}

	if userEmail != nil {
		return &helpers.WebResponse{
			Code:    http.StatusConflict,
			Status:  "Conflict",
			Message: "Email already exists",
		}
	}

	userName, err := service.userRepository.FindUserByUsername(user.Username)
	if err != nil {
		if err.Error() != "user not found" {
			return &helpers.WebResponse{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: err.Error(),
			}
		}
	}

	if userName != nil {
		return &helpers.WebResponse{
			Code:    http.StatusConflict,
			Status:  "Conflict",
			Message: "Username already exists",
		}
	}

	password := helpers.HashPassword([]byte(user.Password))
	user.Password = password
	return service.userRepository.Create(user)
}

func (userService *UserService) Login(user *model.User) (*model.User, error) {

	if err := user.Validate(); err != nil {
		return user, &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
			Data:    nil,
		}
	}

	userExist, err := userService.userRepository.FindUserByEmail(user.Email)
	if err != nil {
		return nil, &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		}
	}

	if userExist == nil {
		return userExist, &helpers.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: err.Error(),
		}
	}

	checkPassword := helpers.ComparePassword([]byte(userExist.Password), []byte(user.Password))
	if !checkPassword {
		return userExist, &helpers.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Data:    nil,
			Message: "Password doesn't match!",
		}
	}

	return userExist, nil
}
