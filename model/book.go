package model

import (
	"fmt"
	"time"
)

type (
	JsonTime time.Time
)

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(jsonTime).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (jsonTime *JsonTime) UnmarshalJSON(b []byte) error {
	if stamp, err := time.Parse("\"2006-01-02 15:04:05\"", string(b[:])); nil != err {
		return err
	} else {
		*jsonTime = JsonTime(stamp)
		return nil
	}
}

type Book struct {
	Id          int       `gorm:"id" json:"-"`
	BookId      string    `gorm:"book_id" json:"bookId"`
	Name        string    `gorm:"name" json:"name"`
	AuthorId    string    `gorm:"author_id" json:"authorId"`
	Price       int       `gorm:"price" json:"price"`
	Time        time.Time `gorm:"time" json:"-"`
	ReleaseTime JsonTime  `json:"time"`
}

func (Book) TableName() string {
	return "book"
}
