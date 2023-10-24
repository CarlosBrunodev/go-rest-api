package handler

import (
	"github.com/CarlosBrunodev/go-rest-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler(){
	logger = config.Getlogger("handler")
	db = config.GetSQLite()
}


