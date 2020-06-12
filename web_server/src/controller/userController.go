package controller

import (
	appService "aprendendo/web_server/src/application/service"
	ctrldto "aprendendo/web_server/src/controller/dto"

	"github.com/gofiber/fiber"
)

func GetUserHandler(ctx *fiber.Ctx) {
	dto := new(ctrldto.UserPostDTO)

	if err := ctx.BodyParser(dto); err != nil {
		ctx.Status(400).Send(err)
		return
	}
	post, err := appService.UserPostService(dto)
	if err != nil {
		ctx.Status(500).Send(err)
	}

	if err := ctx.JSON(post); err != nil {
		ctx.Status(500).Send(err)
		return
	}
	ctx.JSON(post)
}
