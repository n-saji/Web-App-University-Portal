package service

import (
	"CollegeAdministration/daos"

	"gorm.io/gorm"
)

type Service struct {
	daos *daos.AdministrationCloud
}

func New(dbConn *gorm.DB) *Service {
	return &Service{
		daos: daos.New(dbConn),
	}
}
