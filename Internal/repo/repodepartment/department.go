package repodepartment

import (
	"errors"
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

	alldepart, errall := rd.AllDepertment()

	lendepart := len(alldepart)

	if errall != nil {
		return request.RequestDepartment{}, errall
	}
	if lendepart <= 0 || lendepart > 0 {
		reqdeparttomodeldepart.Id = lendepart
	}

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
	tx := rd.db.Raw("Select departments.nama_department, departments.created_at, departments.update_at from departments").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListDepartmentModelToReq(activ)
	return dtmdlttoreq, nil
}

// NameDepartment implements repocontract.RepoDepartment.
func (rd *RepoDepartment) NameDepartment(name string) (data request.RequestDepartment, err error) {
	var activ model.Department

	tx := rd.db.Raw("Select departments.nama_department, departments.created_at, departments.update_at from departments WHERE departments.nama_department= ? ", name).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestDepartment{}, tx.Error
	}
	var activcore = query.ModeldepartmentToReqDepart(activ)

	return activcore, nil

}
