package kandidatservice

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"

	"github.com/go-playground/validator"
)

type Servicekandidat struct {
	rk       repocontract.RepoKandidat
	rd       repocontract.RepoDepartment
	rsm      repocontract.RepoSubmission
	ru       repocontract.RepoUser
	validate *validator.Validate
}

func NewServiceKandidat(rk repocontract.RepoKandidat, rsm repocontract.RepoSubmission, rd repocontract.RepoDepartment, ru repocontract.RepoUser) servicecontract.ServiceKandidat {
	return &Servicekandidat{
		rk:       rk,
		rd:       rd,
		rsm:      rsm,
		ru:       ru,
		validate: validator.New(),
	}
}

func (sk *Servicekandidat) AddFormulirKandidat(newkandidata request.ReqFormulirKandidat, nama string, AdminId uint) (request.ReqFormulirKandidat, error) {
	datapengajuan, errpenajuan := sk.rsm.NamaManager(nama)

	newkandidata.NamaManager = datapengajuan.Nama
	newkandidata.KodePengajuan = datapengajuan.KodePengajuan
	newkandidata.DepartementManager = datapengajuan.NamaDepartment
	if errpenajuan != nil {
		return request.ReqFormulirKandidat{}, errpenajuan
	}
	datauser, erruser := sk.ru.IdUserExist(int(AdminId))

	if erruser != nil {
		return request.ReqFormulirKandidat{}, erruser
	}
	newkandidata.AdminId = uint(datauser.Id)
	// datahp, errhp1 := validasihp.validateAndFormatPhoneNumber(newkandidata.ContactNumber)

	// if errhp1 != nil {
	// 	return request.ReqFormulirKandidat{}, errhp1
	// }
	lennamaref := len(newkandidata.NamaRefrensi)
	lendepref := len(newkandidata.DepartmentRefrensi)
	lennipref := len(newkandidata.NipRefrensi)
	if newkandidata.InformasiJob == "internal" {

		if (lennamaref < 0 || lennamaref < 5) || (lendepref < 0 || lendepref < 5) || (lennipref < 0 || lennipref < 5) {
			return request.ReqFormulirKandidat{}, errors.New("nama tidak boleh kosong atau tidak boleh kurang dari 5 angka maupun huruf")
		}
		return newkandidata, nil
	}
	datarepo, errrepo := sk.rk.AddFormulirKandidat(newkandidata)

	if errrepo != nil {
		return request.ReqFormulirKandidat{}, errrepo
	}
	return datarepo, nil

}
