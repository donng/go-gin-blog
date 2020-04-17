package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status" gorm:"default:1"`
	Tag     int    `json:"tag"`
}

// 声明表名称，默认为 struct 的复数表示
func (Article) TableName() string {
	return "articles"
}

func GetArticles(offset int, limit int) []Article {
	var articles []Article
	db.Offset(offset).Limit(limit).Find(&articles)

	return articles
}

func CreateArticle(title string, content string) (uint, error) {
	var article = Article{
		Title:   title,
		Content: content,
	}

	if err := db.Create(&article).Error; err != nil {
		return 0, err
	}

	return article.ID, nil
}
