package router

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App) {
	app.Get("/", IndexHandler)
	api := app.Group("/api/")
	api.Post("findcorrectanswercountlist", FindCorrectAnswerCountList)
	api.Post("findcontestantmyanswercountlist", FindContestantMyAnswerCountList)
	api.Post("findmyanswercountlist", FindMyAnswerCountList)
	api.Post("findchallengeranswercountlist", FindChallengerCorrectAnswerCountList)
	api.Post("findmaxuser", FindMaxUser)
	api.Post("findmaxbetuser", FindMaxBetUser)
	api.Post("findcloverbet", FindCloverBet)
	api.Post("finddatelist", FindDateList)
}
