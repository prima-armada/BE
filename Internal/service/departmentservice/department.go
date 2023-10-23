package departmentservice

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/validasi"
	"time"

	"github.com/go-playground/validator"
)

type ServicesDepartment struct {
	rd       repocontract.RepoDepartment
	validate *validator.Validate
}

func NewServiceDepartments(rd repocontract.RepoDepartment) servicecontract.ServiceDepartment {
	return &ServicesDepartment{
		rd:       rd,
		validate: validator.New(),
	}
}

func (sd *ServicesDepartment) Department(newDepartment request.RequestDepartment) (request.RequestDepartment, error) {

	_, errexist := sd.rd.NameDepartment(newDepartment.NameDepartment)

	if errexist == nil {
		return request.RequestDepartment{}, errors.New("Nama Departments Sudah Dibuat")
	}

	validerr := sd.validate.Struct(newDepartment)
	if validerr != nil {

		return request.RequestDepartment{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)

	newDepartment.CreatedAt = now

	datarepo, errrepo := sd.rd.AddDepartment(newDepartment)

	if errrepo != nil {
		return request.RequestDepartment{}, errrepo
	}
	return datarepo, nil
}

func (sd *ServicesDepartment) AllDepartment() ([]request.RequestDepartment, error) {
	datarepo, errrepo := sd.rd.AllDepertment()

	if errrepo != nil {
		return []request.RequestDepartment{}, errrepo
	}
	return datarepo, nil
}

func (sd *ServicesDepartment) UpdatedDepartment(id int, update request.RequestDepartment) (data request.RequestDepartment, err error) {

	if id <= 0 {
		return data, errors.New("data tidak ada atau kurang dari 0")
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)
	update.UpdateAt = now
	datarepo, errrepo := sd.rd.UpdatedDepartment(id, update)

	if errrepo != nil {
		return request.RequestDepartment{}, errrepo
	}
	return datarepo, nil
}

func (sd *ServicesDepartment) DeletedDepartment(id int) error {
	_, errrepo := sd.rd.DeletedDepartment(id)

	if errrepo != nil {
		return errrepo
	}
	return nil
}
