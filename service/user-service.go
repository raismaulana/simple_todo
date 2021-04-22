package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/raismaulana/simple_todo/dto"
	"github.com/raismaulana/simple_todo/entity"
	"github.com/raismaulana/simple_todo/repository"
)

type UserService interface {
	UpdateUser(userUpdateDTO dto.UserUpdateDTO) entity.User
	GetAll() []entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func StaticUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) UpdateUser(input dto.UserUpdateDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&input))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	resUser := service.userRepository.UpdateUser(user)
	return resUser
}

func (service *userService) GetAll() []entity.User {
	resUser := service.userRepository.GetAll()
	return resUser
}
