package repomanager

import (
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"
	"time"

	"gorm.io/gorm"
)

type RepoSubmissionManager struct {
	db *gorm.DB
}

func NewRepoSubmissionManager(db *gorm.DB) repocontract.RepoSubmissionManager {
	return &RepoSubmissionManager{
		db: db,
	}
}

// AddSubmissionManager implements repocontract.RepoSubmissionManager.
func (rsm *RepoSubmissionManager) AddSubmissionManager(newSubmission request.ReqSubmissionManager, res time.Time) (request.ReqSubmissionManager, error) {
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
func (rsm *RepoSubmissionManager) GetAllSubmissionManager(id int) ([]request.ReqGetManager, error) {
	modelmanager := []model.ReqGetManager{}

	tx := rsm.db.Raw("SELECT sb.id,u.nama, dp.nama_department ,sb.jumlah,sb.alasan,sb.status_pengajuan, sb.tanggal_kebutuhan,sb.pencharian,sb.golongan,sb.tanggal_pengajuan FROM users u, departments dp ,submissions sb where sb.id_department = dp.id and sb.user_pengajuan = u.id and sb.user_pengajuan = ?", id).Find(&modelmanager)

	if tx.Error != nil {
		return []request.ReqGetManager{}, tx.Error
	}
	list := query.ListModeltoReqmanager(modelmanager)
	fmt.Print("repo", list)
	return list, nil
}
