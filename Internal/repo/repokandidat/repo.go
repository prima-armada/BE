package repokandidat

import (
	"par/domain/contract/repocontract"
	"par/domain/model"
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

func (rk *Repokandidat) AddFormulirKandidat(newkandidat request.ReqFormulirKandidat) (request.ReqFormulirKandidat, error) {
	reqsubmissiontomodel := query.ReqtomodelKandidat(newkandidat)

	tx := rk.db.Create(&reqsubmissiontomodel)

	if tx.Error != nil {
		return request.ReqFormulirKandidat{}, tx.Error
	}
	modeltoreq := query.ModeltoReqKandidat(&reqsubmissiontomodel)
	return modeltoreq, nil

}

func (rk *Repokandidat) GetCodeKandidat(kode string) ([]request.ReqFormulirKandidat, error) {
	modelkandidat := []model.FormulirKandidat{}

	tx := rk.db.Raw("SELECT formulir_kandidats.id,formulir_kandidats.nama_manager ,formulir_kandidats.kode_pengajuan,formulir_kandidats.nama_kandidat,formulir_kandidats.posisi_lamar FROM formulir_kandidats where formulir_kandidats.kode_pengajuan= ?", kode).Find(&modelkandidat)

	if tx.Error != nil {
		return []request.ReqFormulirKandidat{}, tx.Error
	}
	list := query.ListKandidattoreq(modelkandidat)

	return list, nil

}

// GetCodedannamaKandidat implements repocontract.RepoKandidat.
func (rk *Repokandidat) GetCodedannamaKandidat(kode string, nama string) (request.ReqFormulirKandidat, error) {
	modelkandidat := model.FormulirKandidat{}

	tx := rk.db.Raw("SELECT formulir_kandidats.id,formulir_kandidats.nama_manager, formulir_kandidats.departement_manager , formulir_kandidats.kode_pengajuan, formulir_kandidats.nama_kandidat,formulir_kandidats.posisi_lamar FROM formulir_kandidats where formulir_kandidats.kode_pengajuan= ? AND formulir_kandidats.nama_kandidat= ?", kode, nama).Find(&modelkandidat)

	if tx.Error != nil {
		return request.ReqFormulirKandidat{}, tx.Error
	}
	list := query.ModeltoReqKandidat(&modelkandidat)

	// fmt.Print("ini nama kandidat", list.NamaKandidat)

	return list, nil

}
