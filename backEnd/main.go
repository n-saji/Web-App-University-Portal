package main

import (
	"CollegeAdministration/config"
	"CollegeAdministration/daos"
	"CollegeAdministration/handlers"
	"CollegeAdministration/service"
	"database/sql"
	"embed"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/robfig/cron.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {

	config.Init()
	db, err := gorm.Open(postgres.Open(config.DB_URL), &gorm.Config{})
	if err != nil {
		log.Println("Found err while connecting to database", err)
	}

	toRunGooseMigration(config.DB_URL)

	DaosConnection := daos.New(db)
	ServiceConnection := service.New(DaosConnection)
	handler_connection := handlers.New(ServiceConnection)

	s := cron.New()
	go s.AddFunc("@every 10m", ServiceConnection.RunDailyMigrations)
	go ServiceConnection.AccountDetailsMigration()
	s.Start()

	r := handler_connection.GetRouter()
	main_err := r.Run(config.Port)
	if main_err != nil {
		log.Println("MAIN - ERROR ", main_err)
	}
	s.Stop()
}

func toRunGooseMigration(url string) {

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println("db conn failed")
		panic(err)
	}
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Println("Setting Goose Postgres Dialect Failed")
		panic(err)
	}
	if err := goose.Up(db, "migrations", goose.WithAllowMissing()); err != nil {
		log.Println("Goose Up Failed")
		panic(err)
	}
}
