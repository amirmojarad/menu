package service

type UserModel struct {
	Id             uint
	FirstName      string
	LastName       string
	Email          string
	HashedPassword string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	UserModel
}

type SignupRequest struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type SignupResponse struct {
	UserModel
}

type CreateUserRequest struct {
	UserModel
}

type CreateUserResponse struct {
	UserModel
}

type UserFilter struct {
	Id    uint
	Email string
}

type GetUserResponse struct {
	UserModel
}
