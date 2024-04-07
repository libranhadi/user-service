package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"service-user/config"
	"service-user/model"
)

type IUserRepository interface {
	Create(user *model.User) error
	FindUserByEmail(email string) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (userRepo *UserRepository) Create(user *model.User) error {
	query := "INSERT INTO users (firstname, lastname, username , email, password) VALUES ($1,$2, $3, $4, $5)"
	ctx, cancel := config.NewContext()
	defer cancel()
	_, errExec := userRepo.db.ExecContext(ctx, query, &user.Firstname, &user.Lastname, &user.Username, &user.Email, &user.Password)
	if errExec != nil {
		// return errors.New("error, creating user")
		return fmt.Errorf("error, create user %w", errExec)
	}
	return nil
}

func (userRepo *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	query := "SELECT id, password, email FROM users WHERE email = $1"
	ctx, cancel := config.NewContext()
	defer cancel()

	user := &model.User{}

	rows, err := userRepo.db.QueryContext(ctx, query, email)
	if err != nil {
		return user, errors.New("error checking email existence")
		// return user, fmt.Errorf("error checking email existences: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Password, &user.Email)
		if err != nil {
			return user, errors.New("error scanning email existence result")
			// return user, fmt.Errorf("error scanning email existence result: %w", err)
		}
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (userRepo *UserRepository) FindUserByUsername(username string) (*model.User, error) {
	query := "SELECT id, username , password, email FROM users WHERE username = $1"
	ctx, cancel := config.NewContext()
	defer cancel()

	user := &model.User{}

	rows, err := userRepo.db.QueryContext(ctx, query, username)
	if err != nil {
		return user, errors.New("error checking username existence")
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return user, errors.New("error scanning username existence result")
		}
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}
