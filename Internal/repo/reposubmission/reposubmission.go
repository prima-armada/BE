package reposubmission

import (
	"errors"
	"fmt"
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

func (rsm *RepoSubmission) AddSubmission(newSubmission request.ReqSubmission, res time.Time) (request.ReqSubmission, error) {
	reqsubmissiontomodel := query.RequestsubmissionTomodel(newSubmission, res)

	tx := rsm.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqSubmission{}, tx.Error
	}
	timeString := reqsubmissiontomodel.TanggalKebutuhan.Format("02/01/2006")

	modeltoreq := query.ModelsubmissionToRequest(reqsubmissiontomodel, timeString)

	return modeltoreq, nil
}

// GetAllSubmissionUser implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) GetAllSubmissionUser(deparment string) ([]request.ReqGetUsers, error) {
	modelsubmisison := []model.GetUsersSubmission{}

	tx := rsm.db.Raw("select submissions.id, departments.nama_department, submissions.alasan, submissions.pencharian,submissions.posisi_kosong, submissions.tanggal_kebutuhan, userpengajuan.nama as user_pengajuan, user_evaluasi.nama AS nama_evaluasi, user_verifikasi.nama AS nama_verifikasi, user_persetujuan.nama AS nama_persetujuan, submissions.status_pengajuan, submissions.tanggal_verifikasi, submissions.tanggal_evaluasi, submissions.tanggal_pengajuan, submissions.tanggal_disetujui, submissions.kode_pengajuan FROM submissions LEFT JOIN users AS user_verifikasi ON submissions.id_verifikasi = user_verifikasi.id LEFT JOIN users AS user_persetujuan ON submissions.idpersetujuan = user_persetujuan.id LEFT JOIN users AS user_evaluasi ON submissions.id_evaluasi = user_evaluasi.id LEFT JOIN users AS userpengajuan ON submissions.user_pengajuan =userpengajuan.id LEFT JOIN departments ON departments.id = submissions.id_department WHERE departments.nama_department = ?", deparment).Find(&modelsubmisison)

	if tx.Error != nil {
		return []request.ReqGetUsers{}, tx.Error
	}
	list := query.ListModeltoReqsubmission(modelsubmisison)

	return list, nil
}

func (rsm *RepoSubmission) GetNamaManager(namamanager string) ([]request.ReqGetManager, error) {
	modelmanager := []model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and u.nama = ?", namamanager).Find(&modelmanager)

	if tx.Error != nil {
		return []request.ReqGetManager{}, tx.Error
	}
	list := query.ListModeltoReqmanager(modelmanager)

	fmt.Print("ini list", list)

	return list, nil
}

func (rsm *RepoSubmission) GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error) {
	modelAdmin := []model.ReqGetAdmin{}
	tx := rsm.db.
		Table("submissions").
		Select("submissions.id, manager.nama as user_pengajuan, departments.nama_department,submissions.posisi_kosong,submissions.jumlah, submissions.alasan, submissions.pencharian, submissions.tanggal_kebutuhan, submissions.maksimal_gaji, user_evaluasi.nama AS nama_evaluasi, user_verifikasi.nama AS nama_verifikasi, user_persetujuan.nama AS nama_persetujuan, submissions.status_pengajuan, submissions.golongan, submissions.tanggal_verifikasi, submissions.tanggal_evaluasi, submissions.tanggal_pengajuan, submissions.tanggal_disetujui,submissions.kode_pengajuan").
		Joins("LEFT JOIN users AS user_verifikasi ON submissions.id_verifikasi = user_verifikasi.id").
		Joins("LEFT JOIN users AS user_persetujuan ON submissions.idpersetujuan = user_persetujuan.id").
		Joins("LEFT JOIN users AS user_evaluasi ON submissions.id_evaluasi = user_evaluasi.id").
		Joins("LEFT JOIN users AS manager ON submissions.user_pengajuan = manager.id").
		Joins("LEFT JOIN departments ON departments.id = submissions.id_department").
		Find(&modelAdmin)

	if tx.Error != nil {
		return []request.ReqGetAdmin{}, tx.Error
	}
	list := query.ListModeltoReqadmin(modelAdmin)

	return list, nil
}

func (rsm *RepoSubmission) UpdateSubmissionAdmin(idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error) {
	var submission model.Submission

	tx1 := rsm.db.Raw("Select submissions.jumlah, submissions.alasan from submissions WHERE submissions.id= ? ", idsubmission).First(&submission)

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

	tx1 := rsm.db.Raw("Select submissions.jumlah, submissions.alasan from submissions WHERE submissions.id= ? ", idsubmission).First(&submission)

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

	tx1 := rsm.db.Raw("Select submissions.jumlah, submissions.alasan from submissions WHERE submissions.id= ? ", idsubmission).First(&submission)

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

func (rsm *RepoSubmission) CodeSubmission(kode string) (request.ReqGetManager, error) {
	modelmanager := model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and sb.kode_pengajuan= ?", kode).Find(&modelmanager)

	if tx.Error != nil {
		return request.ReqGetManager{}, tx.Error
	}
	list := query.GetModelMnagerToReq(modelmanager)

	return list, nil
}
func (rsm *RepoSubmission) CodeSubmissions(kode string) ([]request.ReqGetManager, error) {
	modelmanager := []model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.kode_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and sb.kode_pengajuan= ?", kode).Find(&modelmanager)

	if tx.Error != nil {
		return []request.ReqGetManager{}, tx.Error
	}
	list := query.ListModeltoReqmanager(modelmanager)

	return list, nil
}
