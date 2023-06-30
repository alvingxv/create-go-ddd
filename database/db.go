package database

import (
	"fmt"
	"os"

	"github.com/alvingxv/create-go-ddd/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
	dialect  = "postgres"
)

// var (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "root"
// 	password = "root"
// 	dbname   = "todos-hacktiv"
// )

var db *gorm.DB

func HandleDatabaseConnection() {
	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})

	if err != nil {
		panic("failed connect to database")
	}

	db.AutoMigrate(entity.Todo{})

}

func GetDatabaseInstance() *gorm.DB {
	return db
}
