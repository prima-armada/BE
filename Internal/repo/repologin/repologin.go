package repologin

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"
	middlewares "par/middleware"

	"gorm.io/gorm"
)

type Repologin struct {
	db *gorm.DB
}

func NewRepoLogin(db *gorm.DB) repocontract.RepoLogin {
	return &Repologin{
		db: db,
	}
}

// Login implements repocontract.RepoLogin.
func (rl *Repologin) Login(nip string, password string) (string, request.RequestUser, error) {
	userdata := model.User{}

	tx := rl.db.Where("nip = ?", nip).First(&userdata)
	if tx.Error != nil {
		return "", request.RequestUser{}, tx.Error
	}
	createtoken, errtoken := middlewares.CreateTokenTeam(userdata.Nip, userdata.Role, int(userdata.ID), userdata.Bagian)

	if errtoken != nil {
		return "", request.RequestUser{}, errors.New("gagal membuat token")
	}

	datamodeltoreq := query.ModeltoReq(&userdata)
	return createtoken, datamodeltoreq, nil
}
