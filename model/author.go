package model

type Author struct {
	Id       int    `gorm:"id" json:"-"`
	AuthorId string `gorm:"author_id" json:"authorId"`
	Name     string `gorm:"name" json:"name"`
	Sex      int    `gorm:"sex" json:"sex"`
	Age      int    `gorm:"age" json:"age"`
}

func (Author) TableName() string {
	return "author"
}
