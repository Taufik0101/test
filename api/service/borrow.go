package service

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"test/api/dto"
	"test/api/model"
)

type BorrowService interface {
	GetBorrow() []model.Borrow
	CreateBorrow(input dto.CreateBorrow) model.Borrow
	PutBorrow(input dto.UpdateBorrow) model.Borrow
	PatchBorrow(input dto.UpdateBorrow) model.Borrow
	DeleteBorrow(id string) bool
}

type borrowService struct {
	borrowConnection *gorm.DB
	borrowCache      *redis.Client
}

func (b borrowService) GetBorrow() []model.Borrow {
	//TODO implement me
	panic("implement me")
}

func (b borrowService) CreateBorrow(input dto.CreateBorrow) model.Borrow {
	//TODO implement me
	panic("implement me")
}

func (b borrowService) PutBorrow(input dto.UpdateBorrow) model.Borrow {
	//TODO implement me
	panic("implement me")
}

func (b borrowService) PatchBorrow(input dto.UpdateBorrow) model.Borrow {
	//TODO implement me
	panic("implement me")
}

func (b borrowService) DeleteBorrow(id string) bool {
	//TODO implement me
	panic("implement me")
}

func NewBorrowService(borrowConn *gorm.DB, borrowCaches *redis.Client) BorrowService {
	return &borrowService{
		borrowConnection: borrowConn,
		borrowCache:      borrowCaches,
	}
}
