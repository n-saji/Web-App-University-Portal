package config

import "os"

var (

	Postgres_User string
	Postgres_Db_Name string
)
func Init(){
	os.Setenv("POSTGRES_USER","postgres")
	os.Setenv("POSTGRES_DB_NAME","collegeport")
	Postgres_User = os.Getenv("POSTGRES_USER")
	Postgres_Db_Name = os.Getenv("POSTGRES_DB_NAME")

}