package repocontract

import "par/domain/request"

type RepoUser interface {
	Register(newRequest request.RequestUser) (data request.RequestUser, err error)
}
