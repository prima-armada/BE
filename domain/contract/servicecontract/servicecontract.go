package servicecontract

import (
	"par/domain/request"
)

type ServiceCase interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
}

type ServiceLogin interface {
	Login(username string, password string) (string, request.RequestUser, error)
}
