package db

import (
	"employeeManagement/app/entities"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbCon *gorm.DB

func GetDBConnection() *gorm.DB {
	return dbCon
}

func SetDBConnection(db *gorm.DB) {
	dbCon = db
}

// Create a new connection to the database
func InitDBConnection() error {
	dbname := "employeeManagementDB"
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	// host := "localhost"
	// port := "5432"
	// user := "postgres"
	// password := "admin123"

	dsn := fmt.Sprintf("host=%s, user=%s password=%s dbname=%s port=%s sslmode=disable", host,
		user, password, dbname, port)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Employee{})

	dbCon = db
	return nil
}
