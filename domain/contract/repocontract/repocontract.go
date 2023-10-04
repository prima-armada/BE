package repocontract

import "par/domain/request"

type RepoUser interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
	AllManager() (data []request.RequestUser, err error)
	AllUser() (data []request.RequestUser, err error)
}
