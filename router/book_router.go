package router

import (
	"net/http"
	"strconv"
	"zen/util"
)

func (bookRouter *BookRouter) Setup(engine *Engine) {
	group := engine.Group("/books")
	group.GET("/", bookRouter.listBooks)
	group.POST("/", bookRouter.saveBook)
}

func (bookRouter *BookRouter) saveBook(ctx *Context) {
	var book Book
	if err := ctx.ShouldBind(&book); nil != err {
		ctx.IndentedJSON(http.StatusOK, util.ErrWithMessage(err.Error(), 0))
		return
	}
	if count, err := bookService.SaveBook(&book); nil != err {
		ctx.IndentedJSON(http.StatusOK, util.ErrWithMessage(err.Error(), 0))
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, util.Ok(count))
	}
}

func (bookRouter *BookRouter) listBooks(ctx *Context) {
	var page = 1
	var size = 10
	var err error
	condition := ctx.Query("page")
	if page, err = strconv.Atoi(condition); nil != err {
		log.WithFields(map[string]interface{}{"page": condition, "error": err}).
			Warn("error parameter, page will be set with 1")
		page = 1
	}
	condition = ctx.Query("size")
	if size, err = strconv.Atoi(condition); nil != err {
		log.WithFields(map[string]interface{}{"size": condition, "error": err}).
			Warn("error parameter, size will be set with 10")
		size = 10
	}
	ctx.IndentedJSON(http.StatusOK, util.Page(bookService.ListBooks(page, size)))
}
