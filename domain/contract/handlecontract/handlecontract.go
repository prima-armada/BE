package handlecontract

import "github.com/labstack/echo/v4"

type HandleUser interface {
	Register(e echo.Context) error
	NamaManager(e echo.Context) error
}

type HandleLogin interface {
	Login(e echo.Context) error
}

type HandleDepartment interface {
	AddDepartment(e echo.Context) error
	AllDepartment(e echo.Context) error
	UpdatedDepartment(e echo.Context) error
	DeletedDepartment(e echo.Context) error
}

type HandleSoal interface {
	Addsoal(e echo.Context) error
	AllSoal(e echo.Context) error
	KategoriSoal(e echo.Context) error
	UpdatedSoal(e echo.Context) error
	Deletedsoal(e echo.Context) error
}
type HandleSubmission interface {
	AddSubmission(e echo.Context) error
	GetAllSubmissionUser(e echo.Context) error
	GetAllSubmissionAdmin(e echo.Context) error
	UpdateSubmissionAdmin(e echo.Context) error
	GetNamaManager(e echo.Context) error
	UpdateSubmissionPresident(e echo.Context) error
	UpdateSubmissionDireksi(e echo.Context) error
	GetCode(e echo.Context) error
}
type HandleKandidat interface {
	AddFormulirKandidat(e echo.Context) error
	GetCodeKandidat(e echo.Context) error
}
type HandleInterview interface {
	AddFormulirInterview(e echo.Context) error
	AddFormulirInterviewFpt(e echo.Context) error
	GetallInterview(e echo.Context) error
	// CekKategoriInterview(e echo.Context) error
}
type HandleDetailProses interface {
	AddProcess(e echo.Context) error
	GetallDetail(e echo.Context) error
	Updatedetail(e echo.Context) error
	GetallDetailManager(e echo.Context) error
	UpdatedDetailAdmin(e echo.Context) error
	UpdatedDetaildireksi(e echo.Context) error
}
type HandlePosisi interface {
	AddPosisi(e echo.Context) error
}
type HandleSoalFPT interface {
	Addsoal(e echo.Context) error
	AllSoal(e echo.Context) error
	KategoriSoal(e echo.Context) error
	UpdatedSoal(e echo.Context) error
	Deletedsoal(e echo.Context) error
}
