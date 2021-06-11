package v1

import "github.com/gofiber/fiber/v2"

type Article struct{}

func (a Article) Get(c *fiber.Ctx) error    { return nil }
func (a Article) List(c *fiber.Ctx) error   { return nil }
func (a Article) Create(c *fiber.Ctx) error { return nil }
func (a Article) Update(c *fiber.Ctx) error { return nil }
func (a Article) Delete(c *fiber.Ctx) error { return nil }

func NewArticle() Article {
	return Article{}
}
