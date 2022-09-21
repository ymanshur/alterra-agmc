package controller

import (
	"day-4/go-restful/constant"
	"day-4/go-restful/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type bookController struct {
	db  map[int]*model.Book
	seq int
}

func (c *bookController) CreateBook(ctx echo.Context) error {
	// Bind
	book := &model.Book{
		ID: c.seq,
	}
	if err := ctx.Bind(book); err != nil {
		return err
	}

	// Validate
	if book.Title == "" || book.Author == "" {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": constant.ErrInvalidInput.Error(),
		})
	}

	// Create book
	c.db[book.ID] = book
	c.seq++

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "success create new book",
		"data":    book,
	})
}

func (c *bookController) GetBook(ctx echo.Context) error {
	// Validate parameter
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Check if exist
	if _, exist := c.db[bookId]; !exist {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get a book",
		"data":    c.db[bookId],
	})
}

func (c *bookController) UpdateBook(ctx echo.Context) error {
	// Validate parameter
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Check if exist
	if _, exist := c.db[bookId]; !exist {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	// Bind
	book := new(model.Book)
	if err := ctx.Bind(book); err != nil {
		return err
	}

	// Validate
	if book.Title == "" || book.Author == "" {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": constant.ErrInvalidInput.Error(),
		})
	}

	// Update book
	c.db[bookId].Title = book.Title
	c.db[bookId].Author = book.Author

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success update a book",
		"data":    c.db[bookId],
	})
}

func (c *bookController) DeleteBook(ctx echo.Context) error {
	// Validate parameter
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	// Check if exist
	if _, exist := c.db[bookId]; !exist {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	// Delete book
	delete(c.db, bookId)

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success delete a book",
	})
}

func (c *bookController) GetAllBook(ctx echo.Context) error {
	data := []model.Book{}
	for _, book := range c.db {
		data = append(data, *book)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "success get all books",
		"data":    data,
	})
}

func NewBookController(db map[int]*model.Book) bookController {
	seq := len(db) + 1
	return bookController{
		db:  db,
		seq: seq,
	}
}
