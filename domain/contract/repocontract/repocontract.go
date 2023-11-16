package repocontract

import (
	"par/domain/request"
	"time"
)

type RepoUser interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
	GetAllManager(roles string) ([]request.RequestUser, error)
	NipUserExist(nip string) (data request.RequestUser, err error)
	IdUserExist(id int) (data request.RequestUser, err error)
	UsernameUserExist(username string) (data request.RequestUser, err error)
	NameExist(name string) (data request.RequestUser, err error)
}
type RepoLogin interface {
	Login(nip string, password string) (string, request.RequestUser, error)
}

type RepoSubmission interface {
	AddSubmissionManager(newSubmission request.ReqSubmissionManager, res time.Time) (request.ReqSubmissionManager, error)
	GetAllSubmissionManager(id int) ([]request.ReqGetManager, error)
	GetNamaManager(namamanager string) ([]request.ReqGetManager, error)
	NamaManager(namamanager string) (request.ReqGetManager, error)
	CodeSubmission(kode string) (request.ReqGetManager, error)
	CodeSubmissions(kode string) ([]request.ReqGetManager, error)
	GetAllSubmissionDireksi(deparment string) ([]request.ReqGetDireksi, error)
	GetAllSubmissionPresident(deparment string) ([]request.ReqGetPresident, error)
	GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error)
	UpdateSubmissionAdmin(idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error)
	UpdateSubmissionPresident(idsubmission int, update request.UpdateVicePresident) (request.UpdateVicePresident, error)
	UpdateSubmissionDireksi(idsubmission int, update request.UpdateDireksi) (request.UpdateDireksi, error)
}
type RepoKandidat interface {
	AddFormulirKandidat(newkandidat request.ReqFormulirKandidat) (request.ReqFormulirKandidat, error)
	GetCodeKandidat(kode string) ([]request.ReqFormulirKandidat, error)
	GetCodedannamaKandidat(kode, nama string) (request.ReqFormulirKandidat, error)
}
type RepoSoal interface {
	AddSoal(newksoal request.RequesSoal) (request.RequesSoal, error)
	KategoriSoal(kategori string) (data request.RequesSoal, err error)
	AllSoal() (data []request.RequesSoal, err error)
	Updatedsoal(id int, update request.RequesSoal) (data request.RequesSoal, err error)
	DeletedSoal(id int) (row int, err error)
}
type RepoDepartment interface {
	AddDepartment(newDepartment request.RequestDepartment) (request.RequestDepartment, error)
	AllDepertment() (data []request.RequestDepartment, err error)
	NameDepartment(name string) (data request.RequestDepartment, err error)
	UpdatedDepartment(id int, update request.RequestDepartment) (data request.RequestDepartment, err error)
	DeletedDepartment(id int) (row int, err error)
	IdDepartment(id int) (data request.RequestDepartment, err error)
}
type RepoInterview interface {
	AddInterview(newinterview request.ReqInterviewKandidat, tanggal time.Time) (request.ReqInterviewKandidat, error)
	GetallInterview(userid int, kode, nama string) (data []request.ReqInterviewKandidat, err error)
	CekKategorInterview(kategori string) (request.ReqInterviewKandidat, error)
}
type RepoProcess interface {
	AddProcess(newProcess request.ReqDetailProsesAdmin) (request.ReqDetailProsesAdmin, error)
	GetallDetail() (data []request.ReqDetailProses, err error)
	Getdetailkandidat(kode, nama, kandidat string) (data request.ReqDetailProsesAdmin, err error)
}
