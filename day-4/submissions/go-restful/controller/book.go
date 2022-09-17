package controller

import (
	"go-restful/constant"
	"go-restful/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success create new book",
		"data":    book,
	})
}

func GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.ErrInvalidUrlParam.Error())
	}

	if _, exist := books[id]; !exist {
		return echo.NewHTTPError(http.StatusNotFound, gorm.ErrRecordNotFound.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get a book",
		"data":    books[id],
	})
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.ErrInvalidUrlParam.Error())
	}

	if _, exist := books[id]; !exist {
		return echo.NewHTTPError(http.StatusNotFound, gorm.ErrRecordNotFound.Error())
	}

	u := new(model.Book)
	if err := c.Bind(u); err != nil {
		return err
	}

	books[id].Title = u.Title
	books[id].Author = u.Author

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update a book",
		"data":    books[id],
	})
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.ErrInvalidUrlParam.Error())
	}

	if _, exist := books[id]; !exist {
		return echo.NewHTTPError(http.StatusNotFound, gorm.ErrRecordNotFound.Error())
	}

	delete(books, id)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete a book",
	})
}

func GetAllBook(c echo.Context) error {
	data := []model.Book{}
	for _, book := range books {
		data = append(data, *book)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all books",
		"data":    data,
	})
}
