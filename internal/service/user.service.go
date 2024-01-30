package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"menu_generator/config"
	"menu_generator/internal/errorext"
)

type UserRepository interface {
	Create(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error)
	Get(ctx context.Context, req UserFilter) (GetUserResponse, error)
}

type User struct {
	userRepository UserRepository
	cfg            *config.AppConfig
	logger         *logrus.Entry
}

func NewUser(cfg *config.AppConfig, logger *logrus.Entry, userRepository UserRepository) *User {
	return &User{
		userRepository: userRepository,
		cfg:            cfg,
		logger:         logger,
	}
}

func (u User) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	if !req.Validate() {
		return LoginResponse{}, errorext.NewValidationError("invalid password or email")
	}

	getUserResponse, err := u.userRepository.Get(ctx, UserFilter{
		Email: req.Email,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	if err = VerifyPassword(req.Password, getUserResponse.HashedPassword); err != nil {
		return LoginResponse{}, errorext.NewValidationError("wrong password: %+v", err)
	}

	return LoginResponse(getUserResponse), nil
}

func (u User) Signup(ctx context.Context, req SignupRequest) (SignupResponse, error) {
	if !req.Validate() {
		return SignupResponse{}, errorext.NewValidationError("invalid password or email")
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return SignupResponse{}, errorext.NewValidationError("password hashing error: %+v", err)
	}

	createUserRequest := CreateUserRequest{
		UserModel{
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Email:          req.Email,
			HashedPassword: hashedPassword,
		},
	}

	createUserResponse, err := u.userRepository.Create(ctx, createUserRequest)
	if err != nil {
		return SignupResponse{}, err
	}

	return SignupResponse(createUserResponse), nil
}
