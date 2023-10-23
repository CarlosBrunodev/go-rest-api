package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err  error
	// Inicializer db 
	db, err = InicializerSQLite() 

	if err != nil {
		return fmt.Errorf("erro inicializer sqlite : %v", err)
	}
	
	return nil
//	return errors.New("fake error")
}

func GetSQLite() *gorm.DB{
	return db
}

func Getlogger(p string) *Logger{
	logger =NewLogger(p)
	return logger
}
