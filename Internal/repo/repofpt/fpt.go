package reposoal

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type Repofpt struct {
	db *gorm.DB
}

func NewRepofpt(db *gorm.DB) repocontract.RepoSoalFpt {
	return &Repofpt{
		db: db,
	}
}

// AddSoal implements repocontract.RepoSoalFpt.
func (rf *Repofpt) AddSoal(newsoal request.RequesSoalFpt) (request.RequesSoalFpt, error) {
	reqtomodel := query.ReqfptlTomodel(newsoal)
	tx := rf.db.Create(&reqtomodel)

	if tx.Error != nil {
		return request.RequesSoalFpt{}, tx.Error
	}
	modeltoreq := query.ModelfptlToreq(&reqtomodel)

	return modeltoreq, nil
}

func (rf *Repofpt) AllSoal() (data []request.RequesSoalFpt, err error) {
	var activ []model.SoalFPT
	tx := rf.db.Raw("SELECT soal_fpts.id,soal_fpts.description,soal_fpts.bobot,soal_fpts.kategori FROM soal_fpts").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.Listmodelotreqfpt(activ)

	return dtmdlttoreq, nil
}

func (rf *Repofpt) DeletedSoal(id int) (row int, err error) {
	soal := model.SoalFPT{}

	tx := rf.db.Unscoped().Where("soal_fpts.id=?", id).Delete(&soal)
	fmt.Println(tx.Statement.SQL.String())
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete soal by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}

func (rf *Repofpt) Updatedsoal(id int, update request.RequesSoalFpt) (data request.RequesSoalFpt, err error) {
	var soal model.Department
	tx := rf.db.Raw("SELECT soal_fpts.id,soal_fpts.description,soal_fpts.bobot,soal_fpts.kategori FROM soal_fpts WHERE soal_fpts.id = ?", id).First(&soal)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequesSoalFpt{}, tx.Error
	}

	reqtomodel := query.ReqfptlTomodel(update)

	tx2 := rf.db.Model(&reqtomodel).Where("id = ?", id).Updates(&reqtomodel)

	if tx2.Error != nil {
		return request.RequesSoalFpt{}, tx2.Error
	}
	modeltoreq := query.ModelfptlToreq(&reqtomodel)

	return modeltoreq, nil
}

func (rf *Repofpt) KategoriSoal(kategori string) (data request.RequesSoalFpt, err error) {
	var activ model.SoalFPT
	tx := rf.db.Raw("SELECT soal_fpts.id,soal_fpts.description,soal_fpts.bobot,soal_fpts.kategori FROM soal_fpts WHERE soal_fpts.kategori = ?", kategori).Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}

	dtmdlttoreq := query.ModelfptlToreq(&activ)

	return dtmdlttoreq, nil
}
