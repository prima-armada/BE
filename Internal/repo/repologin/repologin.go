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
func (rl *Repologin) Login(username string, password string) (string, request.RequestUser, error) {
	userdata := model.User{}

	tx := rl.db.Where("username = ?", username).First(&userdata)
	if tx.Error != nil {
		return "", request.RequestUser{}, tx.Error
	}
	createtoken, errtoken := middlewares.CreateTokenTeam(userdata.Id, userdata.Role)

	if errtoken != nil {
		return "", request.RequestUser{}, errors.New("gagal membuat token")
	}

	datamodeltoreq := query.ModelToReq(userdata)
	return createtoken, datamodeltoreq, nil
}
