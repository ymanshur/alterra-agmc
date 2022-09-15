package controller

import (
	"go-restful/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*model.Book{}
	seq   = 1
)

func CreateBook(c echo.Context) error {
	book := &model.Book{
		ID: seq,
	}
	if err := c.Bind(book); err != nil {
		return err
	}
	books[book.ID] = book
	seq++
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "success create new book",
		"data":    book,
	})
}

func GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func UpdateBook(c echo.Context) error {
	u := new(model.Book)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Title = u.Title
	return c.JSON(http.StatusOK, books[id])
}

func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.NoContent(http.StatusNoContent)
}

func GetAllBook(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}
