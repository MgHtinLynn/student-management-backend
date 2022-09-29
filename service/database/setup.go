package database

import (
	"fmt"
	"github.com/MgHtinLynn/final-year-project-mcc/service/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {

	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	viperHost := viper.Get("DB_HOST").(string)

	viperPort := viper.Get("DB_PORT").(string)

	viperUser := viper.Get("DB_USERNAME").(string)

	viperDB := viper.Get("DB_NAME").(string)

	viperPassword := viper.Get("DB_PASSWORD").(string)

	dbUrl := fmt.Sprintf("host=%v port=%v users=%v dbname=%v password=%v sslmode=disable", viperHost, viperPort, viperUser, viperDB, viperPassword)
	fmt.Println("dbUrl is\t\t", dbUrl)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	//
	if err != nil {
		log.Fatalln(err)
	}
	//
	//fmt.Println(db)

	_ = db.AutoMigrate(&models.User{})

	// Initialise value
	m := models.User{Name: "Htin Lynn", Email: "htinlin01@gmail.com", Phone: "09785360975", Address: "Home"}
	db.Create(&m)

	return db
}
