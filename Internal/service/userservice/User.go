package userservice

import (
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"time"
)

type ServicesCase struct {
	ru repocontract.RepoUser
}

func NewServiceUser(ru repocontract.RepoUser) servicecontract.ServiceCase {
	return &ServicesCase{
		ru: ru,
	}
}

// Register implements servicecontract.ServiceCase.
func (sc *ServicesCase) Register(newRequest request.RequestUser) (data request.RequestUser, err error) {
	newRequest.CreatedAt = time.Now()

	data, err = sc.ru.Register(newRequest)

	return data, nil
}
