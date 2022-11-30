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

type BookService interface {
	GetBook() ([]*model.Book, string)
	CreateBook(input dto.CreateBook) (*model.Book, error)
	PutBook(input dto.PutBook) (*model.Book, error)
	PatchBook(input dto.PatchBook) (*model.Book, error)
	DeleteBook(id string) (bool, error)
}

type bookService struct {
	bookConnection *gorm.DB
	bookCache      *redis.Client
}

func (b bookService) GetBook() ([]*model.Book, string) {
	//TODO implement me
	var books []*model.Book
	var dataFrom string

	//Check On Redis Cache Book
	Get, _ := b.bookCache.Get(context.Background(), "book").Result()

	// If On Redis Is "" Query From Database
	if Get == "" {
		dataFrom = "MySQL"
		b.bookConnection.Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: "created_at",
			},
			Desc: true,
		}).Find(&books)

		bes, err := json.Marshal(&books)
		if err != nil {
			log.Println("Failed To Encode Struct")
		}

		_, _ = b.bookCache.Set(context.Background(), "book", string(bes), time.Hour*24).Result()
	} else {
		// If On Redis Not "", Get From Redis
		dataFrom = "Redis"
		in := []byte(Get)
		errMarshal := json.Unmarshal(in, &books)
		if errMarshal != nil {
			log.Print(errMarshal)
		}
	}

	return books, dataFrom
}

func (b bookService) CreateBook(input dto.CreateBook) (*model.Book, error) {
	//TODO implement me
	newBook := &model.Book{
		Name: input.Name,
		Page: input.Page,
	}

	errSave := b.bookConnection.Create(newBook).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = b.bookCache.Del(context.Background(), "book").Result()

	return newBook, nil
}

func (b bookService) PutBook(input dto.PutBook) (*model.Book, error) {
	//TODO implement me
	newBook := &model.Book{
		Name: input.Name,
		Page: input.Page,
	}

	errSave := b.bookConnection.Where(model.Book{ID: input.ID}).Updates(newBook).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = b.bookCache.Del(context.Background(), "book").Result()

	var NewBook *model.Book
	errGet := b.bookConnection.Where(model.Book{ID: input.ID}).First(&NewBook).Error
	if errGet != nil {
		return nil, errGet
	}

	return NewBook, nil
}

func (b bookService) PatchBook(input dto.PatchBook) (*model.Book, error) {
	//TODO implement me
	newBook := &model.Book{
		Name: input.Name,
		Page: input.Page,
	}

	errSave := b.bookConnection.Where(model.Book{ID: input.ID}).Updates(newBook).Error

	if errSave != nil {
		return nil, errSave
	}

	_, _ = b.bookCache.Del(context.Background(), "book").Result()

	var NewBook *model.Book
	errGet := b.bookConnection.Where(model.Book{ID: input.ID}).First(&NewBook).Error
	if errGet != nil {
		return nil, errGet
	}

	return NewBook, nil
}

func (b bookService) DeleteBook(id string) (bool, error) {
	//TODO implement me
	var book *model.Book
	errDelete := b.bookConnection.Where(model.Book{ID: id}).Delete(&book).Error

	if errDelete != nil {
		return false, errDelete
	}

	_, _ = b.bookCache.Del(context.Background(), "book").Result()

	return true, nil
}

func NewBookService(bookConn *gorm.DB, bookCaches *redis.Client) BookService {
	return &bookService{
		bookConnection: bookConn,
		bookCache:      bookCaches,
	}
}
