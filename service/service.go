package service

import (
	"zen/mapper"
	"zen/model"
)

type (
	BookService struct{}
	AuthorService struct{}
	Book = model.Book
	Author = model.Author
	Page = model.Page
)

var (
	bookMapper   mapper.BookMapper
	authorMapper mapper.AuthorMapper
)
