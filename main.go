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
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Postgres_Host, config.Postgres_User, config.Postgres_Password, config.Postgres_Db_Name, config.Postgres_Port)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Println("Found err while connecting to database", err)
	}

	err1 := db.Migrator().AutoMigrate(
		&models.CourseInfo{},
		&models.StudentInfo{},
		&models.InstructorDetails{},
		&models.InstructorLogin{},
	)
	if err != nil {
		log.Println("error found while migrating", err1)
	}

	DaosConnection := daos.New(db)
	ServiceConnection := service.New(DaosConnection)
	handler_connection := handlers.New(ServiceConnection)

	r := handler_connection.GetRouter()
	err2 := r.Run(":5050")
	if err2 != nil {
		log.Println("MAIN - ERROR ", err2)
	}

}
