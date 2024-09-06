package config

import(
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"

)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "my data base") // create and add the my data base !!!
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}