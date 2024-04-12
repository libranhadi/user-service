package model

import "fmt"

type User struct {
	Id        uint
	Firstname string
	Lastname  string
	Username  string
	Email     string
	Password  string
}

type UserRole struct {
	User User
	Role Role
}

func (u User) Validate() error {
	if len(u.Email) == 0 {
		return fmt.Errorf("email field is required")
	}

	if len(u.Password) == 0 {
		return fmt.Errorf("password field is required")
	}

	if len(u.Email) < 3 || len(u.Email) > 100 {
		return fmt.Errorf("email length must be between 3 and 100 characters")
	}

	if len(u.Password) < 6 {
		return fmt.Errorf("password must be bigger than 6 characters")
	}

	return nil
}
