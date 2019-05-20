package mapper

func (authorMapper *AuthorMapper) SaveAuthor(author *Author) (int, error) {
	ctx := db.Begin()
	if err := ctx.Create(author).Error; nil != err {
		ctx.Rollback()
		return 0, err
	}
	ctx.Commit()
	return 1, nil
}

func (authorMapper * AuthorMapper) ListAuthors(page, size int)(Page, []Author){
	if page <= 0{
		page = 1
	}
	if size <= 0{
		size = 10
	}
	total := 0
	var authors []Author
	db.Table("author").Count(&total).Limit(size).Offset((page - 1) * size).Find(&authors)
	var pages = 0
	if total != 0{
		pages = total / size + 1
	}
	return Page{Page:page, Size:len(authors), Total:total, Pages:pages}, authors
}
