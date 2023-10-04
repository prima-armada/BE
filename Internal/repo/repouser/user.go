package repouser

import (
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"
	"strconv"

	"github.com/google/uuid"
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

	gettall, _ := ru.AllManager()

	lendata := len(gettall)
	if newRequest.Role == "manager" {
		if lendata <= 0 || lendata > 0 {
			lendata += 1
			cnvstring := strconv.Itoa(lendata)
			cnvstring = uuid.New().String()
			datamanager.Id = cnvstring

			tx2 := ru.db.Create(&datamanager)

			if tx2.Error != nil {
				return data, tx2.Error
			}
		}
	}

	datareqtomdel := query.RequserToModel(newRequest)
	gettuser, _ := ru.AllManager()

	lenuser := len(gettuser)

	if lenuser < 0 || lenuser > 0 {
		lenuser += 1
		cnvuser := strconv.Itoa(lenuser)
		fmt.Print("ini repo", cnvuser)
		cnvuser = uuid.New().String()
		datareqtomdel.Id = cnvuser

	}

	tx := ru.db.Create(&datareqtomdel)

	datamodeltoreq := query.ModelToReq(datareqtomdel)
	if tx.Error != nil {
		return data, tx.Error
	}
	return datamodeltoreq, nil
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
	tx := ru.db.Raw("Select users.id, users.role, users.nip, users.password,users.username from managers").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModelUserToReq(activ)
	return dtmdlttoreq, nil
}
