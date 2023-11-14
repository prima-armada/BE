package soalhandler

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

type Handlersoal struct {
	ssl servicecontract.ServiceSoal
}

func NewHandlesSoal(ssl servicecontract.ServiceSoal) handlecontract.HandleSoal {
	return &Handlersoal{
		ssl: ssl,
	}
}

func (hsl *Handlersoal) Addsoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)

	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	Reqsoal := request.RequesSoal{}

	binderr := e.Bind(&Reqsoal)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hsl.ssl.AddSoal(Reqsoal)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	reqtorespon := query.Reqsoaltorespon(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(reqtorespon, http.StatusCreated, false))
}

// AllSoal implements handlecontract.HandleSoal.
func (hsl *Handlersoal) AllSoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)

	if role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hsl.ssl.AllSoal()

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	respondata := query.ListReqDtoressoal(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(respondata, http.StatusOK, false))
}

// KategoriSoal implements handlecontract.HandleSoal.
func (hsl *Handlersoal) KategoriSoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)

	if role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	kategori := e.QueryParam("kategori")

	dataservice, errservice := hsl.ssl.KategoriSoal(kategori)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	reqtorespon := query.Reqsoaltorespon(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(reqtorespon, http.StatusOK, false))

}

// UpdatedSoal implements handlecontract.HandleSoal.
func (hsl *Handlersoal) UpdatedSoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	id := e.QueryParam("idsoal")
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	Reqsoal := request.RequesSoal{}
	cnv, errcnv := strconv.Atoi(id)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse("gagal convert", http.StatusBadRequest, true))
	}
	binderr := e.Bind(&Reqsoal)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hsl.ssl.Updatedsoal(cnv, Reqsoal)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	reqtorespon := query.Reqsoaltorespon(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(reqtorespon, http.StatusCreated, false))
}

// Deletedsoal implements handlecontract.HandleSoal.
func (hsl *Handlersoal) Deletedsoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	id := e.QueryParam("idsoal")
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	cnv, errcnv := strconv.Atoi(id)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse("gagal convert", http.StatusBadRequest, true))
	}
	errservice := hsl.ssl.DeleteSoal(cnv)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	return e.JSON(http.StatusOK, helper.GetResponse("berhasil di hapus", http.StatusOK, false))
}
