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
	AddSubmission(newSubmission request.ReqSubmission, res time.Time) (request.ReqSubmission, error)
	GetNamaManager(namamanager string) ([]request.ReqGetManager, error)
	NamaManager(namamanager string) (request.ReqGetManager, error)
	CodeSubmission(kode string) (request.ReqGetManager, error)
	CodeSubmissions(kode string) ([]request.ReqGetManager, error)
	GetAllSubmissionUser(deparment string) ([]request.ReqGetUsers, error)
	UpdateSubmissionAdmin(idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error)
	UpdateSubmissionPresident(idsubmission int, update request.UpdateVicePresident) (request.UpdateVicePresident, error)
	UpdateSubmissionDireksi(idsubmission int, update request.UpdateDireksi) (request.UpdateDireksi, error)
	GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error)
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
	AddInterviewfpt(newinterview request.ReqInterviewfpt, tanggal time.Time) (request.ReqInterviewfpt, error)
	GetallInterview(userid int, kode, nama string) (data []request.ReqInterviewKandidat, err error)
	CekallInterview(userid int, kode, nama string) (data []request.ReqInterviewKandidat, err error)
	GetallInterviewftp(nama string, kode string) (data []request.ReqInterviewfpt, err error)
	Getallnilaiftp(kode, nama string) (data []request.ReqInterviewfpt, err error)
	CekKategorInterview(kategori string) (request.ReqInterviewKandidat, error)
	// ReqInterviewKfpt
}
type RepoProcess interface {
	AddProcess(newProcess request.ReqDetailProsesAdmin) (request.ReqDetailProsesAdmin, error)
	GetallDetail() (data []request.ReqDetailProses, err error)
	GetdetailkandidatAdmin(kode, nama, kandidat string) (data request.ReqDetailProsesAdmin, err error)
	GetdetailkandidatManager(id int) (data request.ReqDetailProsesManager, err error)
	GetAlldetailManager(department string) (data []request.ReqDetailProsesManager, err error)
	UpdateDetail(id int, update request.ReqDetailProsesManager) (data request.ReqDetailProsesManager, err error)
	UpdateDetailAdmin(update request.ReqDetailProsesAdmin) (data request.ReqDetailProsesAdmin, err error)
	UpdateDetailDireksi(update request.ReqDetailProsesDireksi) (data request.ReqDetailProsesDireksi, err error)
}
type RepoPosisi interface {
	AddPosisi(newProcess request.ReqPosisi) (request.ReqPosisi, error)
	Getdetailposisi(userid int) (request.ReqPosisi, error)
}
type RepoSoalFpt interface {
	AddSoal(newsoal request.RequesSoalFpt) (request.RequesSoalFpt, error)
	KategoriSoal(kategori string) (data request.RequesSoalFpt, err error)
	AllSoal() (data []request.RequesSoalFpt, err error)
	Updatedsoal(id int, update request.RequesSoalFpt) (data request.RequesSoalFpt, err error)
	DeletedSoal(id int) (row int, err error)
}
