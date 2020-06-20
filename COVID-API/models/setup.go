package models

//importamos gorm que es un mini-orm de go
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "data.sqlite")

	if err != nil {
		panic("Fallo en conectar a la base de datos!")
	}

	database.AutoMigrate(&Death{})

	DB = database
}
