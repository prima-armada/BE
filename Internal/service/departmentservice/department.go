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
	newDepartment.CreatedAt = time.Now()
	datarepo, errrepo := sd.rd.AddDepartment(newDepartment)

	if errrepo != nil {
		return request.RequestDepartment{}, errrepo
	}
	return datarepo, nil
}
