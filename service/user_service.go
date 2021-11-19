package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/titoyudha/Go_Gin_RestAPI/dto"
	"github.com/titoyudha/Go_Gin_RestAPI/entity"
	"github.com/titoyudha/Go_Gin_RestAPI/repository"
)

//Contract User Service
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

//new user service created new instance of User Service
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

//UPDATE USER
func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(user))
	if err != nil {
		log.Fatalf("Failed mapping %v", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
