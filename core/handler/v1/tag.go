package v1

import "github.com/gofiber/fiber/v2"

type Tag struct{}

func (t Tag) Get(c *fiber.Ctx) error    { return nil }
func (t Tag) List(c *fiber.Ctx) error   { return nil }
func (t Tag) Create(c *fiber.Ctx) error { return nil }
func (t Tag) Update(c *fiber.Ctx) error { return nil }
func (t Tag) Delete(c *fiber.Ctx) error { return nil }

func NewTag() Tag {
	return Tag{}
}
