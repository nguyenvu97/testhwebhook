package initD

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm-tutorial/Model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//const DB_USERNAME = "root"
//const DB_PASSWORD = "Vu123456"
//const DB_NAME = "test1"
//const DB_HOST = "localhost"
//const DB_PORT = "3306"

func ConnectDB() (*gorm.DB, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
	}

	DB_HOST := viper.GetString("database.host")
	DB_PORT := viper.GetInt("database.port")
	DB_USERNAME := viper.GetString("database.username")
	DB_PASSWORD := viper.GetString("database.password")
	DB_NAME := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return nil, err
	}

	err = db.AutoMigrate(&Model.User{}, &Model.Book{}, &Model.LibraryBook{})
	if err != nil {

		return nil, err
	}

	return db, nil
}
