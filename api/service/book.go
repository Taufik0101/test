package service

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"test/api/dto"
	"test/api/model"
)

type BookService interface {
	GetBook() []model.Book
	CreateBook(input dto.CreateBook) model.Book
	PutBook(input dto.UpdateBook) model.Book
	PatchBook(input dto.UpdateBook) model.Book
	DeleteBook(id string) bool
}

type bookService struct {
	bookConnection *gorm.DB
	bookCache      *redis.Client
}

func (b bookService) GetBook() []model.Book {
	//TODO implement me
	panic("implement me")
}

func (b bookService) CreateBook(input dto.CreateBook) model.Book {
	//TODO implement me
	panic("implement me")
}

func (b bookService) PutBook(input dto.UpdateBook) model.Book {
	//TODO implement me
	panic("implement me")
}

func (b bookService) PatchBook(input dto.UpdateBook) model.Book {
	//TODO implement me
	panic("implement me")
}

func (b bookService) DeleteBook(id string) bool {
	//TODO implement me
	panic("implement me")
}

func NewBookService(bookConn *gorm.DB, bookCaches *redis.Client) BookService {
	return &bookService{
		bookConnection: bookConn,
		bookCache:      bookCaches,
	}
}
