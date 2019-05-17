package mapper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"zen/model"
)

type (
	BookMapper struct{}
	AuthorMapper struct{}
	BookAndAuthor struct{}
	Book = model.Book
	Page = model.Page
	Author = model.Author
)

const (
	dbName = "mysql"
	dbHost = "root:123456@tcp(127.0.0.1:3306)/goods?charset=utf8&parseTime=True&loc=Local"
)

var (
	db  *gorm.DB
	log = logrus.New()
)

// if you want a custom logger, init log here
func Setup() error {
	var err error
	if db, err = gorm.Open(dbName, dbHost); nil != err {
		log.Error("connect data base error: ", err)
		return err
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(5)
	return nil
}

func TearDown() error {
	if nil != db {
		return db.Close()
	}
	return nil
}
