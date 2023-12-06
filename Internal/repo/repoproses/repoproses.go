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
	tx := rp.db.Raw("SELECT formulir_kandidats.curicullum_vitae,detail_proses.id, detail_proses.id_admin, detail_proses.nilai_admin, detail_proses.nilai_manager, detail_proses.nama_kandidat, detail_proses.total_nilai, detail_proses.kode_pengajuan, detail_proses.id_manager, detail_proses.nama_manager, detail_proses.nama_admin, detail_proses.status,detail_proses.kandidat_department,detail_proses.nilai_direksi,detail_proses.nama_direksi FROM detail_proses LEFT JOIN formulir_kandidats ON formulir_kandidats.nama_kandidat = detail_proses.nama_kandidat").Find(&model)

	if tx.Error != nil {
		return data, tx.Error
	}

	modeltoreq := query.Listmodelotreqdetail(model)

	return modeltoreq, nil

}

// Getdetailkandidat implements repocontract.RepoProcess.
func (rp *Repoproses) GetdetailkandidatAdmin(kode string, nama string, kandidat string) (data request.ReqDetailProsesAdmin, err error) {
	model := model.DetailProses{}
	tx := rp.db.Raw("SELECT detail_proses.id, detail_proses.id_admin, detail_proses.nilai_admin, detail_proses.nilai_manager, detail_proses.nama_kandidat, detail_proses.total_nilai, detail_proses.kode_pengajuan, detail_proses.id_manager, detail_proses.nama_manager, detail_proses.nama_admin, detail_proses.status FROM detail_proses WHERE detail_proses.nama_admin = ? AND detail_proses.nama_kandidat = ? AND detail_proses.kode_pengajuan = ?", nama, kandidat, kode).Find(&model)

	if tx.Error != nil {
		return data, tx.Error
	}

	modeltoreq := query.Modelprosesadmintoreq(&model)

	return modeltoreq, nil
}

// GetdetailkandidatManager implements repocontract.RepoProcess.
func (rp *Repoproses) GetdetailkandidatManager(id int) (data request.ReqDetailProsesManager, err error) {
	model := model.DetailProses{}
	tx := rp.db.Raw("SELECT detail_proses.id,  detail_proses.nilai_manager, detail_proses.nama_kandidat, detail_proses.total_nilai, detail_proses.kode_pengajuan, detail_proses.id_manager, detail_proses.nama_manager, detail_proses.status,detail_proses.kandidat_department	 FROM detail_proses WHERE detail_proses.id = ?", id).Find(&model)

	if tx.Error != nil {
		return data, tx.Error
	}

	modeltoreq := query.Modeldetailmanagertoreq(&model)

	return modeltoreq, nil
}

// UpdateDetailAdmin implements repocontract.RepoProcess.
func (rp *Repoproses) UpdateDetailAdmin(update request.ReqDetailProsesAdmin) (data request.ReqDetailProsesAdmin, err error) {
	reqtomodel := query.Reqprosesadmintomodel(update)

	tx2 := rp.db.Model(&reqtomodel).Where("id = ?", update.Id).Updates(&reqtomodel)

	if tx2.Error != nil {
		return data, tx2.Error
	}
	modeltoreq := query.Modelprosesadmintoreq(&reqtomodel)

	return modeltoreq, nil
}

// UpdateDetail implements repocontract.RepoProcess.
func (rp *Repoproses) UpdateDetail(id int, update request.ReqDetailProsesManager) (data request.ReqDetailProsesManager, err error) {
	reqmanagertomodel := query.Reqdetailmanager(update)

	tx2 := rp.db.Model(&reqmanagertomodel).Where("id = ?", id).Updates(&reqmanagertomodel)

	if tx2.Error != nil {
		return data, tx2.Error
	}
	modeltoreq := query.Modeldetailmanagertoreq(&reqmanagertomodel)

	return modeltoreq, nil
}

// GetAlldetailManager implements repocontract.RepoProcess.
func (rp *Repoproses) GetAlldetailManager(department string) (data []request.ReqDetailProsesManager, err error) {
	model := []model.DetailProses{}
	tx := rp.db.Raw("SELECT detail_proses.id,  detail_proses.nilai_manager, detail_proses.nama_kandidat, detail_proses.total_nilai, detail_proses.kode_pengajuan, detail_proses.id_manager, detail_proses.nama_manager, detail_proses.status,detail_proses.kandidat_department	 FROM detail_proses WHERE detail_proses.kandidat_department = ?", department).Find(&model)

	if tx.Error != nil {
		return data, tx.Error
	}

	modeltoreq := query.Listmodelotreqdetailmanager(model)

	return modeltoreq, nil
}

// UpdateDetailDireksi implements repocontract.RepoProcess.
func (rp *Repoproses) UpdateDetailDireksi(update request.ReqDetailProsesDireksi) (data request.ReqDetailProsesDireksi, err error) {
	reqmanagertomodel := query.ReqdetailDireksi(update)

	tx2 := rp.db.Model(&reqmanagertomodel).Where("id = ?", update.Id).Updates(&reqmanagertomodel)

	if tx2.Error != nil {
		return data, tx2.Error
	}
	modeltoreq := query.ModeldetailDireksi(&reqmanagertomodel)

	return modeltoreq, nil
}
