package config

import (
	"os"
	"github.com/CarlosBrunodev/go-rest-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InicializerSQLite() (*gorm.DB, error){

	logger := Getlogger("sqlite")
	dbPath := "./db/main.db" 
	// check if the databese file exists 
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database not found , creating ...")
		err = os.Mkdir("./db",os.ModePerm)
		if err != nil {
			return nil, err
		}
		file , err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create database connect
	db , err := gorm.Open(sqlite.Open(dbPath),&gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil , err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite migrate error: %v", err)
		return nil, err
	}

	return db, nil 
}