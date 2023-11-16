package repoproses

import (
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type Repoproses struct {
	db *gorm.DB
}

func NewRepoproses(db *gorm.DB) repocontract.RepoProcess {
	return &Repoproses{
		db: db,
	}
}

// AddProcess implements repocontract.RepoProcess.
func (rp *Repoproses) AddProcess(newProcess request.ReqDetailProsesAdmin) (request.ReqDetailProsesAdmin, error) {
	reqsubmissiontomodel := query.Reqprosesadmintomodel(newProcess)

	tx := rp.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqDetailProsesAdmin{}, tx.Error
	}

	modeltoreq := query.Modelprosesadmintoreq(&reqsubmissiontomodel)

	return modeltoreq, nil
}

// GetallDetail implements repocontract.RepoProcess.
func (rp *Repoproses) GetallDetail() (data []request.ReqDetailProses, err error) {
	model := []model.DetailProses{}
	tx := rp.db.Raw("SELECT detail_proses.id, detail_proses.id_admin, detail_proses.nilai_admin, detail_proses.nilai_manager, detail_proses.nama_kandidat, detail_proses.total_nilai, detail_proses.kode_pengajuan, detail_proses.id_manager, detail_proses.nama_manager, detail_proses.nama_admin, detail_proses.status FROM detail_proses").Find(&model)

	if tx.Error != nil {
		return data, tx.Error
	}

	modeltoreq := query.Listmodelotreqdetail(model)

	return modeltoreq, nil

}

// Getdetailkandidat implements repocontract.RepoProcess.
func (rp *Repoproses) Getdetailkandidat(kode string, nama string, kandidat string) (data request.ReqDetailProsesAdmin, err error) {
	model := model.DetailProses{}
	tx := rp.db.Raw("SELECT detail_proses.id, detail_proses.id_admin, detail_proses.nilai_admin, detail_proses.nilai_manager, detail_proses.nama_kandidat, detail_proses.total_nilai, detail_proses.kode_pengajuan, detail_proses.id_manager, detail_proses.nama_manager, detail_proses.nama_admin, detail_proses.status FROM detail_proses WHERE detail_proses.nama_admin = ? AND detail_proses.nama_kandidat = ? AND detail_proses.kode_pengajuan = ?", nama, kandidat, kode).Find(&model)

	if tx.Error != nil {
		return data, tx.Error
	}

	modeltoreq := query.Modelprosesadmintoreq(&model)

	return modeltoreq, nil
}

// SELECT detail_proses.id, detail_proses.id_admin, detail_proses.nilai_admin, detail_proses.nilai_manager, detail_proses.nama_kandidat, detail_proses.total_nilai, detail_proses.kode_pengajuan, detail_proses.id_manager, detail_proses.nama_manager, detail_proses.nama_admin, detail_proses.status FROM detail_proses WHERE detail_proses.nama_admin ="furqan" AND detail_proses.nama_kandidat = "ammar" AND detail_proses.kode_pengajuan = "generalaffair60RVPrxN";
