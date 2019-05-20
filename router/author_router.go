package router

import (
	"net/http"
	"strconv"
	"zen/util"
)

func (authorRouter *AuthorRouter) Setup(engine *Engine) {
	group := engine.Group("/authors")
	group.POST("/", authorRouter.saveAuthor)
	group.GET("/", authorRouter.listAuthors)
}

func (authorRouter *AuthorRouter) saveAuthor(ctx *Context) {
	var author Author
	if err := ctx.ShouldBind(&author); nil != err {
		log.WithFields(util.BeanToMap(author)).Error("author cant be saved, maybe cause: ", err.Error())
		ctx.IndentedJSON(http.StatusOK, util.ErrWithMessage(err.Error(), 0))
		return
	}
	if count, err := authorService.SaveAuthor(&author); nil != err {
		log.WithFields(util.BeanToMap(author)).Error("author cant be saved, maybe cause: ", err.Error())
		ctx.IndentedJSON(http.StatusOK, util.ErrWithMessage(err.Error(), 0))
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, util.Ok(count))
	}
}


func (authorRouter * AuthorRouter) listAuthors(ctx * Context){
	page := 1
	size := 10
	var err error
	var condition = ctx.Query("page")
	if page, err = strconv.Atoi(condition); nil != err{
		log.WithFields(map[string]interface{}{"page": condition, "error": err}).
			Warn("error parameter, page will be set with 1")
		page = 1
	}
	condition = ctx.Query("size")
	if size, err = strconv.Atoi(condition); nil != err{
		log.WithFields(map[string]interface{}{"size": condition, "error": err}).
			Warn("error parameter, size will be set with 10")
		size = 10
	}
	ctx.IndentedJSON(http.StatusOK, util.Page(authorService.ListAuthors(page, size)))
}