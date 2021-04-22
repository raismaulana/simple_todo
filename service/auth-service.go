package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/raismaulana/simple_todo/dto"
	"github.com/raismaulana/simple_todo/entity"
	"github.com/raismaulana/simple_todo/helper"
	"github.com/raismaulana/simple_todo/repository"
)

type AuthService interface {
	Login(input dto.LoginDTO) interface{}
	Register(input dto.RegisterDTO) entity.User
	IsDuplicateEmail(email string) bool
	GetRole(userID string) string
}

type authService struct {
	userRepository repository.UserRepository
}

func StaticAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) Login(input dto.LoginDTO) interface{} {
	resUser, err := service.userRepository.FindByEmail(input.Email)
	if err != nil {
		return false
	}
	if resUser.Email == input.Email && helper.PasswordVerify(resUser.Password, input.Password) {
		return resUser
	}
	return false
}

func (service *authService) Register(input dto.RegisterDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&input))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	resUser := service.userRepository.InsertUser(user)
	return resUser
}

// IsDuplicateEmail is a function that will return True if email already exists, otherwise False
func (service *authService) IsDuplicateEmail(email string) bool {
	_, err := service.userRepository.FindByEmail(email)
	if err != nil {
		return false
	}
	return true
}

func (service *authService) GetRole(userID string) string {
	user, err := service.userRepository.FindByID(userID)
	if err != nil {
		return "guest"
	}
	return user.Role
}
