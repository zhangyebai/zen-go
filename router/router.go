package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"zen/model"
	"zen/service"
)

type (
	BookRouter struct{}
	AuthorRouter struct{}
	Engine = gin.Engine
	Context = gin.Context
	Book = model.Book
	Author = model.Author
)

var (
	log           = logrus.New()
	bookService   service.BookService
	authorService service.AuthorService
)

func Setup(engine *Engine) {
	new(BookRouter).Setup(engine)
	new(AuthorRouter).Setup(engine)
}
