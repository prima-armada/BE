package submissionhandler

import (
	"net/http"
	"par/domain/contract/handlecontract"
	"par/domain/contract/servicecontract"
	"par/domain/query"
	"par/domain/request"
	"par/helper"
	middlewares "par/middleware"
	"strconv"
	"time"

	echo "github.com/labstack/echo/v4"
)

type HandlerSubmission struct {
	ss servicecontract.ServiceSubmission
}

func NewHandlesSubmission(ss servicecontract.ServiceSubmission) handlecontract.HandleSubmission {
	return &HandlerSubmission{
		ss: ss,
	}
}

func (hs *HandlerSubmission) AddSubmissionManager(e echo.Context) error {
	Reqmanager := request.ReqSubmissionManager{}
	role := middlewares.ExtractTokenTeamRole(e)
	usermanagar, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "manager" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses manager", http.StatusUnauthorized, true))
	}
	binderr := e.Bind(&Reqmanager)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	res, errConvtime := time.Parse("02/01/2006", Reqmanager.TanggalKebutuhan)
	if errConvtime != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errConvtime.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hs.ss.AddSubmissionManager(Reqmanager, usermanagar, res)

	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusBadRequest, true))
	}
	reqtorespon := query.ReqmanagerToRespon(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(reqtorespon, http.StatusCreated, false))

}

func (hs *HandlerSubmission) GetAllSubmissionManager(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	usermanagar, errtoken := middlewares.ExtractTokenIdUser(e)

	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken, http.StatusUnauthorized, true))
	}
	if role != "manager" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses manager", http.StatusUnauthorized, true))
	}

	dataservice, errservice := hs.ss.GetAllSubmissionManager(usermanagar)

	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusBadRequest, true))
	}
	respon := query.ListReqltoResmanager(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

func (hs *HandlerSubmission) GetAllSubmissionAdmin(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))

	}
	dataservice, errservice := hs.ss.GetAllSubmissionAdmin()
	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusBadRequest, true))
	}
	respon := query.ListReqltoResAdmin(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

func (hs *HandlerSubmission) GetAllSubmissionDireksi(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	departdireksi := middlewares.ExtractTokenTeamDepartment(e)

	if role != "direksi" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses direksi", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hs.ss.GetAllSubmissionDireksi(departdireksi)
	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusBadRequest, true))
	}
	respon := query.ListReqltoResDireksi(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

func (hs *HandlerSubmission) UpdateSubmissionAdmin(e echo.Context) error {
	Reqadmin := request.UpdateAdmin{}
	role := middlewares.ExtractTokenTeamRole(e)
	useradmin, errtoken := middlewares.ExtractTokenIdUser(e)
	idpengajuan := e.Param("id")
	cnv, errcnv := strconv.Atoi(idpengajuan)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errcnv.Error(), http.StatusBadRequest, true))
	}
	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	binderr := e.Bind(&Reqadmin)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hs.ss.UpdateSubmissionAdmin(useradmin, cnv, Reqadmin)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respon := query.ReqsubmissionToResadminupated(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

func (hs *HandlerSubmission) UpdateSubmissionPresident(e echo.Context) error {
	ReqPresident := request.UpdateVicePresident{}
	role := middlewares.ExtractTokenTeamRole(e)
	userpresident, errtoken := middlewares.ExtractTokenIdUser(e)
	idpengajuan := e.Param("id")
	cnv, errcnv := strconv.Atoi(idpengajuan)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errcnv.Error(), http.StatusBadRequest, true))
	}
	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "vicepresident" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses vicepresident", http.StatusUnauthorized, true))
	}
	binderr := e.Bind(&ReqPresident)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hs.ss.UpdateSubmissionPresident(userpresident, cnv, ReqPresident)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respon := query.ReqtoResponPresident(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

func (hsm *HandlerSubmission) UpdateSubmissionDireksi(e echo.Context) error {
	ReqPresident := request.UpdateDireksi{}
	role := middlewares.ExtractTokenTeamRole(e)
	userdireksi, errtoken := middlewares.ExtractTokenIdUser(e)
	idpengajuan := e.Param("id")
	cnv, errcnv := strconv.Atoi(idpengajuan)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errcnv.Error(), http.StatusBadRequest, true))
	}
	if errtoken != nil {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse(errtoken.Error(), http.StatusUnauthorized, true))
	}
	if role != "direksi" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses direksi", http.StatusUnauthorized, true))
	}
	binderr := e.Bind(&ReqPresident)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hsm.ss.UpdateSubmissionDireksi(userdireksi, cnv, ReqPresident)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respon := query.ReqDireksiTores(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

// GetAllSubmissionPresident implements handlecontract.HandleSubmission.
func (hsm *HandlerSubmission) GetAllSubmissionPresident(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	departpresident := middlewares.ExtractTokenTeamDepartment(e)

	if role != "vicepresident" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses vicepresident", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hsm.ss.GetAllSubmissionPresident(departpresident)
	if errservice != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(errservice.Error(), http.StatusBadRequest, true))
	}
	respon := query.ListReqltoResPresident(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}

// GetNamaManager implements handlecontract.HandleSubmission.
func (hs *HandlerSubmission) GetNamaManager(e echo.Context) error {

	namas := e.Param("nama")
	role := middlewares.ExtractTokenTeamRole(e)

	// fmt.Print("ini nama handler", namas)

	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses manager", http.StatusUnauthorized, true))
	}

	dataservice, errservice := hs.ss.GetNamaManager(namas)
	if dataservice == nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse("data tidak ada", http.StatusInternalServerError, true))
	}

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respon := query.ListReqltoResmanager(dataservice)

	return e.JSON(http.StatusOK, helper.GetResponse(respon, http.StatusOK, false))
}
