package daos

import (
	"CollegeAdministration/models"
)

func (ac *AdministrationCloud) AccountMigrationsCreate(req []*models.Account) error {

	err := ac.dbConn.Model(models.Account{}).Create(req).Error
	if err != nil {
		return err
	}
	return nil
}

func (ac *AdministrationCloud) AccountMigrationsUpdate(req []*models.Account) error {

	err := ac.dbConn.Model(models.Account{}).Save(req).Error
	if err != nil {
		return err
	}
	return nil
}
