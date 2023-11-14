package reposubmission

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"
	"time"

	"gorm.io/gorm"
)

type RepoSubmission struct {
	db *gorm.DB
}

func NewRepoSubmission(db *gorm.DB) repocontract.RepoSubmission {
	return &RepoSubmission{
		db: db,
	}
}

func (rsm *RepoSubmission) AddSubmissionManager(newSubmission request.ReqSubmissionManager, res time.Time) (request.ReqSubmissionManager, error) {
	reqsubmissiontomodel := query.RequestmanagerTomodel(newSubmission, res)

	tx := rsm.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqSubmissionManager{}, tx.Error
	}
	timeString := reqsubmissiontomodel.TanggalKebutuhan.Format("02/01/2006")

	modeltoreq := query.ModelmanagerToRequest(reqsubmissiontomodel, timeString)

	return modeltoreq, nil
}

func (rsm *RepoSubmission) GetAllSubmissionManager(id int) ([]request.ReqGetManager, error) {
	modelmanager := []model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and sb.user_pengajuan = ?", id).Find(&modelmanager)

	if tx.Error != nil {
		return []request.ReqGetManager{}, tx.Error
	}
	list := query.ListModeltoReqmanager(modelmanager)

	return list, nil
}

// GetKodePengajuan implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) GetNamaManager(namamanager string) ([]request.ReqGetManager, error) {
	modelmanager := []model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and u.nama = ?", namamanager).Find(&modelmanager)

	if tx.Error != nil {
		return []request.ReqGetManager{}, tx.Error
	}
	list := query.ListModeltoReqmanager(modelmanager)

	return list, nil
}

func (rsm *RepoSubmission) GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error) {
	modelAdmin := []model.ReqGetAdmin{}
	tx := rsm.db.
		Table("submissions").
		Select("submissions.id, manager.nama as user_pengajuan, departments.nama_department, submissions.jumlah, submissions.alasan, submissions.pencharian, submissions.tanggal_kebutuhan, submissions.maksimal_gaji, user_evaluasi.nama AS nama_evaluasi, user_verifikasi.nama AS nama_verifikasi, user_persetujuan.nama AS nama_persetujuan, submissions.status_pengajuan, submissions.golongan, submissions.tanggal_verifikasi, submissions.tanggal_evaluasi, submissions.tanggal_pengajuan, submissions.tanggal_disetujui,submission.kode_pengajuan").
		Joins("LEFT JOIN users AS user_verifikasi ON submissions.id_verifikasi = user_verifikasi.id").
		Joins("LEFT JOIN users AS user_persetujuan ON submissions.idpersetujuan = user_persetujuan.id").
		Joins("LEFT JOIN users AS user_evaluasi ON submissions.id_evaluasi = user_evaluasi.id").
		Joins("LEFT JOIN users AS manager ON submissions.user_pengajuan = manager.id").
		Joins("LEFT JOIN departments ON departments.id = submissions.id_department").
		Find(&modelAdmin)
	// fmt.Print("ini list", modelAdmin)
	if tx.Error != nil {
		return []request.ReqGetAdmin{}, tx.Error
	}
	list := query.ListModeltoReqadmin(modelAdmin)

	return list, nil
}

func (rsm *RepoSubmission) GetAllSubmissionDireksi(deparment string) ([]request.ReqGetDireksi, error) {
	modeldireksi := []model.ReqGetDireksi{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.tanggal_disetujui,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id AND sb.user_pengajuan =u.id and dp.nama_department = ?", deparment).Find(&modeldireksi)

	if tx.Error != nil {
		return []request.ReqGetDireksi{}, tx.Error
	}
	list := query.ListModeltoReqDireksi(modeldireksi)

	return list, nil
}

// GetAllSubmissionPresident implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) GetAllSubmissionPresident(deparment string) ([]request.ReqGetPresident, error) {
	modelpresident := []model.ReqGetPresident{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.tanggal_verifikasi FROM users u, departments dp ,submissions sb where sb.id_department = dp.id AND sb.user_pengajuan =u.id and dp.nama_department = ?", deparment).Find(&modelpresident)

	if tx.Error != nil {
		return []request.ReqGetPresident{}, tx.Error
	}
	list := query.ListmodelltoReqPresident(modelpresident)

	return list, nil
}

func (rsm *RepoSubmission) UpdateSubmissionAdmin(idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error) {
	var submission model.Submission

	tx1 := rsm.db.Raw("Select submissions.jumlah, submissions.alasan, from submissions WHERE submissions.id= ? ", idsubmission).First(&submission)

	if errors.Is(tx1.Error, gorm.ErrRecordNotFound) {

		return request.UpdateAdmin{}, tx1.Error
	}

	reqadmintomodelsubmission := query.ReqadminTomodelsubmissionudated(update)

	tx2 := rsm.db.Model(&reqadmintomodelsubmission).Where("id = ?", idsubmission).Updates(&reqadmintomodelsubmission)

	if tx2.Error != nil {
		return request.UpdateAdmin{}, tx2.Error
	}
	modeltoreq := query.ModelsubmissionToReqadminudated(reqadmintomodelsubmission)

	return modeltoreq, nil
}

func (rsm *RepoSubmission) UpdateSubmissionPresident(idsubmission int, update request.UpdateVicePresident) (request.UpdateVicePresident, error) {
	var submission model.Submission

	tx1 := rsm.db.Raw("Select submissions.jumlah, submissions.alasan, from submissions WHERE submissions.id= ? ", idsubmission).First(&submission)

	if errors.Is(tx1.Error, gorm.ErrRecordNotFound) {

		return request.UpdateVicePresident{}, tx1.Error
	}

	reqtomodel := query.ReqpresidentTomodelsubmissionudated(update)

	tx2 := rsm.db.Model(&reqtomodel).Where("id = ?", idsubmission).Updates(&reqtomodel)

	if tx2.Error != nil {
		return request.UpdateVicePresident{}, tx2.Error
	}
	modeltoreq := query.ModelsubmissionToReqpresidentupdated(reqtomodel)

	return modeltoreq, nil
}

// UpdateSubmissionDireksi implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) UpdateSubmissionDireksi(idsubmission int, update request.UpdateDireksi) (request.UpdateDireksi, error) {
	var submission model.Submission

	tx1 := rsm.db.Raw("Select submissions.jumlah, submissions.alasan, from submissions WHERE submissions.id= ? ", idsubmission).First(&submission)

	if errors.Is(tx1.Error, gorm.ErrRecordNotFound) {

		return request.UpdateDireksi{}, tx1.Error
	}

	reqtomodel := query.ReqdireksiTomodelsubmissionudated(update)

	tx2 := rsm.db.Model(&reqtomodel).Where("id = ?", idsubmission).Updates(&reqtomodel)

	if tx2.Error != nil {
		return request.UpdateDireksi{}, tx2.Error
	}
	modeltoreq := query.ModelDireksiToreq(reqtomodel)

	return modeltoreq, nil
}

// NamaManager implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) NamaManager(namamanager string) (request.ReqGetManager, error) {
	modelmanager := model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and u.nama = ?", namamanager).Find(&modelmanager)

	if tx.Error != nil {
		return request.ReqGetManager{}, tx.Error
	}
	list := query.GetModelMnagerToReq(modelmanager)

	return list, nil
}

// NamaManager implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) CodeSubmission(kode string) (request.ReqGetManager, error) {
	modelmanager := model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and sb.kode_pengajuan= ?", kode).Find(&modelmanager)

	if tx.Error != nil {
		return request.ReqGetManager{}, tx.Error
	}
	list := query.GetModelMnagerToReq(modelmanager)

	return list, nil
}
