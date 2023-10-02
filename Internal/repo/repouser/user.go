package repouser

import (
	"errors"
	"par/domain/contract/repocontract"
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
	datareqtomdel := query.RequserToModel(newRequest)

	tx := ru.db.Create(&datareqtomdel)

	datamodeltoreq := query.ModelToReq(datareqtomdel)
	if tx.Error != nil {
		return data, errors.New("activities query error")
	}
	return datamodeltoreq, nil
}
