package repouser

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type RepoUser struct {
	db *gorm.DB
}

func NewRepoUser(db *gorm.DB) repocontract.RepoUser {
	return &RepoUser{
		db: db,
	}
}

// Register implements repocontract.RepoUser.
func (ru *RepoUser) Register(newRequest request.RequestUser) (data request.RequestUser, err error) {

	datamanager := query.RequserToModelmanager(newRequest)
	dataadmin := query.RequserToModelAdmin(newRequest)
	datareqtomdel := query.RequserToModel(newRequest)

	_, erruserexist := ru.NipUserExist(datareqtomdel.Nip)
	if erruserexist == nil {
		return request.RequestUser{}, errors.New("Nip Sudah Ada")
	}
	tx := ru.db.Create(&datareqtomdel)

	if tx.Error != nil {
		return data, tx.Error
	}
	nipmanager, errnipmanager := ru.Nipmanagerexist(datamanager.Nip)
	nipadmin, errnipadmin := ru.NipadminExist(dataadmin.Nip)

	if errnipmanager == nil && errnipadmin == nil {
		if nipadmin.Nip == nipmanager.Nip {
			return request.RequestUser{}, errors.New("maaf anda sudah terdaftar,Hubungi Admin")
		}

	}
	if nipmanager.Nip == "" {

		if newRequest.Role == "manager" {

			tx2 := ru.db.Create(&datamanager)

			if tx2.Error != nil {
				return request.RequestUser{}, tx2.Error
			}
		}
	} else {
		return request.RequestUser{}, errors.New("anda sudah terdaftar di manager")
	}

	if errnipadmin == nil {
		return request.RequestUser{}, errors.New("anda sudah terdaftar di admin")
	}
	if nipadmin.Nip == "" {

		if newRequest.Role == "admin" {
			newRequest.Bagian = "humancapital"
			tx3 := ru.db.Create(&dataadmin)
			if tx3.Error != nil {
				return request.RequestUser{}, errors.New(tx3.Error.Error())
			}
		}
	}
	datamodeltoreq := query.ModelToReq(datareqtomdel)
	return datamodeltoreq, nil

}

// NipadminExist implements repocontract.RepoUser.
func (ru *RepoUser) NipadminExist(nip string) (data request.RequestUser, err error) {
	var activ model.Admin

	tx := ru.db.Raw("Select admins.id, admins.nip, admins.nama from admins WHERE admins.nip= ? ", nip).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModeladminToRequser(activ)
	return activcore, nil
}

// Nipmanagerexist implements repocontract.RepoUser.
func (ru *RepoUser) Nipmanagerexist(nip string) (data request.RequestUser, err error) {
	var activ model.Manager

	tx := ru.db.Raw("Select managers.id, managers.nip, managers.nama from managers WHERE managers.nip= ? ", nip).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModelmanagerToRequser(activ)
	fmt.Print("ini data id", activcore)
	return activcore, nil
}

func (ru *RepoUser) NipUserExist(nip string) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.nip, users.password from users WHERE users.nip= ? ", nip).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModeltoReq(activ)
	// fmt.Print("ini data id", activcore)
	return activcore, nil
}

func (ru *RepoUser) AllManager() (data []request.RequestUser, err error) {
	var activ []model.Manager
	tx := ru.db.Raw("Select managers.id, managers.nip, managers.nama from managers").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModelToReq(activ)
	return dtmdlttoreq, nil
}

// AllUser implements repocontract.RepoUser.
func (ru *RepoUser) AllUser() (data []request.RequestUser, err error) {
	var activ []model.User
	tx := ru.db.Raw("Select users.id, users.role, users.nip, users.password,users.username from users").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModelUserToReq(activ)
	return dtmdlttoreq, nil
}

// AllAdmin implements repocontract.RepoUser.
func (ru *RepoUser) AllAdmin() (data []request.RequestUser, err error) {
	var activ []model.Admin
	tx := ru.db.Raw("Select admins.id, admins.nip, admins.nama from admins").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModelToRequest(activ)
	return dtmdlttoreq, nil
}
