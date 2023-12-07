package kandidatservice

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/validasi"
	"par/validasihp"

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

func (sk *Servicekandidat) AddFormulirKandidat(newkandidata request.ReqFormulirKandidat, AdminId uint) (request.ReqFormulirKandidat, error) {

	validerr := sk.validate.Struct(newkandidata)
	if validerr != nil {

		return request.ReqFormulirKandidat{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	dataname, errname := sk.ru.NameExist(newkandidata.NamaManager)
	datahp1, errhp1 := validasihp.ValidateAndFormatPhoneNumber(newkandidata.ContactNumber)
	if errname != nil {
		return request.ReqFormulirKandidat{}, errname
	}

	if errhp1 != nil {
		return request.ReqFormulirKandidat{}, errhp1
	}
	newkandidata.ContactNumber = datahp1

	kodeajuan, errajuan := sk.rsm.CodeSubmission(newkandidata.KodePengajuan)

	if errajuan != nil {
		return request.ReqFormulirKandidat{}, errors.New("data kode pengajuan tidak ada")
	}
	fmt.Print("code", kodeajuan)

	if kodeajuan.StatusPengajuan != "disetujui" {
		return request.ReqFormulirKandidat{}, errors.New("status belum di setujui oleh direksi")
	}

	newkandidata.KodePengajuan = kodeajuan.KodePengajuan
	newkandidata.DepartementManager = dataname.Bagian

	datauser, erruser := sk.ru.IdUserExist(int(AdminId))

	if erruser != nil {
		return request.ReqFormulirKandidat{}, errors.New("data user not found")
	}
	newkandidata.AdminId = uint(datauser.Id)

	lennamaref := len(newkandidata.NamaRefrensi)
	lendepref := len(newkandidata.DepartmentRefrensi)
	lennipref := len(newkandidata.NipRefrensi)
	if newkandidata.InformasiJob == "internal" {

		if (lennamaref < 0 || lennamaref < 5) || (lendepref < 0 || lendepref < 5) || (lennipref < 0 || lennipref < 5) {
			return request.ReqFormulirKandidat{}, errors.New("Kolom refrensi tidak boleh kosong, karena anda memilih" + newkandidata.InformasiJob)
		}
		return newkandidata, nil
	}
	cekdata, errdata := sk.rsm.GetAllSubmissionAdmin()

	if errdata != nil {
		return request.ReqFormulirKandidat{}, errdata
	}
	for _, val := range cekdata {
		if newkandidata.KodePengajuan == val.KodePengajuan {
			// fmt.Print(a)
			if newkandidata.PosisiLamar != val.PosisiKosong {
				return request.ReqFormulirKandidat{}, errors.New("posisi yang di lamar tidak sesuai dengan posisi kosong pengajuan")
			}
			if newkandidata.NamaManager != val.UserPengajuan && newkandidata.DepartementManager != val.NamaDepartment {
				return request.ReqFormulirKandidat{}, errors.New("nama manager yang anda masukkan tidak sama dengan kode pengajuan")
			}
		}

	}

	cekcode, errcode := sk.rk.GetCodedannamaKandidat(newkandidata.KodePengajuan, newkandidata.NamaKandidat)
	if errcode != nil {
		return request.ReqFormulirKandidat{}, errcode
	}
	if cekcode.NamaManager == newkandidata.NamaManager && cekcode.NamaKandidat == newkandidata.NamaKandidat {
		return request.ReqFormulirKandidat{}, errors.New("nama " + cekcode.NamaKandidat + " sudah anda buat")
	}

	datarepo, errrepo := sk.rk.AddFormulirKandidat(newkandidata)

	if errrepo != nil {
		return request.ReqFormulirKandidat{}, errrepo
	}
	return datarepo, nil

}

func (sk *Servicekandidat) GetCodeKandidat(kode string) (data []request.ReqFormulirKandidat, err error) {
	data, err = sk.rk.GetCodeKandidat(kode)

	if err != nil {
		return []request.ReqFormulirKandidat{}, err
	}
	return data, nil
}
