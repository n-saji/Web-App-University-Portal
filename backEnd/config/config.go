package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	Postgres_User     string
	Postgres_Db_Name  string
	Postgres_Password string
	Postgres_Port     int64
	Postgres_Host     string
	Port              string
	DB_URL            string
)

func Init() {
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_DB_NAME", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "password")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_HOST", "db")
	os.Setenv("PORT", ":5050")

	Postgres_User = os.Getenv("POSTGRES_USER")
	Postgres_Db_Name = os.Getenv("POSTGRES_DB_NAME")
	Postgres_Password = os.Getenv("POSTGRES_PASSWORD")
	Postgres_Port, _ = strconv.ParseInt(os.Getenv("POSTGRES_PORT"), 10, 64)
	Postgres_Host = os.Getenv("POSTGRES_HOST")
	Port = os.Getenv("PORT")

	os.Setenv("db_url", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", Postgres_Host, Postgres_User, Postgres_Password, Postgres_Db_Name, Postgres_Port))
	DB_URL = os.Getenv("db_url")
}
