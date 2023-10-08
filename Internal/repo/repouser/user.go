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

	gettall, _ := ru.AllManager()

	lendata := len(gettall)

	nipmanager, errnipmanager := ru.Nipmanagerexist(newRequest.Nip)

	nipadmin, errnipadmin := ru.NipadminExist(newRequest.Nip)

	if errnipmanager == nil && errnipadmin == nil {
		if nipadmin.Nip == nipmanager.Nip {
			return request.RequestUser{}, errors.New("maaf anda sudah terdaftar,Hubungi Admin")
		}

	}

	if nipmanager.Nip == "" {

		if newRequest.Role == "manager" {
			if lendata <= 0 || lendata > 0 {
				lendata += 1

				datamanager.Id = lendata

				tx2 := ru.db.Create(&datamanager)

				if tx2.Error != nil {
					return request.RequestUser{}, tx2.Error
				}
			}
		}

	} else {
		return request.RequestUser{}, errors.New("anda sudah terdaftar di manager")
	}

	if errnipadmin == nil {
		return request.RequestUser{}, errors.New("anda sudah terdaftar di admin")
	}
	admin, _ := ru.AllAdmin()
	if nipadmin.Nip == "" {
		lenadmins := len(admin)
		if newRequest.Role == "admin" {
			if lenadmins <= 0 || lenadmins > 0 {
				lenadmins += 1

				dataadmin.Id = lenadmins
				tx3 := ru.db.Create(&dataadmin)

				if tx3.Error != nil {
					return request.RequestUser{}, errors.New(tx3.Error.Error())
				}
			}
		}
	}

	datareqtomdel := query.RequserToModel(newRequest)
	gettuser, erralluser := ru.AllUser()
	if erralluser != nil {
		return request.RequestUser{}, errors.New(erralluser.Error())
	}

	lenuser := len(gettuser)

	if lenuser <= 0 || lenuser > 0 {
		lenuser += 1

		datareqtomdel.Id = lenuser

	}

	tx := ru.db.Create(&datareqtomdel)

	datamodeltoreq := query.ModelToReq(datareqtomdel)
	if tx.Error != nil {
		return data, tx.Error
	}
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
