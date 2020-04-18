package model

type Article struct {
	Model
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Status  int    `json:"status" gorm:"default:1"`
	TagID   int    `json:"tag_id"`
}

// 声明表名称，默认为 struct 的复数表示
func (Article) TableName() string {
	return "articles"
}

func GetArticles(offset int, limit int, tagID int) ([]Article, error) {
	var articles []Article

	tx := db.Order("id desc").Offset(offset).Limit(limit)
	if tagID != 0 {
		tx.Where("tag_id = ?", tagID)
	}
	if err := db.Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func CreateArticle(title string, desc string, content string, tagID int) (uint, error) {
	var article = Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		TagID:   tagID,
	}

	if err := db.Create(&article).Error; err != nil {
		return 0, err
	}

	return article.ID, nil
}

func FindArticle(id int) (Article, error) {
	var article Article
	if err := db.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}

func RemoveArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}
	return nil
}

func ModifyArticle(id int, title string, desc string, content string, tagID int) error {
	var article Article
	var err error
	if article, err = FindArticle(id); err != nil {
		return err
	}

	if title != "" {
		article.Title = title
	}
	if desc != "" {
		article.Desc = desc
	}
	if content != "" {
		article.Content = content
	}
	if tagID != 0 {
		article.TagID = tagID
	}

	if err = db.Save(&article).Error; err != nil {
		return err
	}

	return nil
}
