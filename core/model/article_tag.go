package model

type ArticleTag struct {
	Model
	TagID     uint `json:"tag_id" gorm:"comment:'tag ID'"`
	ArticleID uint `json:"article_id" gorm:"comment:'article ID'"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
