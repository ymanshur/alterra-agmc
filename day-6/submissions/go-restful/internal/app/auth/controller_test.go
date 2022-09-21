package auth

import (
	"day-4/go-restful/lib"
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
	baseURL = "/api/auth"
)

func TestSignUpSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, baseURL+"/signup", strings.NewReader(
		`{"name": "Fityah Salamah", "email": "fityah.salamah@gmail.com", "password": "1234"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	database.Seed(db, model.User{})

	// Assertions
	if assert.NoError(t, c.SignUp(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignUpUnprocessableEntity(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, baseURL+"/signup", strings.NewReader(
		`{"name": "", "email": "fityah.salamah@gmail.com", "password": "1234"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	database.Seed(db, model.User{})

	// Assertions
	if assert.NoError(t, c.SignUp(ctx)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}

func TestSignInSuccess(t *testing.T) {
	// Setup
	e := echo.New()
	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, baseURL+"/signin", strings.NewReader(
		`{"email": "fityah.salamah@gmail.com", "password": "1234"}`,
	))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	database.Seed(db, model.User{})
	db.Create(&mockData)

	// Assertions
	if assert.NoError(t, c.SignIn(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
