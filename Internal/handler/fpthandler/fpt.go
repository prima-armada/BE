package fpthandler

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

type Handlerfpt struct {
	ssl servicecontract.ServiceSoalFpt
}

func NewHandlesFpt(ssl servicecontract.ServiceSoalFpt) handlecontract.HandleSoalFPT {
	return &Handlerfpt{
		ssl: ssl,
	}
}

// Addsoal implements handlecontract.HandleSoalFPT.
func (hf *Handlerfpt) Addsoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)

	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	Reqsoal := request.RequesSoalFpt{}

	binderr := e.Bind(&Reqsoal)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hf.ssl.AddSoal(Reqsoal)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	// reqtorespon := query.Reqsoaltorespon(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(dataservice, http.StatusCreated, false))
}

func (hf *Handlerfpt) AllSoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)

	if role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	dataservice, errservice := hf.ssl.AllSoal()

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	// respondata := query.ListReqDtoressoal(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(dataservice, http.StatusOK, false))
}

func (hf *Handlerfpt) Deletedsoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	id := e.QueryParam("idsoal")
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	cnv, errcnv := strconv.Atoi(id)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse("gagal convert", http.StatusBadRequest, true))
	}
	errservice := hf.ssl.DeleteSoal(cnv)
	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	return e.JSON(http.StatusOK, helper.GetResponse("berhasil di hapus", http.StatusOK, false))
}
func (hf *Handlerfpt) KategoriSoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)

	if role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin dan manager", http.StatusUnauthorized, true))
	}
	kategori := e.QueryParam("kategori")

	dataservice, errservice := hf.ssl.KategoriSoal(kategori)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}

	// reqtorespon := query.Reqsoaltorespon(dataservice)
	return e.JSON(http.StatusOK, helper.GetResponse(dataservice, http.StatusOK, false))
}

func (hf *Handlerfpt) UpdatedSoal(e echo.Context) error {
	role := middlewares.ExtractTokenTeamRole(e)
	id := e.QueryParam("idsoal")
	if role != "admin" || role == "" {
		return e.JSON(http.StatusUnauthorized, helper.GetResponse("Hanya Bisa Diakses admin", http.StatusUnauthorized, true))
	}
	Reqsoal := request.RequesSoalFpt{}
	cnv, errcnv := strconv.Atoi(id)

	if errcnv != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse("gagal convert", http.StatusBadRequest, true))
	}
	binderr := e.Bind(&Reqsoal)

	if binderr != nil {
		return e.JSON(http.StatusBadRequest, helper.GetResponse(binderr.Error(), http.StatusBadRequest, true))
	}
	dataservice, errservice := hf.ssl.Updatedsoal(cnv, Reqsoal)

	if errservice != nil {
		return e.JSON(http.StatusInternalServerError, helper.GetResponse(errservice.Error(), http.StatusInternalServerError, true))
	}
	// reqtorespon := query.Reqsoaltorespon(dataservice)
	return e.JSON(http.StatusCreated, helper.GetResponse(dataservice, http.StatusCreated, false))
}
