package service

import (
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

func (bookService *BookService) ListBooks(page int, size int) (Page, []Book) {
	if page <= 0{
		page = 1
	}
	if size <= 0{
		size = 10
	}
	return bookMapper.ListBooks(page, size)
}

func (bookService *BookService) SaveBook(book *Book) (int, error) {
	book.BookId = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	book.Time = time.Time(book.ReleaseTime)
	return bookMapper.SaveBook(book)
}
