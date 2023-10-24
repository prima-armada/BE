package repocontract

import "par/domain/request"

type RepoUser interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllManager() (data []request.RequestUser, err error)
	AllAdmin() (data []request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
	NipadminExist(nip string) (data request.RequestUser, err error)
	NipUserExist(nip string) (data request.RequestUser, err error)
	Nipmanagerexist(nip string) (data request.RequestUser, err error)
}
type RepoLogin interface {
	Login(nip string, password string) (string, request.RequestUser, error)
}
type RepoDepartment interface {
	AddDepartment(newDepartment request.RequestDepartment) (request.RequestDepartment, error)
	AllDepertment() (data []request.RequestDepartment, err error)
	NameDepartment(name string) (data request.RequestDepartment, err error)
	UpdatedDepartment(id int, update request.RequestDepartment) (data request.RequestDepartment, err error)
	DeletedDepartment(id int) (row int, err error)
}
