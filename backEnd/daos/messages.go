package daos

import (
	"CollegeAdministration/models"
	"fmt"

	"github.com/google/uuid"
)

func (dao *Daos) InsertIntoMessages(req *models.Messages) error {

	err := dao.dbConn.Model(models.Messages{}).Create(req).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) UpdateMessageStatusForAccountId(account_id uuid.UUID) error {

	err := dao.dbConn.Model(models.Messages{}).Where("account_id = ?", account_id).Update("is_read", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) UpdateMessageStatusforMsgId(msg_id uuid.UUID) error {

	err := dao.dbConn.Model(models.Messages{}).Where("id = ?", msg_id).Update("is_read", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) GetActiveMessagesForAccountId(account_id uuid.UUID) ([]models.Messages, error) {

	var messages []models.Messages
	err := dao.dbConn.Model(models.Messages{}).Where("account_id = ? and is_read = false", account_id).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (dao *Daos) DeleteMessageByAccountId(account_id uuid.UUID) error {

	fmt.Println("Deleting messages for account_id: ", account_id)
	q := dao.dbConn.Model(models.Messages{}).Where("account_id = ?", account_id).Delete(&models.Messages{})
	fmt.Println(q.RowsAffected)
	if q.Error != nil {
		return q.Error
	}
	return nil
}
