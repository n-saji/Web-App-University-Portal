package daos

import (
	"gorm.io/gorm"
)

type AdministrationCloud struct {
	dbConn *gorm.DB
}

func New(conn *gorm.DB) *AdministrationCloud {
	return &AdministrationCloud{dbConn: conn}
}
