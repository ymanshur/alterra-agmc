package controller

import (
	cons "day-6/go-restful/constant"
	db "day-6/go-restful/database"
	lib "day-6/go-restful/lib"
	m "day-6/go-restful/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	// Define .env path for test
	mockUserEnvFilePath = "../.env.test"
	mockUserDB          = db.New(db.Config{
		User:     cons.Getenv("DB_USER", mockUserEnvFilePath),
		Password: cons.Getenv("DB_PASS", mockUserEnvFilePath),
		Host:     cons.Getenv("DB_HOST", mockUserEnvFilePath),
		Port:     cons.Getenv("DB_PORT", mockUserEnvFilePath),
		Name:     cons.Getenv("DB_NAME", mockUserEnvFilePath),
	})
	mockUsers = []m.User{
		{
			Name:     "Fityah Salamah",
			Email:    "fityah.salamah@gmail.com",
			Password: "1234",
		},
	}
)

func TestCreateUserCreated(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(
		`{"name": "Fityah Salamah", "email": "fityah.salamah@gmail.com", "password": "1234"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	defer db.Drop(mockUserDB, m.User{})
	db.Load(mockUserDB, m.User{})
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.CreateUser(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestCreateUserUnprocessableEntity(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(
		`{"name": "", "email": "fityah.salamah@gmail.com", "password": "1234"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	defer db.Drop(mockUserDB, m.User{})
	db.Load(mockUserDB, m.User{})
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.CreateUser(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestGetUserSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	defer db.Drop(mockUserDB, m.User{})
	db.Load(mockUserDB, m.User{})
	mockUserDB.Create(&mockUsers)
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.GetUser(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserInvalidParam(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("a")
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.GetUser(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateUserSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/users/:id", strings.NewReader(
		`{"name": "Fityah Salamah", "email": "fityahsalamah@gmail.com", "password": "1234"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	defer db.Drop(mockUserDB, m.User{})
	db.Load(mockUserDB, m.User{})
	mockUserDB.Create(&mockUsers)
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.UpdateUser(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserBadRequest(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, "/users/:id", strings.NewReader(
		`{"name": "", "email": "fityahsalamah@gmail.com", "password": "1234}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	defer db.Drop(mockUserDB, m.User{})
	db.Load(mockUserDB, m.User{})
	mockUserDB.Create(&mockUsers)
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.UpdateUser(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	defer db.Drop(mockUserDB, m.User{})
	db.Load(mockUserDB, m.User{})
	mockUserDB.Create(&mockUsers)
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.DeleteUser(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserInvalidParam(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("a")
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.DeleteUser(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetAllUserSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	defer db.Drop(mockUserDB, m.User{})
	db.Load(mockUserDB, m.User{})
	mockUserDB.Create(&mockUsers)
	c := NewUserController(mockUserDB)

	// Assertions
	if assert.NoError(t, c.GetAllUser(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
