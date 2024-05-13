package controller

import (
	"github.com/gofiber/fiber/v2"

	S "github.com/amirrstm/go-clinique/app/repository"
	H "github.com/amirrstm/go-clinique/handler"
	U "github.com/amirrstm/go-clinique/utils"
)

func GetUsers(ctx *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(ctx)

	if txErr != nil {
		return H.BuildError(ctx, "Unable to get transaction", fiber.StatusInternalServerError, txErr)
	}

	users, serviceErr := S.GetUsers(dbTrx, ctx.UserContext())

	if serviceErr != nil {
		return H.BuildError(ctx, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	return H.Success(ctx, fiber.Map{
		"ok":    1,
		"users": users,
	})
}
