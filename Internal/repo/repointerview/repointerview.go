package repointerview

import (
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"
	"time"

	"gorm.io/gorm"
)

type Repointerview struct {
	db *gorm.DB
}

func NewRepoInterview(db *gorm.DB) repocontract.RepoInterview {
	return &Repointerview{
		db: db,
	}
}
func (ri *Repointerview) AddInterview(newinterview request.ReqInterviewKandidat, tanggal time.Time) (request.ReqInterviewKandidat, error) {
	reqsubmissiontomodel := query.Reqinterviewtomodel(newinterview, tanggal)

	tx := ri.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqInterviewKandidat{}, tx.Error
	}
	timeString := reqsubmissiontomodel.TanggalWwawancara.Format("02/01/2006")

	modeltoreq := query.ModelinterviewToRequest(&reqsubmissiontomodel, timeString)

	return modeltoreq, nil

}

func (ri *Repointerview) AddInterviewfpt(newinterview request.ReqInterviewfpt, tanggal time.Time) (request.ReqInterviewfpt, error) {
	reqsubmissiontomodel := query.Reqinterviewfpttomodel(newinterview, tanggal)

	tx := ri.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqInterviewfpt{}, tx.Error
	}
	timeString := reqsubmissiontomodel.TanggalWwawancara.Format("02/01/2006")

	modeltoreq := query.Modelfpttoreq(&reqsubmissiontomodel, timeString)

	return modeltoreq, nil
}
func (ri *Repointerview) CekKategorInterview(kategori string) (request.ReqInterviewKandidat, error) {
	modelinterview := model.InterviewKandidat{}
	tx := ri.db.Raw("SELECT interview_kandidats.id,interview_kandidats.nama_user,interview_kandidats.departement_user ,interview_kandidats.kategori_soal,interview_kandidats.nama_kandidat,interview_kandidats.user_id FROM interview_kandidats where interview_kandidats.kategori_soal= ?", kategori).Find(&modelinterview)
	if tx.Error != nil {
		return request.ReqInterviewKandidat{}, tx.Error
	}
	modeltoreq := query.ModelinterviewToRequest2(&modelinterview)

	return modeltoreq, nil
}

// GetallInterview implements repocontract.RepoInterview.
func (ri *Repointerview) GetallInterview(userid int, kode string, nama string) (data []request.ReqInterviewKandidat, err error) {
	modelinterview := []model.InterviewKandidat{}
	tx := ri.db.Raw("SELECT interview_kandidats.id, interview_kandidats.nama_user, interview_kandidats.departement_user, interview_kandidats.departement_kandidat, interview_kandidats.kode_pengajuan, interview_kandidats.id_soal, interview_kandidats.nama_kandidat, SUM(interview_kandidats.nilai) AS nilai, interview_kandidats.user_id FROM interview_kandidats where interview_kandidats.kode_pengajuan= ? AND interview_kandidats.nama_kandidat= ? AND interview_kandidats.user_id =?", kode, nama, userid).Find(&modelinterview)
	if tx.Error != nil {
		return []request.ReqInterviewKandidat{}, tx.Error
	}
	modeltoreq := query.Listmodelotreqinterview(modelinterview)

	return modeltoreq, nil
}

func (ri *Repointerview) CekallInterview(userid int, kode string, nama string) (data []request.ReqInterviewKandidat, err error) {
	modelinterview := []model.InterviewKandidat{}
	tx := ri.db.Raw("SELECT interview_kandidats.id, interview_kandidats.nama_user, interview_kandidats.kategori_soal, interview_kandidats.departement_user, interview_kandidats.departement_kandidat, interview_kandidats.kode_pengajuan, interview_kandidats.id_soal, interview_kandidats.nama_kandidat, interview_kandidats.nilai AS nilai, interview_kandidats.user_id FROM interview_kandidats where interview_kandidats.kode_pengajuan= ? AND interview_kandidats.nama_kandidat= ? AND interview_kandidats.user_id =?", kode, nama, userid).Find(&modelinterview)
	if tx.Error != nil {
		return []request.ReqInterviewKandidat{}, tx.Error
	}
	modeltoreq := query.Listmodelotreqinterview(modelinterview)

	return modeltoreq, nil
}

func (ri *Repointerview) GetallInterviewftp(nama string, kode string) (data []request.ReqInterviewfpt, err error) {
	modelinterview := []model.InterviewFPT{}
	tx := ri.db.Raw("SELECT interview_fpts.nama_kandidat,interview_fpts.departement_user,interview_fpts.departement_kandidat,interview_fpts.kategori_soal,interview_fpts.tanggal_wwawancara,interview_fpts.kode_pengajuan,interview_fpts.nilai FROM interview_fpts WHERE interview_fpts.nama_kandidat = ?  and interview_fpts.kode_pengajuan = ?", nama, kode).Find(&modelinterview)
	if tx.Error != nil {
		return []request.ReqInterviewfpt{}, tx.Error
	}
	modeltoreq := query.Listmodelotreqinterviewfpt(modelinterview)

	return modeltoreq, nil
}

// Getallnilaiftp implements repocontract.RepoInterview.
func (ri *Repointerview) Getallnilaiftp(kode string, nama string) (data []request.ReqInterviewfpt, err error) {
	modelinterview := []model.InterviewFPT{}
	tx := ri.db.Raw("SELECT interview_fpts.nama_kandidat,interview_fpts.departement_user,interview_fpts.departement_kandidat,interview_fpts.kategori_soal,interview_fpts.tanggal_wwawancara,interview_fpts.kode_pengajuan,SUM(interview_fpts.nilai) AS nilai FROM interview_fpts WHERE interview_fpts.nama_kandidat = ? and interview_fpts.kode_pengajuan = ?", nama, kode).Find(&modelinterview)
	if tx.Error != nil {
		return []request.ReqInterviewfpt{}, tx.Error
	}
	modeltoreq := query.Listmodelotreqinterviewfpt(modelinterview)

	return modeltoreq, nil
}
