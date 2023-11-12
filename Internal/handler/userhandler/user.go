package userhandler

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

type HandlerUser struct {
	um servicecontract.ServiceCase
}

func NewHandleUser(um servicecontract.ServiceCase) handlecontract.HandleUser {
	return &HandlerUser{
		um: um,
	}
}

// Register implements handlecontract.HandleUser.
func (Hu *HandlerUser) Register(e echo.Context) error {
	requestRegister := request.RequestUser{}

	binderr := e.Bind(&requestRegister)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	data, errservice := Hu.um.Register(requestRegister)
	// fmt.Print("ini data handler", data)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ReqtoResponUser(data)

	return e.JSON(http.StatusCreated, helper.GetResponse(respondata, http.StatusCreated, false))
}

// NamaManager implements handlecontract.HandleUser.
func (hu *HandlerUser) NamaManager(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	roles := e.Param("roles")
	dataservice, errservice := hu.um.GetAllManager(roles)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ListreqlUserToRes(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respondata, http.StatusOK, false))
}
