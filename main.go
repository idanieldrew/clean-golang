package main

import (
	"clean-golang/app/cmd"
	"clean-golang/app/infrastructure/database/fecades"
	"clean-golang/app/infrastructure/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime/debug"
)

func init() {
	defer func() {
		if initErr := recover(); initErr != nil {
			logger.Error(string(debug.Stack()))
			os.Exit(1)
		}
	}()

	// load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to db
	dbErr := fecades.NewDb(os.Getenv("DB_CONNECTION"))
	if dbErr != nil {
		logger.Error(dbErr.Error())
		return
	}
}

func main() {
	cmd.Run()
}
