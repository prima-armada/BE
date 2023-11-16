package interviewservice

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/kriteria"
	"par/validasi"
	"time"

	"github.com/go-playground/validator"
)

type Serviceinterview struct {
	ri       repocontract.RepoInterview
	rd       repocontract.RepoDepartment
	rk       repocontract.RepoKandidat
	ru       repocontract.RepoUser
	rs       repocontract.RepoSoal
	validate *validator.Validate
}

func NewServiceinterview(ri repocontract.RepoInterview, rk repocontract.RepoKandidat, rd repocontract.RepoDepartment, ru repocontract.RepoUser, rs repocontract.RepoSoal) servicecontract.Serviceinterview {
	return &Serviceinterview{
		ri:       ri,
		rd:       rd,
		rk:       rk,
		ru:       ru,
		rs:       rs,
		validate: validator.New(),
	}
}

// AddInterview implements servicecontract.Serviceinterview.
func (si *Serviceinterview) AddInterview(newinterview request.ReqInterviewKandidat) (request.ReqInterviewKandidat, error) {
	validerr := si.validate.Struct(newinterview)
	if validerr != nil {

		return request.ReqInterviewKandidat{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	repouser, erruser := si.ru.IdUserExist(int(newinterview.UserId))

	if erruser != nil {
		return request.ReqInterviewKandidat{}, erruser
	}
	newinterview.NamaUser = repouser.Name

	newinterview.DepartementUser = repouser.Bagian
	newinterview.Role = repouser.Role

	newinterview.UserId = uint(repouser.Id)

	kriteriakandidat := kriteria.KriteriaKandidat(newinterview.Kriteria)

	newinterview.Nilai = float64(kriteriakandidat)

	getcodedannama, errcode := si.rk.GetCodedannamaKandidat(newinterview.KodePengajuan, newinterview.NamaKandidat)

	if errcode != nil {
		return request.ReqInterviewKandidat{}, nil
	}
	newinterview.NamaKandidat = getcodedannama.NamaKandidat
	newinterview.DepartementKandidat = getcodedannama.DepartementManager
	if newinterview.Role == "manager" {
		if repouser.Bagian != newinterview.DepartementKandidat {
			return request.ReqInterviewKandidat{}, errors.New("hanya boleh interview sesama department")
		}
	}
	res, errConvtime := time.Parse("02/01/2006", newinterview.TanggalWwawancara)
	if errConvtime != nil {
		return request.ReqInterviewKandidat{}, errConvtime
	}
	soal, errsoal := si.rs.KategoriSoal(newinterview.KategoriSoal)

	if errsoal != nil {
		return request.ReqInterviewKandidat{}, errsoal
	}
	if soal.Kategori == "" {
		return request.ReqInterviewKandidat{}, errors.New("kategori tidak ada")
	}
	newinterview.IdSoal = soal.Id

	cekinterview, errcek := si.ri.CekKategorInterview(newinterview.KategoriSoal)

	if errcek != nil {
		return request.ReqInterviewKandidat{}, errcek
	}

	if cekinterview.NamaKandidat == newinterview.NamaKandidat && cekinterview.UserId == newinterview.UserId && cekinterview.KategoriSoal == newinterview.KategoriSoal {
		return request.ReqInterviewKandidat{}, errors.New("kategori soal tidak boleh sama ")
	}
	datarepo, errrepo := si.ri.AddInterview(newinterview, res)

	if errrepo != nil {
		return request.ReqInterviewKandidat{}, errrepo
	}
	return datarepo, nil
}

// GetallInterview implements servicecontract.Serviceinterview.
func (si *Serviceinterview) GetallInterview(userid int, kode string, nama string) (data []request.ReqInterviewKandidat, err error) {
	data, err = si.ri.GetallInterview(userid, kode, nama)

	if err != nil {
		return data, err
	}
	return data, nil
}
