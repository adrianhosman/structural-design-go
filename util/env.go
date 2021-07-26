package util

import (
	"os"

	"github.com/adrianhosman/structural-design-go/model"
)

func GetEnv() string {
	env := os.Getenv(model.EnvKey)
	if env == "" {
		env = model.LocalEnv
	}
	return env
}
