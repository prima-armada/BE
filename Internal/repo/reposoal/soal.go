package reposoal

import (
	"errors"
	"fmt"
	"log"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type Reposoal struct {
	db *gorm.DB
}

func NewReposoal(db *gorm.DB) repocontract.RepoSoal {
	return &Reposoal{
		db: db,
	}
}

// AddSoal implements repocontract.RepoSoal.
func (rsl *Reposoal) AddSoal(newksoal request.RequesSoal) (request.RequesSoal, error) {
	reqtomodel := query.ReqsoalTomodel(newksoal)
	tx := rsl.db.Create(&reqtomodel)

	if tx.Error != nil {
		return request.RequesSoal{}, tx.Error
	}
	modeltoreq := query.Modelsoaltoreq(&reqtomodel)

	return modeltoreq, nil
}

func (rsl *Reposoal) AllSoal() (data []request.RequesSoal, err error) {
	var activ []model.SoalInterview
	tx := rsl.db.Raw("Select soal_interviews.id, soal_interviews.kategori, soal_interviews.description from soal_interviews").Find(&activ)
	if tx.Error != nil {
		return []request.RequesSoal{}, tx.Error
	}

	// log.Print(activ, "soal repo")
	dtmdlttoreq := query.Listmodelotreqsoal(activ)
	log.Print(dtmdlttoreq, "soal repo")
	return dtmdlttoreq, nil
}

// KategoriSoal implements repocontract.RepoSoal.
func (rsl *Reposoal) KategoriSoal(kategori string) (data request.RequesSoal, err error) {

	var activ model.SoalInterview
	tx := rsl.db.Raw("Select soal_interviews.id, soal_interviews.kategori, soal_interviews.description from soal_interviews WHERE soal_interviews.kategori = ?", kategori).Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}

	dtmdlttoreq := query.Modelsoaltoreq(&activ)

	return dtmdlttoreq, nil
}

// Updatedsoal implements repocontract.RepoSoal.
func (rsl *Reposoal) Updatedsoal(id int, update request.RequesSoal) (data request.RequesSoal, err error) {
	var soal model.Department
	tx := rsl.db.Raw("Select soal_interviews.id, soal_interviews.kategori, soal_interviews.description from soal_interviews WHERE soal_interviews.id = ?", id).First(&soal)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequesSoal{}, tx.Error
	}

	reqtomodel := query.ReqsoalTomodel(update)

	tx2 := rsl.db.Model(&reqtomodel).Where("id = ?", id).Updates(&reqtomodel)

	if tx2.Error != nil {
		return request.RequesSoal{}, tx2.Error
	}
	modeltoreq := query.Modelsoaltoreq(&reqtomodel)

	return modeltoreq, nil
}

// DeleteSoal implements repocontract.RepoSoal.
func (rsl *Reposoal) DeletedSoal(id int) (row int, err error) {
	soal := model.SoalInterview{}

	tx := rsl.db.Unscoped().Where("soal_interviews.id=?", id).Delete(&soal)
	fmt.Println(tx.Statement.SQL.String())
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete soal by id failed, error query")
	}
	return int(tx.RowsAffected), nil
}
