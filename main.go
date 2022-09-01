package main

import (
	"clean-golang/app/infrastructure/database/fecades"
	"clean-golang/app/infrastructure/logger"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime/debug"
)

func main() {
	defer func() {
		if initErr := recover(); initErr != nil {
			logger.Error(string(debug.Stack()))
			os.Exit(1)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql, err := fecades.NewMysqlFacade()
	if err != nil {
		return
	}

	if err != nil {
		logger.Error(err.Error())
	}
	_ = mysql
	log.Fatalln("ok")
}
