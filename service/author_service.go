package service

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func (authorService *AuthorService) SaveAuthor(author *Author) (int, error) {
	author.AuthorId = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	return authorMapper.SaveAuthor(author)
}

func (authorService * AuthorService) ListAuthors(page, size int)(Page, []Author){
	return authorMapper.ListAuthors(page, size)
}

