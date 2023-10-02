package query

import (
	"par/domain/model"
	"par/domain/request"
	"par/domain/respon"
)

func RequserToModel(data request.RequestUser) model.User {
	return model.User{
		Id:        data.Id,
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}

func ModelToReq(data model.User) request.RequestUser {
	return request.RequestUser{
		Id:        data.Id,
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}

func ReqtoResponUser(data request.RequestUser) respon.ResponseUser {
	return respon.ResponseUser{
		Id:        data.Id,
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}
