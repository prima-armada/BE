package departmenthandler

import (
	"net/http"
	"par/domain/contract/handlecontract"
	"par/domain/contract/servicecontract"
	"par/domain/query"
	"par/domain/request"
	"par/helper"
	middlewares "par/middleware"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type HandlerDepartment struct {
	sd servicecontract.ServiceDepartment
}

func NewHandlesDepartment(sd servicecontract.ServiceDepartment) handlecontract.HandleDepartment {
	return &HandlerDepartment{
		sd: sd,
	}
}

// AddDepartment implements handlecontract.HandleDepartment.
func (hd *HandlerDepartment) AddDepartment(e echo.Context) error {

	role := middlewares.ExtractTokenTeamRole(e)

	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	Reqdepartment := request.RequestDepartment{}

	binderr := e.Bind(&Reqdepartment)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hd.sd.Department(Reqdepartment)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	reqtorespon := query.ReqDepartmentToRespondepart(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(reqtorespon, http.StatusCreated, false))
}

// AllDepartment implements handlecontract.HandleDepartment.
func (hd *HandlerDepartment) AllDepartment(e echo.Context) error {

	role := middlewares.ExtractTokenTeamRole(e)

	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hd.sd.AllDepartment()

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ListReqDepartmentToRespondepart(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respondata, http.StatusOK, false))
}

func (hd *HandlerDepartment) UpdatedDepartment(e echo.Context) error {

	role := middlewares.ExtractTokenTeamRole(e)
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	id := e.Param("id")
	updatedepartment := request.RequestDepartment{}

	binderr := e.Bind(&updatedepartment)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	cnv, errcnv := strconv.Atoi(id)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errcnv.Error(), http.StatusBadRequest, true))
	}

	dataservice, errservice := hd.sd.UpdatedDepartment(cnv, updatedepartment)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	reqtorespon := query.ReqDepartUpdatementToRespondepart(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(reqtorespon, http.StatusOK, false))

}
