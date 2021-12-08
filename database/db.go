package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Connexion() *gorm.DB{
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=3308 sslmode=disable "
	ds, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Printf("Successfully connected! %T", ds)
	return ds
}