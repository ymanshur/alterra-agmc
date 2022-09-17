package controller

import (
	"day-4/go-restful/constant"
	"day-4/go-restful/model"
	"day-4/go-restful/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	repo *repository.UserRepository
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	user := new(model.User)
	ctx.Bind(&user)

	createdUser, err := c.repo.Create(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "success create new user",
		"data":    createdUser,
	})
}

func (c *UserController) GetUser(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.ErrInvalidUrlParam.Error())
	}

	user, err := c.repo.Get(uint(userId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get a user",
		"data":    user,
	})
}

func (c *UserController) UpdateUser(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.ErrInvalidUrlParam.Error())
	}

	// Return http.StatusNotFound if user does not exist
	if _, err := c.repo.Get(uint(userId)); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	user := new(model.User)
	ctx.Bind(&user)

	updatedUser, err := c.repo.Update(uint(userId), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success update a user",
		"data":    updatedUser,
	})
}

func (c *UserController) DeleteUser(ctx echo.Context) error {
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.ErrInvalidUrlParam.Error())
	}

	if err := c.repo.Delete(uint(userId)); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success delete a user",
	})
}

func (c *UserController) GetAllUser(ctx echo.Context) error {
	users, err := c.repo.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get all users",
		"data":    users,
	})
}

func (c *UserController) LoginUser(ctx echo.Context) error {
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

func NewUserController(g *gorm.DB) UserController {
	return UserController{
		repo: &repository.UserRepository{
			DB: g,
		},
	}
}
