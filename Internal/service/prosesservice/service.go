package prosesservice

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/validasi"

	"github.com/go-playground/validator"
)

type Serviceprocess struct {
	rp       repocontract.RepoProcess
	ri       repocontract.RepoInterview
	rk       repocontract.RepoKandidat
	ru       repocontract.RepoUser
	rs       repocontract.RepoSoal
	validate *validator.Validate
}

func NewServiceprocess(rp repocontract.RepoProcess, ri repocontract.RepoInterview, rk repocontract.RepoKandidat, ru repocontract.RepoUser, rs repocontract.RepoSoal) servicecontract.ServiceProcess {
	return &Serviceprocess{
		rp:       rp,
		ri:       ri,
		rk:       rk,
		ru:       ru,
		rs:       rs,
		validate: validator.New(),
	}
}

func (sp *Serviceprocess) AddProcess(id int, newProcess request.ReqDetailProsesAdmin) (request.ReqDetailProsesAdmin, error) {
	validerr := sp.validate.Struct(newProcess)
	if validerr != nil {

		return request.ReqDetailProsesAdmin{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	soal, _ := sp.rs.AllSoal()
	panjangsoal := len(soal)
	repouser, erruser := sp.ru.IdUserExist(id)
	if erruser != nil {
		return request.ReqDetailProsesAdmin{}, erruser
	}
	newProcess.IDAdmin = uint(repouser.Id)
	newProcess.NamaAdmin = repouser.Name

	pengajuan, errpengajuan := sp.ri.GetallInterview(int(newProcess.IDAdmin), newProcess.KodePengajuan, newProcess.NamaKandidat)
	// fmt.Print("ini panajang interview", len(pengajuan))
	// fmt.Print("ini panjang soal", panjangsoal)
	if errpengajuan != nil {
		return request.ReqDetailProsesAdmin{}, errpengajuan
	}
	var nilai float64
	var departmentkandidat string
	for _, val := range pengajuan {
		if newProcess.NamaKandidat != val.NamaKandidat {
			return request.ReqDetailProsesAdmin{}, errors.New("nama kandidat tidak ada")
		}

		nilai += val.Nilai
		departmentkandidat = val.DepartementKandidat
	}
	nilaiakhir := nilai / float64(panjangsoal)
	newProcess.NilaiAdmin = nilaiakhir
	newProcess.Status = "interview hr"
	newProcess.TotalNilai = newProcess.NilaiAdmin

	newProcess.KandidatDepartment = departmentkandidat
	ceknamadankode, errcode := sp.rp.Getdetailkandidat(newProcess.KodePengajuan, newProcess.NamaAdmin, newProcess.NamaKandidat)

	if errcode != nil {
		return request.ReqDetailProsesAdmin{}, errors.New("data sudah ada")
	}
	if ceknamadankode.NamaAdmin != "" || ceknamadankode.KodePengajuan != "" || ceknamadankode.NamaKandidat != "" {
		return request.ReqDetailProsesAdmin{}, errors.New("data sudah ada")
	}

	datarepo, errrepo := sp.rp.AddProcess(newProcess)

	if errrepo != nil {
		return request.ReqDetailProsesAdmin{}, errrepo
	}
	return datarepo, nil

}

func (sp *Serviceprocess) GetallDetail() (data []request.ReqDetailProses, err error) {
	data, err = sp.rp.GetallDetail()

	if err != nil {
		return data, err
	}
	return data, nil
}

// // GetallInterview implements servicecontract.Serviceinterview.
// func (si *Serviceinterview) GetallInterview(userid int, kode string, nama string) (data []request.ReqInterviewKandidat, err error) {
// 	data, err = si.ri.GetallInterview(userid, kode, nama)

// 	if err != nil {
// 		return data, err
// 	}
// 	return data, nil
// }
