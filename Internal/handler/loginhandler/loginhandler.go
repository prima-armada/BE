package loginhandler

import (
	"net/http"
	"par/domain/contract/handlecontract"
	"par/domain/contract/servicecontract"
	"par/domain/query"
	"par/domain/request"
	"par/helper"

	echo "github.com/labstack/echo/v4"
)

type HandlerLogin struct {
	sl servicecontract.ServiceLogin
}

func NewHandlLogin(sl servicecontract.ServiceLogin) handlecontract.HandleLogin {
	return &HandlerLogin{
		sl: sl,
	}
}

// Login implements handlecontract.HandleLogin.
func (hl *HandlerLogin) Login(e echo.Context) error {
	reques := request.RequestUser{}

	errbind := e.Bind(&reques)
	if errbind != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errbind.Error(), http.StatusBadRequest, true))
	}

	token, dataservice, errservice := hl.sl.Login(reques.Username, reques.Password)

	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respon := query.ReqtoRepon(dataservice, token)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))

}
