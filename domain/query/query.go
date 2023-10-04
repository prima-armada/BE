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
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
	}
}

func RequserToModelmanager(data request.RequestUser) model.Manager {
	return model.Manager{
		Id:        data.IdManager,
		Nip:       data.Nip,
		Nama:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}
func ModelmanagerToRequser(data model.Manager) request.RequestUser {
	return request.RequestUser{
		IdManager: data.Id,
		Nip:       data.Nip,
		Name:      data.Nama,
		CreatedAt: data.CreatedAt,
	}
}
func ListModelToReq(data []model.Manager) (datareq []request.RequestUser) {
	for _, val := range data {
		datareq = append(datareq, ModelmanagerToRequser(val))
	}
	return datareq
}
func ModelToReq(data model.User) request.RequestUser {
	return request.RequestUser{
		Id:        data.Id,
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
	}
}
func ListModelUserToReq(data []model.User) (datareq []request.RequestUser) {
	for _, val := range data {
		datareq = append(datareq, ModelToReq(val))
	}
	return datareq
}
func ReqtoResponUser(data request.RequestUser) respon.ResponseUser {
	return respon.ResponseUser{
		Id:        data.Id,
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Username:  data.Username,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}
