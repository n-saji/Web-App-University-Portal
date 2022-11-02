package main

import (
	"CollegeAdministration/config"
	"CollegeAdministration/daos"
	"CollegeAdministration/handlers"
	"CollegeAdministration/models"
	"CollegeAdministration/service"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	config.Init()
	url := fmt.Sprintf("host=localhost user=%s dbname=%s port=5432 sslmode=disable", config.Postgres_User, config.Postgres_Db_Name)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("Found err while connecting to database", err)
	}

	err1 := db.Migrator().AutoMigrate(&models.CollegeAdminstration{}, &models.CoursesAvailable{})
	if err != nil {
		log.Println("error found while migrating", err1)
	}

	DaosConnection := daos.New(db)
	ServiceConnection := service.New(DaosConnection)
	handler_connection := handlers.New(ServiceConnection)

	r := handler_connection.GetRouter()
	r.Run(":5050")
	fmt.Println("No Errors Yeepee!!")

}
