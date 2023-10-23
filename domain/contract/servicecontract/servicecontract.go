package servicecontract

import (
	"par/domain/request"
)

type ServiceCase interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
}

type ServiceLogin interface {
	Login(nip string, password string) (string, request.RequestUser, error)
}

type ServiceDepartment interface {
	Department(newDepartment request.RequestDepartment) (request.RequestDepartment, error)
	AllDepartment() ([]request.RequestDepartment, error)
	UpdatedDepartment(id int, update request.RequestDepartment) (data request.RequestDepartment, err error)
	DeletedDepartment(id int) error
}
