package repokandidat

import (
	"par/domain/contract/repocontract"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type Repokandidat struct {
	db *gorm.DB
}

func NewRepoKandidat(db *gorm.DB) repocontract.RepoKandidat {
	return &Repokandidat{
		db: db,
	}
}

// AddFormulirKandidat implements repocontract.RepoKandidat.
func (rk *Repokandidat) AddFormulirKandidat(newkandidat request.ReqFormulirKandidat) (request.ReqFormulirKandidat, error) {
	reqsubmissiontomodel := query.ReqtomodelKandidat(newkandidat)

	tx := rk.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqFormulirKandidat{}, tx.Error
	}
	modeltoreq := query.ModeltoReqKandidat(&reqsubmissiontomodel)
	return modeltoreq, nil

}
