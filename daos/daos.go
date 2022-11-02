package daos

import (
	"gorm.io/gorm"
)

type AdminstrationCloud struct {
	dbConn *gorm.DB
}

func New(conn *gorm.DB) *AdminstrationCloud {
	return &AdminstrationCloud{dbConn: conn}
}
