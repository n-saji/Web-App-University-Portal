package daos

import (
	"CollegeAdministration/models"
	"log"

	"gorm.io/gorm"
)

type AdminstrationCloud struct {
	dbConn *gorm.DB
}

func New(conn *gorm.DB) *AdminstrationCloud {
	return &AdminstrationCloud{dbConn: conn}
}

func (AC *AdminstrationCloud) RunMigrations() {
	
	err := AC.dbConn.Where("is_valid = ?", false).Delete(&models.Token_generator{}).Error
	if err != nil {
		log.Println("Error run cron migrations", err)
	}
}
