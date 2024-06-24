package database

import (
	"fmt"
	"redis_gorm_fiber/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectionMySQLDB creates a new MySQL database connection.
func ConnectionMySQLDB(config *config.Config) *gorm.DB {
	// Create the data source name (dsn) for the MySQL connection.
	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3333)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	// Open the MySQL connection and store the result in db.
	// gorm.Open returns a *gorm.DB and an error.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to MySQL DB success")
	// Return the GORM database connection.
	return db
}
