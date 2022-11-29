package service

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"test/api/dto"
	"test/api/model"
)

type UserService interface {
	GetUser() []model.User
	CreateUser(input dto.CreateUser) model.User
	PutUser(input dto.UpdateUser) model.User
	PatchUser(input dto.UpdateUser) model.User
	DeleteUser(id string) bool
}

type userService struct {
	userConnection *gorm.DB
	userCache      *redis.Client
}

func (u userService) GetUser() []model.User {
	//TODO implement me
	panic("implement me")
}

func (u userService) CreateUser(input dto.CreateUser) model.User {
	//TODO implement me
	panic("implement me")
}

func (u userService) PutUser(input dto.UpdateUser) model.User {
	//TODO implement me
	panic("implement me")
}

func (u userService) PatchUser(input dto.UpdateUser) model.User {
	//TODO implement me
	panic("implement me")
}

func (u userService) DeleteUser(id string) bool {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userConn *gorm.DB, userCaches *redis.Client) UserService {
	return &userService{
		userConnection: userConn,
		userCache:      userCaches,
	}
}
