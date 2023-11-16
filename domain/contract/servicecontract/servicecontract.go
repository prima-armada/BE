package servicecontract

import (
	"par/domain/request"
	"time"
)

type ServiceCase interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
	GetAllManager(roles string) ([]request.RequestUser, error)
}

type ServiceLogin interface {
	Login(nip string, password string) (string, request.RequestUser, error)
}

type ServiceSubmission interface {
	AddSubmissionManager(newSubmission request.ReqSubmissionManager, idManager int, res time.Time) (request.ReqSubmissionManager, error)
	GetNamaManager(namamanager string) ([]request.ReqGetManager, error)
	GetAllSubmissionManager(id int) ([]request.ReqGetManager, error)
	GetAllSubmissionDireksi(deparment string) ([]request.ReqGetDireksi, error)
	GetAllSubmissionPresident(deparment string) ([]request.ReqGetPresident, error)
	GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error)
	UpdateSubmissionAdmin(iduser int, idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error)
	UpdateSubmissionPresident(iduser int, idsubmission int, update request.UpdateVicePresident) (request.UpdateVicePresident, error)
	UpdateSubmissionDireksi(iduser int, idsubmission int, update request.UpdateDireksi) (request.UpdateDireksi, error)
	CodeSubmission(kode string) ([]request.ReqGetManager, error)
}
type ServiceKandidat interface {
	AddFormulirKandidat(newkandidata request.ReqFormulirKandidat, AdminId uint) (request.ReqFormulirKandidat, error)
	GetCodeKandidat(kode string) ([]request.ReqFormulirKandidat, error)
}
type ServiceSoal interface {
	AddSoal(newksoal request.RequesSoal) (request.RequesSoal, error)
	KategoriSoal(kategori string) (data request.RequesSoal, err error)
	AllSoal() (data []request.RequesSoal, err error)
	Updatedsoal(id int, update request.RequesSoal) (data request.RequesSoal, err error)
	DeleteSoal(id int) error
}

type ServiceDepartment interface {
	Department(newDepartment request.RequestDepartment) (request.RequestDepartment, error)
	AllDepartment() ([]request.RequestDepartment, error)
	UpdatedDepartment(id int, update request.RequestDepartment) (data request.RequestDepartment, err error)
	DeletedDepartment(id int) error
}
type Serviceinterview interface {
	AddInterview(newinterview request.ReqInterviewKandidat) (request.ReqInterviewKandidat, error)
	GetallInterview(userid int, kode, nama string) (data []request.ReqInterviewKandidat, err error)
}
type ServiceProcess interface {
	AddProcess(id int, newProcess request.ReqDetailProsesAdmin) (request.ReqDetailProsesAdmin, error)
	GetallDetail() (data []request.ReqDetailProses, err error)
}
