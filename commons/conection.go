package commons

import (
	"log"

	"github.com/JuanDiegoCastellanos/apirest_go_mysql/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "juan_caste:@Manolo221212@/EasyFund")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
func Migrate() {
	db := GetConnection()

	defer db.Close()
	log.Println("Migrating...")

	db.AutoMigrate(&models.Person{})
}
