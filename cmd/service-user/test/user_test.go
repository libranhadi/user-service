package test

import (
	"errors"
	"service-user/helpers"
	"service-user/model"
	"service-user/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repoMock *UserRepositoryMock) FindUserByEmail(email string) (*model.User, error) {
	args := repoMock.Mock.Called(email)
	if args.Get(0) == nil {
		return nil, errors.New("user not found")
	}
	user := args.Get(0).(*model.User)
	return user, nil
}

func (repoMock *UserRepositoryMock) Create(user *model.User) error {
	args := repoMock.Mock.Called(user)
	return args.Error(0)
}

func TestRegisterWithMockRepository_Success(t *testing.T) {
	mockRepo := new(UserRepositoryMock)
	user := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}

	mockRepo.On("FindUserByEmail", user.Email).Return(nil, errors.New("user not found"))
	mockRepo.On("Create", user).Return(nil)
	serviceMock := service.NewUserService(mockRepo)
	err := serviceMock.Register(user)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegisterWithMockRepository_Failed(t *testing.T) {
	mockRepo := new(UserRepositoryMock)
	user := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}
	objError := errors.New("email already exist")
	mockRepo.On("FindUserByEmail", user.Email).Return(nil, objError)
	mockRepo.On("Create", user).Return(objError)
	serviceMock := service.NewUserService(mockRepo)
	err := serviceMock.Register(user)
	assert.EqualError(t, err, objError.Error())
	mockRepo.AssertExpectations(t)
}

func TestLoginWithMockRepository_Success(t *testing.T) {
	mockRepo := new(UserRepositoryMock)

	userExist := &model.User{
		Email:    "email_test_1@example.com",
		Password: helpers.HashPassword([]byte("password")),
	}

	mockRepo.On("FindUserByEmail", userExist.Email).Return(userExist, nil)

	userLogin := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}

	serviceMock := service.NewUserService(mockRepo)
	_, err := serviceMock.Login(userLogin)
	assert.Nil(t, err)
	assert.Equal(t, userExist.Email, userLogin.Email)
	mockRepo.AssertExpectations(t)
}

func TestLoginWithMockRepository_Failed(t *testing.T) {
	mockRepo := new(UserRepositoryMock)

	userLogin := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}
	objError := errors.New("user not found")
	mockRepo.On("FindUserByEmail", userLogin.Email).Return(nil, objError)

	serviceMock := service.NewUserService(mockRepo)
	_, err := serviceMock.Login(userLogin)
	assert.EqualError(t, err, objError.Error())
	mockRepo.AssertExpectations(t)

}

func TestLoginWithMockRepository_Failed_PasswordNotMatch(t *testing.T) {
	mockRepo := new(UserRepositoryMock)

	userExist := &model.User{
		Email:    "email_test_1@example.com",
		Password: helpers.HashPassword([]byte("test")),
	}

	objError := errors.New("Password doesn't match!")

	mockRepo.On("FindUserByEmail", userExist.Email).Return(userExist, objError)

	userLogin := &model.User{
		Email:    "email_test_1@example.com",
		Password: "password",
	}

	serviceMock := service.NewUserService(mockRepo)
	_, err := serviceMock.Login(userLogin)
	assert.EqualError(t, err, objError.Error())
	mockRepo.AssertExpectations(t)
}
