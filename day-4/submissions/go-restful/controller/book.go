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

type BookHandler struct {
	books map[int]*model.Book
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	book := &model.Book{
		ID: seq,
	}
	if err := c.Bind(book); err != nil {
		return err
	}

	if book.Title == "" || book.Author == "" {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": constant.ErrInvalidInput.Error(),
		})
	}

	h.books[book.ID] = book
	seq++

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success create new book",
		"data":    book,
	})
}

func (h *BookHandler) GetBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	if _, exist := h.books[bookId]; !exist {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get a book",
		"data":    h.books[bookId],
	})
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	if _, exist := h.books[bookId]; !exist {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	if book.Title == "" || book.Author == "" {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": constant.ErrInvalidInput.Error(),
		})
	}

	h.books[bookId].Title = book.Title
	h.books[bookId].Author = book.Author

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update a book",
		"data":    h.books[bookId],
	})
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": constant.ErrInvalidUrlParam.Error(),
		})
	}

	if _, exist := h.books[bookId]; !exist {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": constant.ErrRecordNotFound.Error(),
		})
	}

	delete(h.books, bookId)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete a book",
	})
}

func (h *BookHandler) GetAllBook(c echo.Context) error {
	data := []model.Book{}
	for _, book := range h.books {
		data = append(data, *book)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all books",
		"data":    data,
	})
}
