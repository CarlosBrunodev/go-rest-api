package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
//	return errors.New("fake error")
}

func Getlogger(p string) *Logger{
	logger =NewLogger(p)
	return logger
}
