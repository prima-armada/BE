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
	AddSubmissionManager(e echo.Context) error
	GetAllSubmissionManager(e echo.Context) error
	GetAllSubmissionDireksi(e echo.Context) error
	GetAllSubmissionPresident(e echo.Context) error
	GetAllSubmissionAdmin(e echo.Context) error
	UpdateSubmissionAdmin(e echo.Context) error
	GetNamaManager(e echo.Context) error
	UpdateSubmissionPresident(e echo.Context) error
	UpdateSubmissionDireksi(e echo.Context) error
}
type HandleKandidat interface {
	AddFormulirKandidat(e echo.Context) error
}
