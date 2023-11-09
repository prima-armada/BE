package servicecontract

import (
	"par/domain/request"
	"time"
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

type ServiceSubmission interface {
	AddSubmissionManager(newSubmission request.ReqSubmissionManager, idManager int, res time.Time) (request.ReqSubmissionManager, error)
	GetAllSubmissionManager(id int) ([]request.ReqGetManager, error)
	GetAllSubmissionDireksi(deparment string) ([]request.ReqGetDireksi, error)
	GetAllSubmissionPresident(deparment string) ([]request.ReqGetPresident, error)
	GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error)
	UpdateSubmissionAdmin(iduser int, idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error)
	UpdateSubmissionPresident(iduser int, idsubmission int, update request.UpdateVicePresident) (request.UpdateVicePresident, error)
	UpdateSubmissionDireksi(iduser int, idsubmission int, update request.UpdateDireksi) (request.UpdateDireksi, error)
}
