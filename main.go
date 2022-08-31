package main

import (
	"clean-golang/app/infrastructure/logger"
	"os"
	"runtime/debug"
)

func main() {
	defer func() {
		if initErr := recover(); initErr != nil {
			logger.Error(string(debug.Stack()))
		}
		os.Exit(1)
	}()
}
