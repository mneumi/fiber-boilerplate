package model

type Tag struct {
	Model
	Name  string `json:"name" gorm:"comment:'tag name'"`
	State uint8  `json:"state" gorm:"comment:'tag status 0=disable 1=enable'"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
