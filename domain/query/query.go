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

func ModeltoReq(data model.User) request.RequestUser {
	return request.RequestUser{
		Id:        data.Id,
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
	}
}
func ReqtoRepon(data request.RequestUser, token string) respon.ResponseUser {
	return respon.ResponseUser{
		Id:        data.Id,
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
		Token:     token,
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

func RequserToModelAdmin(data request.RequestUser) model.Admin {
	return model.Admin{
		Id:        data.IdAdmin,
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
func ModeladminToRequser(data model.Admin) request.RequestUser {
	return request.RequestUser{
		IdAdmin:   data.Id,
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
func ListModelToRequest(data []model.Admin) (datareq []request.RequestUser) {
	for _, val := range data {
		datareq = append(datareq, ModeladminToRequser(val))
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
		CreatedAt: data.CreatedAt,
	}
}
func ModeldepartmentToReqDepart(data model.Department) request.RequestDepartment {
	return request.RequestDepartment{
		NameDepartment: data.NamaDepartment,
		CreatedAt:      data.CreatedAt,
		UpdateAt:       data.UpdateAt,
	}
}
func ModelUpdatedepartmentToReqDepart(data model.Department) request.RequestDepartment {
	return request.RequestDepartment{
		NameDepartment: data.NamaDepartment,
		UpdateAt:       data.UpdateAt,
	}
}

func ReqDepartmentTomodelDepart(data request.RequestDepartment) model.Department {
	return model.Department{
		Id:             data.Id,
		NamaDepartment: data.NameDepartment,
		CreatedAt:      data.CreatedAt,
	}
}
func ReqDepartmentTomodelDepartudated(data request.RequestDepartment) model.Department {
	return model.Department{
		Id:             data.Id,
		NamaDepartment: data.NameDepartment,
		UpdateAt:       data.UpdateAt,
	}
}
func ListDepartmentModelToReq(data []model.Department) (datareq []request.RequestDepartment) {
	for _, val := range data {
		datareq = append(datareq, ModeldepartmentToReqDepart(val))
	}
	return datareq
}

func ReqDepartmentToRespondepart(data request.RequestDepartment) respon.ResponseDeparment {
	return respon.ResponseDeparment{

		NameDepartment: data.NameDepartment,
		CreatedAt:      data.CreatedAt,
		UpdateAt:       data.UpdateAt,
	}
}
func ReqDepartUpdatementToRespondepart(data request.RequestDepartment) respon.ResponseDeparment {
	return respon.ResponseDeparment{

		NameDepartment: data.NameDepartment,
		UpdateAt:       data.UpdateAt,
	}
}
func ListReqDepartmentToRespondepart(data []request.RequestDepartment) (datareq []respon.ResponseDeparment) {
	for _, val := range data {
		datareq = append(datareq, ReqDepartmentToRespondepart(val))
	}
	return datareq
}
