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

// AddSubmissionManager implements repocontract.RepoSubmissionManager.
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

// GetAllSubmissionManager implements repocontract.RepoSubmissionManager.
func (rsm *RepoSubmission) GetAllSubmissionManager(id int) ([]request.ReqGetManager, error) {
	modelmanager := []model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and sb.user_pengajuan = ?", id).Find(&modelmanager)

	if tx.Error != nil {
		return []request.ReqGetManager{}, tx.Error
	}
	list := query.ListModeltoReqmanager(modelmanager)

	return list, nil
}

// GetAllSubmissionAdmin implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error) {
	modelAdmin := []model.ReqGetAdmin{}
	tx := rsm.db.
		Table("submissions").
		Select("submissions.id, manager.nama as user_pengajuan, departments.nama_department, submissions.jumlah, submissions.alasan, submissions.pencharian, submissions.tanggal_kebutuhan, submissions.maksimal_gaji, user_evaluasi.nama AS nama_evaluasi, user_verifikasi.nama AS nama_verifikasi, user_persetujuan.nama AS nama_persetujuan, submissions.status_pengajuan, submissions.golongan, submissions.tanggal_verifikasi, submissions.tanggal_evaluasi, submissions.tanggal_pengajuan, submissions.tanggal_disetujui").
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

// GetAllSubmissionDireksi implements repocontract.RepoSubmission.
func (rsm *RepoSubmission) GetAllSubmissionDireksi(deparment string) ([]request.ReqGetDireksi, error) {
	modeldireksi := []model.ReqGetDireksi{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan,sb.tanggal_disetujui FROM users u, departments dp ,submissions sb where sb.id_department = dp.id AND sb.user_pengajuan =u.id and dp.nama_department = ?", deparment).Find(&modeldireksi)

	if tx.Error != nil {
		return []request.ReqGetDireksi{}, tx.Error
	}
	list := query.ListModeltoReqDireksi(modeldireksi)

	return list, nil
}

// UpdateSubmissionAdmin implements repocontract.RepoSubmission.
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
