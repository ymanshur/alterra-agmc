package controller

import (
	"day-4/go-restful/constant"
	"day-4/go-restful/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	seq = 1
)

type BookController struct {
	DB map[int]*model.Book
}

func (c *BookController) CreateBook(ctx echo.Context) error {
	book := &model.Book{
		ID: seq,
	}
	if err := ctx.Bind(book); err != nil {
		return err
	}

	if book.Title == "" || book.Author == "" {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": constant.ErrInvalidInput.Error(),
		})
	}

	c.DB[book.ID] = book
	seq++

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "success create new book",
		"data":    book,
	})
}

func (c *BookController) GetBook(ctx echo.Context) error {
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	if _, exist := c.DB[bookId]; !exist {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get a book",
		"data":    c.DB[bookId],
	})
}

func (c *BookController) UpdateBook(ctx echo.Context) error {
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	if _, exist := c.DB[bookId]; !exist {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	book := new(model.Book)
	if err := ctx.Bind(book); err != nil {
		return err
	}

	if book.Title == "" || book.Author == "" {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": constant.ErrInvalidInput.Error(),
		})
	}

	c.DB[bookId].Title = book.Title
	c.DB[bookId].Author = book.Author

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success update a book",
		"data":    c.DB[bookId],
	})
}

func (c *BookController) DeleteBook(ctx echo.Context) error {
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	if _, exist := c.DB[bookId]; !exist {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	delete(c.DB, bookId)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success delete a book",
	})
}

func (c *BookController) GetAllBook(ctx echo.Context) error {
	data := []model.Book{}
	for _, book := range c.DB {
		data = append(data, *book)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get all books",
		"data":    data,
	})
}

func NewBookController(db map[int]*model.Book) BookController {
	return BookController{
		DB: db,
	}
}
