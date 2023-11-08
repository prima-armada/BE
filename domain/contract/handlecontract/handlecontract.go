package handlecontract

import "github.com/labstack/echo/v4"

type HandleUser interface {
	Register(e echo.Context) error
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
type HandleSubmission interface {
	AddSubmissionManager(e echo.Context) error
	GetAllSubmissionManager(e echo.Context) error
	GetAllSubmissionDireksi(e echo.Context) error
	GetAllSubmissionAdmin(e echo.Context) error
	UpdateSubmissionAdmin(e echo.Context) error
}
