package repocontract

import (
	"par/domain/request"
	"time"
)

type RepoUser interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
	NipUserExist(nip string) (data request.RequestUser, err error)
	IdUserExist(id int) (data request.RequestUser, err error)
	UsernameUserExist(username string) (data request.RequestUser, err error)
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

type RepoSubmission interface {
	AddSubmissionManager(newSubmission request.ReqSubmissionManager, res time.Time) (request.ReqSubmissionManager, error)
	GetAllSubmissionManager(id int) ([]request.ReqGetManager, error)
	GetAllSubmissionDireksi(deparment string) ([]request.ReqGetDireksi, error)
	GetAllSubmissionPresident(deparment string) ([]request.ReqGetPresident, error)
	GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error)
	UpdateSubmissionAdmin(idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error)
	UpdateSubmissionPresident(idsubmission int, update request.UpdateVicePresident) (request.UpdateVicePresident, error)
	UpdateSubmissionDireksi(idsubmission int, update request.UpdateDireksi) (request.UpdateDireksi, error)
}
