package initializer

import (
	"github.com/jjonline/go-lib-backend/logger"
)

//go:noinline
func iniLogger() *logger.Logger {
	return logger.New("info", "stdout", "module")
}
