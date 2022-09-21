package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	filenames []string
}

func (e *Env) load() error {
	for _, filename := range e.filenames {
		if err := godotenv.Load(filename); err != nil {
			// Comment following line so that the testing process doesn't stop (breaking).
			// log.Fatalf("Error getting .env file, %v", err)
			return err
		}
	}
	return nil
}

func (e *Env) Get(key, fallback string) string {
	val, _ := os.LookupEnv(key)
	if val == "" {
		return fallback
	}
	return val
}

func NewEnv(filenames ...string) (*Env, error) {
	env := Env{
		filenames,
	}
	if err := env.load(); err != nil {
		return nil, err
	}
	return &env, nil
}
