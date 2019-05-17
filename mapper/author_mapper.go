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
