package daos

import (
	"CollegeAdministration/models"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (dao *Daos) InsertInstructorDetails(id *models.InstructorDetails) error {
	err := dao.dbConn.Table("instructor_details").Create(id).Error
	if err != nil {
		log.Println("Not able to insert instructor details", err)
		return fmt.Errorf("error while inserting instructor details %s", err.Error())
	}
	return nil
}

func (dao *Daos) GetAllInstructor() ([]*models.InstructorDetails, error) {
	var id []*models.InstructorDetails
	err := dao.dbConn.Order("instructor_name ASC").Find(&id).Error
	if err != nil {
		return nil, fmt.Errorf("not able to retrieve instructor details")
	}
	return id, nil
}

func (dao *Daos) GetAllInstructorOrderByCondition(order_clause string) ([]*models.InstructorDetails, error) {
	var id []*models.InstructorDetails
	q := dao.dbConn
	switch order_clause {
	case "instructor_code":

		q = q.Order("instructor_code ASC")

	case "instructor_name":

		q = q.Order("instructor_name ASC")

	case "department":

		q = q.Order("department ASC")

	// case "course_name":

	// 	q = q.Order("course_id ASC")

	// case "students_enrolled":
	// 	q = q.Raw("select *, (jsonb_array_length(info->'students_list')) as ct from instructor_details id  where id.info ->> 'students_list' is not null group by id.id union select *,0 as ct from instructor_details id where id.info ->> 'students_list' is  null order by ct desc ; ")

	default:
		return nil, fmt.Errorf("no order by clause given")
	}
	
	q = q.Find(&id)
	err := q.Error
	if err != nil {
		return nil, fmt.Errorf("not able to retrieve instructor details")
	}
	return id, nil
}

func (dao Daos) GetInstructor(id_exits *models.InstructorDetails) (*models.InstructorDetails, error) {
	var id models.InstructorDetails
	err := dao.dbConn.Where(&id_exits).Find(&id).Error
	if err != nil {
		return nil, fmt.Errorf("record not found")
	}
	return &id, nil
}

func (dao *Daos) DeleteInstructor(name string) error {

	err := dao.dbConn.Where("instructor_name = ?", name).Delete(models.InstructorDetails{}).Error

	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) GetInstructorWithName(name string) (*models.InstructorDetails, error) {

	var is models.InstructorDetails
	err := dao.dbConn.Model(models.InstructorDetails{}).Select("*").Where("instructor_name = ?", name).Find(&is).Error
	if err != nil {
		return nil, err
	}
	return &is, nil
}

func (dao *Daos) GetInstructorWithSpecifics(condition models.InstructorDetails) ([]*models.InstructorDetails, error) {

	var is []*models.InstructorDetails
	err := dao.dbConn.Model(models.InstructorDetails{}).Select("*").Where(condition).Find(&is).Error
	if err != nil {
		return nil, err
	}
	return is, nil
}

func (dao *Daos) UpdateInstructor(req_id *models.InstructorDetails, condition *models.InstructorDetails) error {

	q := dao.dbConn.Model(models.InstructorDetails{}).Where(condition).Updates(req_id)
	if q.Error != nil {
		return q.Error
	}
	return nil
}

// func (dao *Daos) UpdateInstructorInfo(req_id *models.InstructorDetails, condition *models.InstructorDetails) error {
// 	q := dao.dbConn.Model(models.InstructorDetails{}).Where(condition).Update("info", req_id.Info)
// 	if q.Error != nil {
// 		return q.Error
// 	}
// 	return nil
// }

func (dao *Daos) RetieveInstructorDetailsWithCondition(req models.InstructorDetails) ([]*models.InstructorDetails, error) {
	var list []*models.InstructorDetails
	err := dao.dbConn.Debug().Model(models.InstructorDetails{}).Select("*").Where(req).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (dao *Daos) DeleteInstructorWithConditions(id *models.InstructorDetails) error {

	err := dao.dbConn.Where(id).Delete(&models.InstructorDetails{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (dao *Daos) DisableToken(token uuid.UUID) error {

	err := dao.dbConn.Table("token_generators").Where("token = ?", token).Update("is_valid", "false").Error
	if err != nil {
		return err
	}
	return nil
}
