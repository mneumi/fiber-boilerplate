package response

import (
	"fiber-boilerplate/pkg/errcode"
	"fiber-boilerplate/pkg/utils/pagintation"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Ctx *fiber.Ctx
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

func New(ctx *fiber.Ctx) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(detail interface{}) error {
	ret := fiber.Map{
		"code": 0,
		"msg":  "成功",
	}

	if detail != nil {
		ret["details"] = detail
	}

	return r.Ctx.Status(fiber.StatusOK).JSON(ret)
}

func (r *Response) ToErrorResponse(err *errcode.Error) error {
	response := fiber.Map{
		"code": err.GetCode(),
		"msg":  err.GetMsg(),
	}

	details := err.GetDetails()

	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.Status(err.StatusCode()).JSON(response)

	return nil
}

func (r *Response) ToResponseList(list interface{}, totalRows int) error {
	return r.Ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"list": list,
		"pager": Pager{
			Page:      pagintation.GetPage(r.Ctx),
			PageSize:  pagintation.GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}
