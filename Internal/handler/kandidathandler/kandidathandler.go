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
	nama := e.QueryParam("manager")
	role := middlewares.ExtractTokenTeamRole(e)
	useradmin, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses manager", http.StatusUnauthorized, true))
	}

	binderr := e.Bind(&reqformulir)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	dataservice, errservice := hk.sk.AddFormulirKandidat(reqformulir, nama, uint(useradmin))

	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusBadRequest, true))
	}

	respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusCreated, helper.GetResponse(respon, http.StatusCreated, false))

}
