package controller

import (
	"fmt"
	"time"

	"github.com/amirrstm/go-clinique/app/dto"
	"github.com/amirrstm/go-clinique/pkg/config"
	"github.com/amirrstm/go-clinique/pkg/validator"

	R "github.com/amirrstm/go-clinique/app/repository"
	H "github.com/amirrstm/go-clinique/handler"
	U "github.com/amirrstm/go-clinique/utils"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// GetNewAccessToken method for create a new access token.
//
//	@Description	Create a new access token.
//	@Summary		create a new access token
//	@Tags			Token
//	@Accept			json
//	@Produce		json
//	@Param			login			body		dto.Auth		true	"Request for token"
//	@Failure		400,404,401,500	{object}	ErrorResponse
//	@Success		200				{object}	TokenResponse
//	@Router			/v1/login [post]
func GetNewAccessToken(c *fiber.Ctx) error {
	dbTrx, txErr := U.StartNewPGTrx(c)

	if txErr != nil {
		return H.ErrorHandler(c, fiber.StatusBadRequest, txErr)

	}

	login := &dto.Auth{}

	validate := validator.NewValidator()
	if err := validate.Struct(login); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input found",
			"errors":  validator.ValidatorErrors(err),
		})
	}

	if err := c.BodyParser(login); err != nil {
		return H.ErrorHandler(c, fiber.StatusBadRequest, err)

	}

	user, serviceErr := R.GetByUsername(dbTrx, c.UserContext(), login.Username)
	if serviceErr != nil {
		return H.BuildError(c, serviceErr.Message, serviceErr.Code, serviceErr.Error)
	}

	isValid := IsValidPassword([]byte(user.Password), []byte(login.Password))
	if !isValid {
		return H.BuildError(c, "Password is wrong!", fiber.StatusUnauthorized, nil)

	}

	if !user.IsActive.Bool {
		return H.BuildError(c, "user not active anymore", fiber.StatusUnauthorized, nil)

	}

	// Generate a new Access token.
	token, tokenErr := GenerateNewAccessToken(user.ID, user.IsAdmin.Bool)
	if tokenErr != nil {
		// Return status 500 and token generation error.
		return H.ErrorHandler(c, fiber.StatusInternalServerError, tokenErr)

	}

	return H.Success(c, fiber.Map{
		"message":      fmt.Sprintf("Token will be expired within %d minutes", config.AppCfg().JWTSecretExpireMinutesCount),
		"access_token": token,
	})

}

func GenerateNewAccessToken(userID int, isAdmin bool) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = isAdmin
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AppCfg().JWTSecretExpireMinutesCount)).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppCfg().JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GeneratePasswordHash(password []byte) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func IsValidPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
