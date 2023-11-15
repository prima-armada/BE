package interviewhandler

import (
	"net/http"
	"par/domain/contract/handlecontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/helper"
	middlewares "par/middleware"

	echo "github.com/labstack/echo/v4"
)

type Handlerinterview struct {
	si servicecontract.Serviceinterview
}

func NewHandlesInterview(si servicecontract.Serviceinterview) handlecontract.HandleInterview {
	return &Handlerinterview{
		si: si,
	}
}

func (hi *Handlerinterview) AddFormulirInterview(e echo.Context) error {
	reqinterview := request.ReqInterviewKandidat{}

	role := middlewares.ExtractTokenTeamRole(e)
	user, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	reqinterview.UserId = uint(user)

	binderr := e.Bind(&reqinterview)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	dataservice, errservice := hi.si.AddInterview(reqinterview)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	// respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusCreated, helper.GetResponse(dataservice, http.StatusCreated, false))

}

// GetallInterview implements handlecontract.HandleInterview.
func (hi *Handlerinterview) GetallInterview(e echo.Context) error {
	kode := e.QueryParam("code")
	nama := e.QueryParam("nama")
	role := middlewares.ExtractTokenTeamRole(e)
	user, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hi.si.GetallInterview(user, kode, nama)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	// respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusCreated, helper.GetResponse(dataservice, http.StatusOK, false))
}
