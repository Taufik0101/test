package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"test/api/dto"
	"test/api/model"
	"time"
)

type UserService interface {
	GetUser() ([]*model.User, string)
	CreateUser(input dto.CreateUser) (*model.User, error)
	PutUser(input dto.PutUser) (*model.User, error)
	PatchUser(input dto.PatchUser) (*model.User, error)
	DeleteUser(id string) (bool, error)
}

type userService struct {
	userConnection *gorm.DB
	userCache      *redis.Client
}

func (u userService) GetUser() ([]*model.User, string) {
	//TODO implement me
	var users []*model.User
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

func (u userService) CreateUser(input dto.CreateUser) (*model.User, error) {
	//TODO implement me
	newUser := &model.User{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	errSave := u.userConnection.Create(newUser).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = u.userCache.Del(context.Background(), "user").Result()

	return newUser, nil
}

func (u userService) PutUser(input dto.PutUser) (*model.User, error) {
	//TODO implement me
	newUser := &model.User{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	errSave := u.userConnection.Where(model.User{ID: input.ID}).Updates(newUser).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = u.userCache.Del(context.Background(), "user").Result()

	var NewUser *model.User
	errGet := u.userConnection.Where(model.User{ID: input.ID}).First(&NewUser).Error
	if errGet != nil {
		return nil, errGet
	}

	return NewUser, nil
}

func (u userService) PatchUser(input dto.PatchUser) (*model.User, error) {
	//TODO implement me
	newUser := &model.User{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	errSave := u.userConnection.Where(model.User{ID: input.ID}).Updates(newUser).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = u.userCache.Del(context.Background(), "user").Result()

	var NewUser *model.User
	errGet := u.userConnection.Where(model.User{ID: input.ID}).First(&NewUser).Error
	if errGet != nil {
		return nil, errGet
	}

	return NewUser, nil
}

func (u userService) DeleteUser(id string) (bool, error) {
	//TODO implement me
	var user *model.User
	errDelete := u.userConnection.Where(model.User{ID: id}).Delete(&user).Error

	if errDelete != nil {
		return false, errDelete
	}

	_, _ = u.userCache.Del(context.Background(), "user").Result()

	return true, nil
}

func NewUserService(userConn *gorm.DB, userCaches *redis.Client) UserService {
	return &userService{
		userConnection: userConn,
		userCache:      userCaches,
	}
}
