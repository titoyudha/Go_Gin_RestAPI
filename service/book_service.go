package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/titoyudha/Go_Gin_RestAPI/dto"
	"github.com/titoyudha/Go_Gin_RestAPI/entity"
	"github.com/titoyudha/Go_Gin_RestAPI/repository"
)

type BookService interface {
	Insert(b dto.BookCreatedDTO) entity.Book
	Update(b dto.BookUpdateDTO) entity.Book
	Delete(b entity.Book)
	GetAll() []entity.Book
	FindById(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b dto.BookCreatedDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed Map %v: ", err)
	}
	response := service.bookRepository.InsertBook(book)
	return response
}

func (service *bookService) Update(b dto.BookCreatedDTO) entity.Book {
	book := entity.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed Map %v: ", err)
	}
	response := service.bookRepository.UpdateBook(book)
	return response
}

func (service *bookService) Delete(b entity.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) GetAll() []entity.Book {
	return service.bookRepository.GetAllBook()
}

func (service *bookService) FindById(bookID uint64) entity.Book {
	return service.bookRepository.GetBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	book := service.bookRepository.GetBookByID(bookID)
	id := fmt.Sprintf("%v", book.UserID)
	return userID == id
}
