package loginservice

import (
	"errors"
	"par/bycripts"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
)

type Servicelogin struct {
	rl repocontract.RepoLogin
}

func NewServiceLogin(rl repocontract.RepoLogin) servicecontract.ServiceLogin {
	return &Servicelogin{
		rl: rl,
	}
}

// Login implements servicecontract.ServiceLogin.
func (sl *Servicelogin) Login(nip string, password string) (string, request.RequestUser, error) {
	if nip == "" || password == "" {
		return "", request.RequestUser{}, errors.New("inputan tidak boleh kosong")
	}

	token, datarepo, errrepo := sl.rl.Login(nip, password)
	checkpw := bycripts.CheckPassword(datarepo.Password, password)

	if checkpw != nil {
		return "", request.RequestUser{}, errors.New("password anda salah")
	}
	if errrepo != nil {
		return "", request.RequestUser{}, errors.New(errrepo.Error())
	}
	return token, datarepo, nil
}
