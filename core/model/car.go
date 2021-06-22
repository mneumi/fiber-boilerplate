package model

import "fiber-boilerplate/pkg/global"

type CarDetail struct {
	ID uint `json:"id"`
	// BrandId  string  `json:"brand_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

func (c *CarDetail) List(page int, pageSize int, pageOffset int) ([]CarDetail, int64) {
	cars := make([]CarDetail, 0)
	var total int64

	err := global.DB.Model(&CarDetail{}).Count(&total).
		Offset(pageOffset).Limit(pageSize).Find(&cars).Error

	if err != nil {
		panic(err)
	}

	return cars, total
}
