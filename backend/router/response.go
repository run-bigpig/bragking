package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/bragking/backend/consts"
	"github.com/run-bigpig/bragking/backend/types"
)

func success(ctx *fiber.Ctx, data interface{}) error {
	res := types.Response{
		Code: 0,
		Data: data,
		Msg:  consts.Success,
	}
	return ctx.JSON(res)
}

func fail(ctx *fiber.Ctx, code int, error error) error {
	res := types.Response{
		Code: code,
		Data: nil,
		Msg:  error.Error(),
	}
	return ctx.JSON(res)
}
