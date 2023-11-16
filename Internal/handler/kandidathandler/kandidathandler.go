package kandidathandler

import (
	"net/http"
	"par/domain/contract/handlecontract"
	"par/domain/contract/servicecontract"
	"par/domain/query"
	"par/domain/request"
	"par/helper"
	middlewares "par/middleware"

	echo "github.com/labstack/echo/v4"
)

type Handlerkandidat struct {
	sk servicecontract.ServiceKandidat
}

func NewHandlesKandidat(sk servicecontract.ServiceKandidat) handlecontract.HandleKandidat {
	return &Handlerkandidat{
		sk: sk,
	}
}

func (hk *Handlerkandidat) AddFormulirKandidat(e echo.Context) error {
	reqformulir := request.ReqFormulirKandidat{}

	role := middlewares.ExtractTokenTeamRole(e)
	useradmin, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}

	binderr := e.Bind(&reqformulir)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	dataservice, errservice := hk.sk.AddFormulirKandidat(reqformulir, uint(useradmin))

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusCreated, helper.GetResponse(respon, http.StatusCreated, false))

}

func (hk *Handlerkandidat) GetCodeKandidat(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	kode := e.QueryParam("code")

	if role != "admin" && role != "manager" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hk.sk.GetCodeKandidat(kode)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	respon := query.Listtoreqresponkandidat(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))

}
