package servicecontract

import (
	"par/domain/request"
)

type ServiceCase interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
}
