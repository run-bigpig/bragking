package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/bragking/backend/logic"
	"github.com/run-bigpig/bragking/backend/types"
)

func IndexHandler(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}

func FindDateList(ctx *fiber.Ctx) error {
	var req types.DateListReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	list, err := logic.NewStatistic(ctx).FindDateList(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, list)
}

func FindCloverBet(ctx *fiber.Ctx) error {
	total, err := logic.NewStatistic(ctx).FindCloverBet()
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, total)
}

func FindMaxBetUser(ctx *fiber.Ctx) error {
	var req types.MaxCountUserReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	list, err := logic.NewStatistic(ctx).FindMaxBetUser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, list)
}

func FindMaxUser(ctx *fiber.Ctx) error {
	var req types.MaxCountUserReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	list, err := logic.NewStatistic(ctx).FindMaxUser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, list)
}

func FindCorrectAnswerCountList(ctx *fiber.Ctx) error {
	var req types.AnswerStatisticReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	list, err := logic.NewStatistic(ctx).FindCorrectAnswerCountList(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, list)
}

func FindMyAnswerCountList(ctx *fiber.Ctx) error {
	var req types.AnswerStatisticReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	list, err := logic.NewStatistic(ctx).FindMyAnswerCountList(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, list)
}

func FindChallengerCorrectAnswerCountList(ctx *fiber.Ctx) error {
	var req types.AnswerStatisticReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	list, err := logic.NewStatistic(ctx).FindChallengerCorrectAnswerCountList(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, list)
}

func FindContestantMyAnswerCountList(ctx *fiber.Ctx) error {
	var req types.AnswerStatisticReq
	err := ctx.BodyParser(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	list, err := logic.NewStatistic(ctx).FindContestantMyAnswerCountList(&req)
	if err != nil {
		return fail(ctx, 500, err)
	}
	return success(ctx, list)
}
