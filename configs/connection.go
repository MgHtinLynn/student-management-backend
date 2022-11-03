package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	//
	viperDBURL := viper.Get("DB_URL").(string)

	fmt.Println("viperDBURL", viperDBURL)

	//dbUrl := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", viperHost, viperPort, viperUser, viperDB, viperPassword, viperSSLMode)

	fmt.Println("dbUrl is\t\t", viperDBURL)

	db, err := gorm.Open(postgres.Open(viperDBURL), &gorm.Config{})
	//
	if err != nil {
		fmt.Println("err", err)
	}
	//
	fmt.Println(db)
	//if !db.Migrator().HasTable(&model.User{}) {
	//	_ = db.Migrator().CreateTable(&model.User{})
	//}
	// _ = db.AutoMigrate(&models.User{})

	//	Initialise value
	//m := models.User{Name: "Htin Lynn", Email: "htinlin01@gmail.com", Password: util.HashPassword("password"), Phone: "09785360975", Address: "Home", Active: true}
	//db.Create(&m)

	return db
}
