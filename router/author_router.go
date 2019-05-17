package router

import (
	"net/http"
	"zen/util"
)

func (authorRouter *AuthorRouter) Setup(engine *Engine) {
	group := engine.Group("/authors")
	group.POST("/", authorRouter.saveAuthor)
}

func (authorRouter *AuthorRouter) saveAuthor(ctx *Context) {
	var author Author
	if err := ctx.ShouldBind(&author); nil != err {
		log.WithFields(util.BeanToMap(author)).Error("author cant be saved, maybe cause: ", err.Error())
		ctx.IndentedJSON(http.StatusOK, util.ErrWithMessage(err.Error(), 0))
	}
	if count, err := authorService.SaveAuthor(&author); nil != err {
		log.WithFields(util.BeanToMap(author)).Error("author cant be saved, maybe cause: ", err.Error())
		ctx.IndentedJSON(http.StatusOK, util.ErrWithMessage(err.Error(), 0))
	} else {
		ctx.IndentedJSON(http.StatusOK, util.Ok(count))
	}
}
