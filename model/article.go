package model

type Article struct {
	Model
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	TagID   int    `json:"tag_id"`
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
