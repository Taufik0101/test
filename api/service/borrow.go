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

type BorrowService interface {
	GetBorrow() ([]*model.Borrow, string)
	CreateBorrow(input dto.CreateBorrow) (*model.Borrow, error)
	PutBorrow(input dto.PutBorrow) (*model.Borrow, error)
	PatchBorrow(input dto.PatchBorrow) (*model.Borrow, error)
	DeleteBorrow(id string) (bool, error)
}

type borrowService struct {
	borrowConnection *gorm.DB
	borrowCache      *redis.Client
}

func (b borrowService) GetBorrow() ([]*model.Borrow, string) {
	//TODO implement me
	var borrows []*model.Borrow
	var dataFrom string

	//Check On Redis Cache Borrow
	Get, _ := b.borrowCache.Get(context.Background(), "borrow").Result()

	// If On Redis Is "" Query From Database
	if Get == "" {
		dataFrom = "MySQL"
		b.borrowConnection.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: "created_at",
			},
			Desc: true,
		}).Preload("Users").Preload("Books").Find(&borrows)

		bes, err := json.Marshal(&borrows)
		if err != nil {
			log.Println("Failed To Encode Struct")
		}

		_, _ = b.borrowCache.Set(context.Background(), "borrow", string(bes), time.Hour*24).Result()
	} else {
		// If On Redis Not "", Get From Redis
		dataFrom = "Redis"
		in := []byte(Get)
		errMarshal := json.Unmarshal(in, &borrows)
		if errMarshal != nil {
			log.Print(errMarshal)
		}
	}

	return borrows, dataFrom
}

func (b borrowService) CreateBorrow(input dto.CreateBorrow) (*model.Borrow, error) {
	//TODO implement me
	newBorrow := &model.Borrow{
		UserID: input.UserID,
		BookID: input.BookID,
	}

	errSave := b.borrowConnection.Create(newBorrow).Error

	if errSave != nil {
		return nil, errSave
	}

	var NewBorrow *model.Borrow
	b.borrowConnection.Model(&model.Borrow{}).Where(model.Borrow{ID: newBorrow.ID}).
		Preload("Users").Preload("Books").First(&NewBorrow)

	_, _ = b.borrowCache.Del(context.Background(), "borrow").Result()

	return NewBorrow, nil
}

func (b borrowService) PutBorrow(input dto.PutBorrow) (*model.Borrow, error) {
	//TODO implement me
	newBorrow := &model.Borrow{
		UserID: input.UserID,
		BookID: input.BookID,
	}

	errSave := b.borrowConnection.Where(model.Borrow{ID: input.ID}).Updates(newBorrow).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = b.borrowCache.Del(context.Background(), "borrow").Result()

	var NewBorrow *model.Borrow
	errGet := b.borrowConnection.Where(model.Borrow{ID: input.ID}).Preload("Users").Preload("Books").First(&NewBorrow).Error
	if errGet != nil {
		return nil, errGet
	}

	return NewBorrow, nil
}

func (b borrowService) PatchBorrow(input dto.PatchBorrow) (*model.Borrow, error) {
	//TODO implement me
	newBorrow := &model.Borrow{
		UserID: input.UserID,
		BookID: input.BookID,
	}

	errSave := b.borrowConnection.Where(model.Borrow{ID: input.ID}).Updates(newBorrow).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = b.borrowCache.Del(context.Background(), "borrow").Result()

	var NewBorrow *model.Borrow
	errGet := b.borrowConnection.Where(model.Borrow{ID: input.ID}).Preload("Users").Preload("Books").First(&NewBorrow).Error
	if errGet != nil {
		return nil, errGet
	}

	return NewBorrow, nil
}

func (b borrowService) DeleteBorrow(id string) (bool, error) {
	//TODO implement me
	var borrow *model.Borrow
	errDelete := b.borrowConnection.Where(model.Borrow{ID: id}).Delete(&borrow).Error

	if errDelete != nil {
		return false, errDelete
	}

	_, _ = b.borrowCache.Del(context.Background(), "borrow").Result()

	return true, nil
}

func NewBorrowService(borrowConn *gorm.DB, borrowCaches *redis.Client) BorrowService {
	return &borrowService{
		borrowConnection: borrowConn,
		borrowCache:      borrowCaches,
	}
}
