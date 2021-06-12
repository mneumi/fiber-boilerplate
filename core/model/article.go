package model

type Article struct {
	Model
	Title         string `json:"title" gorm:"comment:'article title'"`
	Desc          string `json:"desc" gorm:"comment:'article desc'"`
	Content       string `json:"content" gorm:"comment:'ariticle content'"`
	CoverImageUrl string `json:"cover_image_url" gorm:"comment:'cover image'"`
	State         uint8  `json:"state" gorm:"comment:'article status 0=disable 1=enable'"`
}

func (a Article) TableName() string {
	return "blog_article"
}
