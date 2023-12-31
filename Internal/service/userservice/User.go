package userservice

import (
	"errors"
	"par/bycripts"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/validasi"
	"time"

	"github.com/go-playground/validator"
)

type ServicesCase struct {
	ru       repocontract.RepoUser
	validate *validator.Validate
}

func NewServiceUser(ru repocontract.RepoUser) servicecontract.ServiceCase {
	return &ServicesCase{
		ru:       ru,
		validate: validator.New(),
	}
}

// Register implements servicecontract.ServiceCase.
func (sc *ServicesCase) Register(newRequest request.RequestUser) (data request.RequestUser, err error) {

	validerr := sc.validate.Struct(newRequest)
	if validerr != nil {

		return request.RequestUser{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	newRequest.CreatedAt = time.Now()
	haspw := bycripts.Bcript(newRequest.Password)
	newRequest.Password = haspw
	if newRequest.Role == "direksi" {
		newRequest.Bagian = "none"
	}
	datarepo, errrepo := sc.ru.Register(newRequest)

	if errrepo != nil {
		return request.RequestUser{}, errors.New(errrepo.Error())
	}

	return datarepo, nil
}

// GetAllManager implements servicecontract.ServiceCase.
func (sc *ServicesCase) GetAllManager(roles string) ([]request.RequestUser, error) {
	datarepo, errrepo := sc.ru.GetAllManager(roles)

	if errrepo != nil {
		return nil, errrepo
	}
	return datarepo, nil
}
