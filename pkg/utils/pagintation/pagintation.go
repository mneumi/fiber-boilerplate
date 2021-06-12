package pagintation

import (
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/utils/convert"

	"github.com/gofiber/fiber/v2"
)

func GetPage(ctx *fiber.Ctx) int {
	page := convert.StrTo(ctx.Query("page")).MustInt()

	if page == 0 {
		return 1
	}

	return page
}

func GetPageSize(ctx *fiber.Ctx) int {
	pageSize := convert.StrTo(ctx.Query("page_size")).MustInt()

	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	offset := 0

	if page > 0 {
		offset = (page - 1) * pageSize
	}

	return offset
}
