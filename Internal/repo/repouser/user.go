package repouser

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/model"
	"par/domain/query"
	"par/domain/request"

	"gorm.io/gorm"
)

type RepoUser struct {
	db *gorm.DB
}

// UsernameUserExist implements repocontract.RepoUser.

func NewRepoUser(db *gorm.DB) repocontract.RepoUser {
	return &RepoUser{
		db: db,
	}
}

// Register implements repocontract.RepoUser.
func (ru *RepoUser) Register(newRequest request.RequestUser) (data request.RequestUser, err error) {

	datareqtomdel := query.RequserToModel(newRequest)

	_, erruserexist := ru.NipUserExist(datareqtomdel.Nip)
	_, errusernameexist := ru.UsernameUserExist(datareqtomdel.Username)
	if erruserexist == nil || errusernameexist == nil {
		return request.RequestUser{}, errors.New("anda sudah terdaftar")
	}

	tx := ru.db.Create(&datareqtomdel)

	if tx.Error != nil {
		return data, tx.Error
	}

	datamodeltoreq := query.ModeltoReq(&datareqtomdel)
	return datamodeltoreq, nil

}
func (ru *RepoUser) UsernameUserExist(username string) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.nip, users.password from users WHERE users.username= ? ", username).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModeltoReq(&activ)
	// fmt.Print("ini data id", activcore)
	return activcore, nil
}

func (ru *RepoUser) NipUserExist(nip string) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.nip, users.password from users WHERE users.nip= ? ", nip).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModeltoReq(&activ)
	// fmt.Print("ini data id", activcore)
	return activcore, nil
}

// AllUser implements repocontract.RepoUser.
func (ru *RepoUser) AllUser() (data []request.RequestUser, err error) {
	var activ []model.User
	tx := ru.db.Raw("Select users.id, users.role, users.nip, users.password,users.username,users.bagian from users").Find(&activ)
	if tx.Error != nil {
		return data, tx.Error
	}
	dtmdlttoreq := query.ListModelUserToReq(activ)
	return dtmdlttoreq, nil
}

// IdUserExist implements repocontract.RepoUser.
func (ru *RepoUser) IdUserExist(id int) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.nip, users.password,users.bagian,users.nama,users.role from users WHERE users.id= ? ", id).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModeltoReq(&activ)

	return activcore, nil
}

// GetAllManager implements repocontract.RepoUser.
func (ru *RepoUser) GetAllManager(roles string) ([]request.RequestUser, error) {
	var activ []model.User
	tx := ru.db.Raw("Select users.id, users.role, users.nip, users.password,users.username,users.nama,users.bagian from users where users.role = ?", roles).Find(&activ)
	if tx.Error != nil {
		return []request.RequestUser{}, tx.Error
	}
	dtmdlttoreq := query.ListModelUserToReq(activ)
	return dtmdlttoreq, nil
}

// NameExist implements repocontract.RepoUser.
func (ru *RepoUser) NameExist(name string) (data request.RequestUser, err error) {
	var activ model.User

	tx := ru.db.Raw("Select users.id, users.nip, users.password,users.username,users.nama,users.role,users.bagian from users WHERE users.nama= ? ", name).First(&activ)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {

		return request.RequestUser{}, tx.Error
	}
	var activcore = query.ModeltoReq(&activ)

	return activcore, nil
}
