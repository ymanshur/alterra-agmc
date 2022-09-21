package auth

import (
	"day-6/go-restful/internal/factory"
	"day-6/go-restful/internal/model"
	"day-6/go-restful/internal/repository"
	"day-6/go-restful/pkg/constant"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type controller struct {
	repo *repository.User
}

func (c *controller) SignUp(ctx echo.Context) error {
	// Bind
	user := new(model.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	// Validate
	if err := ctx.Validate(user); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}

	// Create user
	createdUser, err := c.repo.Save(user)
	if err != nil {
		errorMessage := err.Error()
		if strings.Contains(errorMessage, "Duplicate entry") {
			return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
				"message": errorMessage,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": errorMessage,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success create new user",
		"data":    createdUser,
	})
}

func (c *controller) SignIn(ctx echo.Context) error {
	// Bind
	user := new(model.User)
	ctx.Bind(&user)

	// Validate
	if user.Email == "" || user.Password == "" {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": constant.ErrReqEmailPassword,
		})
	}

	// Find user
	loggedInUser, err := c.repo.FindByEmailPassword(user)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
				"message": constant.ErrCredentialNotMatch,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success logged in",
		"data":    loggedInUser,
	})
}

func NewController(f *factory.Factory) *controller {
	return &controller{
		repo: f.UserRepository,
	}
}
