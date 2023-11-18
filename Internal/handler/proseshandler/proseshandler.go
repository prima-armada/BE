package proseshandler

import (
	"net/http"
	"par/domain/contract/handlecontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/helper"
	middlewares "par/middleware"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type HandlerProses struct {
	sp servicecontract.ServiceProcess
}

func NewHandlesProcess(sp servicecontract.ServiceProcess) handlecontract.HandleDetailProses {
	return &HandlerProses{
		sp: sp,
	}
}

func (hp *HandlerProses) AddProcess(e echo.Context) error {
	reqproses := request.ReqDetailProsesAdmin{}

	role := middlewares.ExtractTokenTeamRole(e)
	user, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role == "admin" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}

	binderr := e.Bind(&reqproses)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}

	dataservice, errservice := hp.sp.AddProcess(user, reqproses)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	// respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusCreated, helper.GetResponse(dataservice, http.StatusCreated, false))

}

func (hp *HandlerProses) GetallDetail(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)

	if role != "admin" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin ", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hp.sp.GetallDetail()
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	// respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(dataservice, http.StatusOK, false))
}

func (hp *HandlerProses) Updatedetail(e echo.Context) error {
	reqproses := request.ReqDetailProsesManager{}
	IdProses := e.QueryParam("idproses")

	cnv, errcnv := strconv.Atoi(IdProses)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errcnv, http.StatusBadRequest, true))
	}
	role := middlewares.ExtractTokenTeamRole(e)
	user, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "manager" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	binderr := e.Bind(&reqproses)
	reqproses.IdManager = uint(user)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hp.sp.UpdateDetail(cnv, reqproses)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	// respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(dataservice, http.StatusOK, false))
}

// GetallDetailManager implements handlecontract.HandleDetailProses.
func (hp *HandlerProses) GetallDetailManager(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	user, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "manager" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hp.sp.GetAlldetailManager(user)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	// respon := query.ReqtoResponKandidat(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(dataservice, http.StatusOK, false))
}
