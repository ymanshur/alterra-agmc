package controller

import (
	"day-6/go-restful/constant"
	"day-6/go-restful/model"
	"day-6/go-restful/repository"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type userController struct {
	repo *repository.UserRepository
}

func (c *userController) CreateUser(ctx echo.Context) error {
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
	createdUser, err := c.repo.Create(user)
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

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "success create new user",
		"data":    createdUser,
	})
}

func (c *userController) GetUser(ctx echo.Context) error {
	// Validate parameter
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Get user
	user, err := c.repo.Get(uint(userId))
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

func (c *userController) UpdateUser(ctx echo.Context) error {
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
	updatedUser, err := c.repo.Update(uint(userId), user)
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

func (c *userController) DeleteUser(ctx echo.Context) error {
	// Validate parameter
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Delete user
	if err := c.repo.Delete(uint(userId)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success delete a user",
	})
}

func (c *userController) GetAllUser(ctx echo.Context) error {
	users, err := c.repo.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get all users",
		"data":    users,
	})
}

func (c *userController) LoginUser(ctx echo.Context) error {
	user := new(model.User)
	ctx.Bind(&user)

	// Throws bad request error
	// if user.Email == "" || user.Password == "" {
	// 	return echo.ErrBadRequest
	// }

	loggedInUser, err := c.repo.Login(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success logged in",
		"data":    loggedInUser,
	})
}

func NewUserController(db *gorm.DB) userController {
	return userController{
		repo: &repository.UserRepository{
			DB: db,
		},
	}
}
