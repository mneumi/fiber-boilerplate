package v1

import (
	"fiber-boilerplate/core/model"
	"fiber-boilerplate/pkg/response"
	"fiber-boilerplate/pkg/utils/pagintation"

	"github.com/gofiber/fiber/v2"
)

type Car struct{}

func (a Car) List(ctx *fiber.Ctx) error {
	page := pagintation.GetPage(ctx)
	pageSize := pagintation.GetPageSize(ctx)
	pageOffset := pagintation.GetPageOffset(page, pageSize)

	carModel := model.CarDetail{}

	cars, total := carModel.List(page, pageSize, pageOffset)

	response.New(ctx).ToResponseList(cars, int(total))

	return nil
}

func NewCar() Car {
	return Car{}
}
