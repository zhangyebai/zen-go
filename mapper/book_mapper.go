package mapper

import "zen/model"

func (bookMapper *BookMapper) ListBooks(page int, size int) (Page, []Book) {
	var books []Book
	var total int
	var offset = 0
	if size <= 0 {
		size = 10
	}
	if page <= 0 {
		page = 1
	}
	offset = (page - 1) * size
	db.Table("book").Count(&total).Limit(size).Offset(offset).Find(&books)
	var pages = 0
	if total > 0 {
		pages = total/size + 1
	}
	for idx := range books {
		books[idx].ReleaseTime = model.JsonTime(books[idx].Time)
	}
	return Page{Page: page, Size: len(books), Total: total, Pages: pages}, books
}

func (bookMapper *BookMapper) SaveBook(book *Book) (int, error) {
	ctx := db.Begin()
	if err := ctx.Create(book).Error; nil != err {
		ctx.Rollback()
		return 0, err
	}
	ctx.Commit()
	return 1, nil
}
