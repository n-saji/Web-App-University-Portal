package daos

import (
	"CollegeAdministration/models"
	"log"

	"github.com/google/uuid"
)

// func (dao *Daos) AccountCreation(req []*models.Account) error {

// 	err := dao.dbConn.Model(models.Account{}).Create(req).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (dao *Daos) AccountMigrationsUpdate(req []*models.Account) error {

	err := dao.dbConn.Model(models.Account{}).Save(req).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) CreateAccount(acc *models.Account) error {

	err := dao.dbConn.Model(models.Account{}).Create(acc).Error
	if err != nil {
		log.Println("unable to persist")
		return err
	}
	return nil

}

func (dao *Daos) DeleteAccount(id uuid.UUID) error {

	err := dao.dbConn.Model(models.Account{}).Delete("id", id).Error
	if err != nil {
		log.Println("unable to delete account")
		return err
	}
	return nil
}

func (dao *Daos) GetAccountIDsByType(accountType string) ([]models.Account, error) {

	var accounts []models.Account
	err := dao.dbConn.Model(models.Account{}).Select("id").Where("type = ?", accountType).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (dao *Daos) GetAccountByID(id uuid.UUID) (*models.Account, error) {

	var acc models.Account
	err := dao.dbConn.Model(models.Account{}).Where("id = ?", id).Find(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

func (dao *Daos) GetAccountNameById(id uuid.UUID) (*models.Account, error) {

	var acc models.Account
	err := dao.dbConn.Model(models.Account{}).Select("name").Where("id = ?", id).Find(&acc).Error
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
