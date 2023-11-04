package repodepartment

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type RepoDepartment struct {
	db *gorm.DB
}

func NewRepoDepartments(db *gorm.DB) repocontract.RepoDepartment {
	return &RepoDepartment{
		db: db,
	}
}

// AddDepartment implements repocontract.RepoDepartment.
func (rd *RepoDepartment) AddDepartment(newDepartment request.RequestDepartment) (request.RequestDepartment, error) {
	reqdeparttomodeldepart := query.ReqDepartmentTomodelDepart(newDepartment)

	// alldepart, errall := rd.AllDepertment()

	// lendepart := len(alldepart)

	// if errall != nil {
	// 	return request.RequestDepartment{}, errall
	// }
	// if lendepart <= 0 || lendepart > 0 {
	// 	lendepart++
	// 	reqdeparttomodeldepart.Id = lendepart

	// }

	tx := rd.db.Create(&reqdeparttomodeldepart)

	if tx.Error != nil {
		return request.RequestDepartment{}, tx.Error
	}
	modeltoreq := query.ModeldepartmentToReqDepart(reqdeparttomodeldepart)

	return modeltoreq, nil
}

// AllDepertment implements repocontract.RepoDepartment.
func (rd *RepoDepartment) AllDepertment() (data []request.RequestDepartment, err error) {
	var activ []model.Department
	tx := rd.db.Raw("Select departments.id, departments.nama_department, departments.created_at, departments.updated_at from departments").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListDepartmentModelToReq(activ)
	fmt.Print("ini repo", dtmdlttoreq)
	return dtmdlttoreq, nil
}

// NameDepartment implements repocontract.RepoDepartment.
func (rd *RepoDepartment) NameDepartment(name string) (data request.RequestDepartment, err error) {
	var activ model.Department

	tx := rd.db.Raw("Select departments.nama_department, departments.created_at, departments.updated_at from departments WHERE departments.nama_department= ? ", name).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestDepartment{}, tx.Error
	}
	var activcore = query.ModeldepartmentToReqDepart(activ)

	return activcore, nil

}

func (rd *RepoDepartment) UpdatedDepartment(id int, update request.RequestDepartment) (data request.RequestDepartment, err error) {

	var depart model.Department

	tx1 := rd.db.Raw("Select departments.nama_department, departments.created_at, departments.updated_at from departments WHERE departments.id= ? ", id).First(&depart)

	if errors.Is(tx1.Error, gorm.ErrRecordNotFound) {

		return request.RequestDepartment{}, tx1.Error
	}

	reqdeparttomodeldepart := query.ReqDepartmentTomodelDepartudated(update)

	tx2 := rd.db.Model(&reqdeparttomodeldepart).Where("id = ?", id).Updates(&reqdeparttomodeldepart)

	if tx2.Error != nil {
		return request.RequestDepartment{}, tx2.Error
	}
	modeltoreq := query.ModelUpdatedepartmentToReqDepart(reqdeparttomodeldepart)

	return modeltoreq, nil
}

func (rd *RepoDepartment) DeletedDepartment(id int) (row int, err error) {
	Depart := model.Department{}

	tx := rd.db.Unscoped().Delete(&Depart, id)

	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete department by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}
