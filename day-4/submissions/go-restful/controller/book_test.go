package controller

import (
	"day-4/go-restful/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	envFilePath = "../.env.test"
	mockBookDB  = map[int]*model.Book{
		1: {ID: 1, Title: "Ulysses", Author: "James Joyce"},
	}
)

func TestCreateBookCreated(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(
		`{"title": "Ulysses", "author": "James Joyce"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)
	c := NewBookController(map[int]*model.Book{})

	// Assertions
	if assert.NoError(t, c.CreateBook(ctx)) {
		assert.Equal(t, http.StatusCreated, res.Code)
	}
}

func TestCreateBookUnprocessableEntity(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(
		`{"title": "", "author": "James Joyce"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	c := NewBookController(map[int]*model.Book{})

	// Assertions
	if assert.NoError(t, c.CreateBook(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestGetBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := NewBookController(mockBookDB)

	// Assertions
	if assert.NoError(t, c.GetBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookBadRequest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("a")
	c := NewBookController(mockBookDB)

	// Assertions
	if assert.NoError(t, c.GetBook(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/:id", strings.NewReader(
		`{"title": "Ulysses", "author": "James Joyce M."}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := NewBookController(mockBookDB)

	// Assertions
	if assert.NoError(t, c.UpdateBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookBadRequest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/:id", strings.NewReader(
		`{"title": "Ulysses", "author": ""}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := NewBookController(mockBookDB)

	// Assertions
	if assert.NoError(t, c.UpdateBook(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	c := NewBookController(mockBookDB)

	// Assertions
	if assert.NoError(t, c.DeleteBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookRecordNotFound(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("2")
	c := NewBookController(mockBookDB)

	// Assertions
	if assert.NoError(t, c.DeleteBook(ctx)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestGetAllBookSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	c := NewBookController(mockBookDB)

	// Assertions
	if assert.NoError(t, c.GetAllBook(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
