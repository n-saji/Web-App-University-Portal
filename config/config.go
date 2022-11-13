package config

import (
	"os"
	"strconv"
)

var (

	Postgres_User string
	Postgres_Db_Name string
	Postgres_Password string
	Postgres_Port	int64
	Postgres_Host string
)
func Init(){
	os.Setenv("POSTGRES_USER","postgres")
	os.Setenv("POSTGRES_DB_NAME","postgres")
	os.Setenv("POSTGRES_PASSWORD","password")
	os.Setenv("POSTGRES_PORT","5432")
	os.Setenv("POSTGRES_HOST","localhost")
	Postgres_User = os.Getenv("POSTGRES_USER")
	Postgres_Db_Name = os.Getenv("POSTGRES_DB_NAME")
	Postgres_Password = os.Getenv("POSTGRES_PASSWORD")
	Postgres_Port,_ = strconv.ParseInt(os.Getenv("POSTGRES_PORT"),10,64)
	Postgres_Host = os.Getenv("POSTGRES_HOST")
}