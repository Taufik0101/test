package service

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"test/api/dto"
	"test/api/model"
	"time"
)

type UserService interface {
	GetUser() ([]model.User, string)
	CreateUser(input dto.CreateUser) model.User
	PutUser(input dto.UpdateUser) model.User
	PatchUser(input dto.UpdateUser) model.User
	DeleteUser(id string) bool
}

type userService struct {
	userConnection *gorm.DB
	userCache      *redis.Client
}

func (u userService) GetUser() ([]model.User, string) {
	var users []model.User
	var dataFrom string

	//Check On Redis Cache User
	Get, _ := u.userCache.Get(context.Background(), "user").Result()

	// If On Redis Is "" Query From Database
	if Get == "" {
		dataFrom = "MySQL"
		u.userConnection.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: "id",
			},
			Desc: false,
		}).Find(&users)

		b, err := json.Marshal(&users)
		if err != nil {
			log.Println("Failed To Encode Struct")
		}

		_, _ = u.userCache.Set(context.Background(), "user", string(b), time.Hour*24).Result()
	} else {
		// If On Redis Not "", Get From Redis
		dataFrom = "Redis"
		in := []byte(Get)
		errMarshal := json.Unmarshal(in, &users)
		if errMarshal != nil {
			log.Print(errMarshal)
		}
	}

	return users, dataFrom
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
