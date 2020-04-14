package model

import (
	"github.com/jinzhu/gorm"
)

type Articles struct {
	gorm.Model
	Title   string
	Content string
	Status  int
	Tag     int
}

// 声明表名称，默认为 struct 的复数表示
func (Articles) TableName() string {
	return "articles"
}

func GetArticles(offset int, limit int) interface{} {
	return 123
}
