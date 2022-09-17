package controller

import (
	"day-4/go-restful/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[int]*model.Book{
		1: {ID: 1, Title: "Ulysses", Author: "James Joyce"},
	}
	path = "/books"
)

func TestCreateBookCreated(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(
		`{"title": "Ulysses", "author": "James Joyce"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)
	c := &BookController{map[int]*model.Book{}}

	// Assertions
	if assert.NoError(t, c.CreateBook(ctx)) {
		assert.Equal(t, http.StatusCreated, res.Code)
	}
}

func TestCreateBookUnprocessableEntity(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(
		`{"title": "", "author": "James Joyce"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	c := &BookController{map[int]*model.Book{}}

	// Assertions
	if assert.NoError(t, c.CreateBook(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestGetBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/:id", path), nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := &BookController{mockDB}

	// Assertions
	if assert.NoError(t, c.GetBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookBadRequest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/:id", path), nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("a")
	c := &BookController{mockDB}

	// Assertions
	if assert.NoError(t, c.GetBook(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/:id", path), strings.NewReader(
		`{"title": "Ulysses", "author": "James Joyce M."}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := &BookController{mockDB}

	// Assertions
	if assert.NoError(t, c.UpdateBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookBadRequest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/:id", path), strings.NewReader(
		`{"title": "Ulysses", "author": ""}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := &BookController{mockDB}

	// Assertions
	if assert.NoError(t, c.UpdateBook(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/:id", path), nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := &BookController{mockDB}

	// Assertions
	if assert.NoError(t, c.DeleteBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookRecordNotFound(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/:id", path), nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")
	c := &BookController{mockDB}

	// Assertions
	if assert.NoError(t, c.DeleteBook(ctx)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestGetAllBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	c := &BookController{mockDB}

	// Assertions
	if assert.NoError(t, c.GetAllBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
