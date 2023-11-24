package posisihandler

import (
	"net/http"
	"par/domain/contract/handlecontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/helper"
	middlewares "par/middleware"

	echo "github.com/labstack/echo/v4"
)

type Handlerposisi struct {
	sp servicecontract.ServicePosisi
}

func NewHandlesPosisi(sp servicecontract.ServicePosisi) handlecontract.HandlePosisi {
	return &Handlerposisi{
		sp: sp,
	}
}

// AddPosisi implements handlecontract.HandlePosisi.
func (hp *Handlerposisi) AddPosisi(e echo.Context) error {
	Reqposisi := request.ReqPosisi{}
	role := middlewares.ExtractTokenTeamRole(e)
	user, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses manager", http.StatusUnauthorized, true))
	}
	binderr := e.Bind(&Reqposisi)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	dataservice, errservice := hp.sp.AddPosisi(user, Reqposisi)

	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusBadRequest, true))
	}
	// reqtorespon := query.ReqmanagerToRespon(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(dataservice, http.StatusCreated, false))

}
