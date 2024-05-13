package repository

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	D "github.com/amirrstm/go-clinique/app/dto"
	M "github.com/amirrstm/go-clinique/app/models"
	T "github.com/amirrstm/go-clinique/types"
)

func GetUsers(dbTrx boil.ContextExecutor, ctx context.Context) ([]*M.User, *T.ServiceError) {
	products, err := M.Users().All(ctx, dbTrx)

	if err != nil {
		return nil, &T.ServiceError{
			Message: "Unable to get users",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	return products, nil
}

func CreateUser(dbTrx boil.ContextExecutor, ctx context.Context, body *D.CreateUser) (*M.User, *T.ServiceError) {
	user := M.User{
		Email:     body.Email,
		Password:  body.Password,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	if err := user.Insert(ctx, dbTrx, boil.Infer()); err != nil {
		return nil, &T.ServiceError{
			Error:   err,
			Message: "Unable to create user",
			Code:    fiber.StatusInternalServerError,
		}
	}

	return &user, nil
}

func GetByUsername(dbTrx boil.ContextExecutor, ctx context.Context, username string) (*M.User, *T.ServiceError) {
	user, err := M.Users(qm.Where("username = ?", username)).One(ctx, dbTrx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &T.ServiceError{
				Message: "User not found",
				Error:   err,
				Code:    fiber.StatusNotFound,
			}
		}

		return nil, &T.ServiceError{
			Message: "Unable to get user",
			Error:   err,
			Code:    fiber.StatusInternalServerError,
		}
	}

	return user, nil
}
