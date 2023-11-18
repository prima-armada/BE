package prosesservice

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/kriteria"
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

	newProcess.TotalNilai = newProcess.NilaiAdmin

	cekstatus, errstatus := kriteria.CekSTATUS(newProcess.TotalNilai)

	if errstatus != nil {
		return request.ReqDetailProsesAdmin{}, errstatus
	}

	newProcess.Status = cekstatus
	newProcess.KandidatDepartment = departmentkandidat

	ceklen, _ := sp.ri.CekallInterview(int(newProcess.IDAdmin), newProcess.KodePengajuan, newProcess.NamaKandidat)

	panjanginterview := len(ceklen)

	if panjanginterview != panjangsoal {
		return request.ReqDetailProsesAdmin{}, errors.New("penilaiain interview tidak sama dengan soal interview")
	}
	ceknamadankode, errcode := sp.rp.GetdetailkandidatAdmin(newProcess.KodePengajuan, newProcess.NamaAdmin, newProcess.NamaKandidat)

	if errcode != nil {
		return request.ReqDetailProsesAdmin{}, errcode
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

// UpdateDetail implements servicecontract.ServiceProcess.
func (sp *Serviceprocess) UpdateDetail(id int, update request.ReqDetailProsesManager) (data request.ReqDetailProsesManager, err error) {

	repousers, erruser := sp.ru.IdUserExist(int(update.IdManager))

	update.NamaManager = repousers.Name

	if erruser != nil {
		return data, erruser
	}
	cekinterview, errinterview := sp.ri.GetallInterview(repousers.Id, update.KodePengajuan, update.NamaKandidat)
	if errinterview != nil {
		return data, errinterview
	}

	var nilai float64

	for _, val := range cekinterview {
		if update.KodePengajuan == val.KodePengajuan {
			nilai += val.Nilai
		}
	}

	soal, _ := sp.rs.AllSoal()
	panjangsoal := len(soal)

	ceklen, _ := sp.ri.CekallInterview(repousers.Id, update.KodePengajuan, update.NamaKandidat)

	panjanginterview := len(ceklen)

	if panjanginterview != panjangsoal {
		return data, errors.New("penilaiain interview tidak sama dengan soal interview")
	}

	update.NilaiManager = nilai / float64(panjangsoal)
	cekproseses, errproses := sp.rp.GetallDetail()

	if errproses != nil {
		return data, errproses
	}
	var total float64

	for _, val := range cekproseses {
		if update.KodePengajuan == val.KodePengajuan {
			total = val.TotalNilai

		}

	}

	getdetail, errdetail := sp.rp.GetdetailkandidatManager(id)

	if getdetail.Id == 0 {
		return data, errors.New("id tidak ada")
	}

	if errdetail != nil {
		return data, errdetail
	}
	if repousers.Bagian != getdetail.KandidatDepartment {
		return data, errors.New("anda tidak berhak untuk update kandidat ini")
	}
	update.KandidatDepartment = getdetail.KandidatDepartment
	update.TotalNilai = (update.NilaiManager + total) / 2

	status, errstatus := kriteria.CekSTATUSformanager(update.TotalNilai)

	if errstatus != nil {
		return data, errstatus
	}
	update.Status = status

	datarepo, errepo := sp.rp.UpdateDetail(id, update)

	if errepo != nil {
		return data, errepo
	}
	return datarepo, nil
}

// GetAlldetailManager implements servicecontract.ServiceProcess.
func (sp *Serviceprocess) GetAlldetailManager(id int) (data []request.ReqDetailProsesManager, err error) {
	repousers, erruser := sp.ru.IdUserExist(id)

	if erruser != nil {
		return data, erruser
	}
	datarepo, errrepo := sp.rp.GetAlldetailManager(repousers.Bagian)

	if errrepo != nil {
		return data, errrepo
	}
	return datarepo, nil
}
