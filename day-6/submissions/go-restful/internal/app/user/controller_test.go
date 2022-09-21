package user

import (
	"day-6/go-restful/database"
	"day-6/go-restful/internal/factory"
	"day-6/go-restful/internal/model"
	"day-6/go-restful/internal/repository"
	"day-6/go-restful/pkg/util"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	basePath, _ = os.Getwd()
	env, _      = util.NewEnv(filepath.Join(basePath, "../../../.env.test")) //.env path for test
	db          = database.New(database.Config{
		User:     env.Get("DB_USER", ""),
		Password: env.Get("DB_PASS", ""),
		Host:     env.Get("DB_HOST", ""),
		Port:     env.Get("DB_PORT", ""),
		Name:     env.Get("DB_NAME", ""),
	})
	f        = factory.Factory{UserRepository: repository.NewUserRepository(db)}
	c        = NewController(&f)
	mockData = []model.User{
		{
			Name:     "Fityah Salamah",
			Email:    "fityah.salamah@gmail.com",
			Password: "1234",
		},
	}
	baseURL = "/api/users"
)

func TestGetSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, baseURL+"/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	database.Seed(db, model.User{})
	db.Create(&mockData)

	// Assertions
	if assert.NoError(t, c.Get(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetInvalidParam(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, baseURL+"/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("a")

	// Assertions
	if assert.NoError(t, c.Get(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, baseURL+"/:id", strings.NewReader(
		`{"name": "Fityah Salamah", "email": "fityahsalamah@gmail.com", "password": "1234"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	database.Seed(db, model.User{})
	db.Create(&mockData)

	// Assertions
	if assert.NoError(t, c.Update(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBadRequest(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPut, baseURL+"/:id", strings.NewReader(
		`{"name": "", "email": "fityahsalamah@gmail.com", "password": "1234}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	// Assertions
	if assert.NoError(t, c.Update(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestDeleteSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, baseURL+"/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	database.Seed(db, model.User{})
	db.Create(&mockData)

	// Assertions
	if assert.NoError(t, c.Delete(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteInvalidParam(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, baseURL+"/:id", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("a")

	// Assertions
	if assert.NoError(t, c.Delete(ctx)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetAllSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, baseURL, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	database.Seed(db, model.User{})
	db.Create(&mockData)

	// Assertions
	if assert.NoError(t, c.GetAll(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
