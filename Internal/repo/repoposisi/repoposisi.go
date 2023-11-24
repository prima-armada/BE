package repoposisi

import (
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type Repoposisi struct {
	db *gorm.DB
}

func NewRepoposisi(db *gorm.DB) repocontract.RepoPosisi {
	return &Repoposisi{
		db: db,
	}
}

func (rp *Repoposisi) AddPosisi(newProcess request.ReqPosisi) (request.ReqPosisi, error) {
	reqsubmissiontomodel := query.Reqposisitomodel(newProcess)

	tx := rp.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqPosisi{}, tx.Error
	}

	modeltoreq := query.ModelPosisitoreq(&reqsubmissiontomodel)

	return modeltoreq, nil
}

// Getdetailposisi implements repocontract.RepoPosisi.
func (rp *Repoposisi) Getdetailposisi(userid int) (request.ReqPosisi, error) {
	modelposisi := model.Position{}

	tx := rp.db.Raw("SELECT positions.id,positions.user_id,positions.level_kosong,positions.department FROM positions WHERE positions.user_id = ?", userid).Find(&modelposisi)

	if tx.Error != nil {
		return request.ReqPosisi{}, tx.Error
	}
	modeltoreq := query.ModelPosisitoreq(&modelposisi)
	return modeltoreq, nil
}
