package user

import (
	"day-6/go-restful/internal/factory"
	"day-6/go-restful/internal/model"
	"day-6/go-restful/internal/repository"
	"day-6/go-restful/pkg/constant"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type controller struct {
	repo *repository.User
}

func (c *controller) Get(ctx echo.Context) error {
	// Validate parameter
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Get user
	user, err := c.repo.FindById(uint(userId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": err.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get a user",
		"data":    user,
	})
}

func (c *controller) Update(ctx echo.Context) error {
	// Validate parameter
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Bind
	user := new(model.User)
	ctx.Bind(&user)

	// Validate
	if err := ctx.Validate(user); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": err.Error(),
		})
	}

	// Update user
	updatedUser, err := c.repo.UpdateById(uint(userId), user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success update a user",
		"data":    updatedUser,
	})
}

func (c *controller) Delete(ctx echo.Context) error {
	// Validate parameter
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Delete user
	if err := c.repo.DeleteById(uint(userId)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success delete a user",
	})
}

func (c *controller) GetAll(ctx echo.Context) error {
	users, err := c.repo.FindAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get all users",
		"data":    users,
	})
}

func NewController(f *factory.Factory) *controller {
	return &controller{
		repo: f.UserRepository,
	}
}
