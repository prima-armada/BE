package query

import (
	"par/domain/model"
	"par/domain/request"
	"par/domain/respon"
)

func RequserToModel(data request.RequestUser) model.User {
	return model.User{

		Role:     data.Role,
		Nip:      data.Nip,
		Password: data.Password,
		Username: data.Username,
		Nama:     data.Name,
		Bagian:   data.Bagian,
	}
}

func ModeltoReq(data model.User) request.RequestUser {
	return request.RequestUser{
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
	}
}
func ReqtoRepon(data request.RequestUser, token string) respon.ResponseUser {
	return respon.ResponseUser{
		Role:      data.Role,
		Nip:       data.Nip,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
	}
}
func ReqtoResponLogin(data request.RequestUser, token string) respon.LoginRespon {
	return respon.LoginRespon{
		Role:     data.Role,
		Nip:      data.Nip,
		Username: data.Username,
		Token:    token,
	}
}

func ListModelUserToReq(data []model.User) (datareq []request.RequestUser) {
	for _, val := range data {

		datareq = append(datareq, ModeltoReq(val))
	}
	return datareq
}
func ReqtoResponUser(data request.RequestUser) respon.ResponseUser {
	return respon.ResponseUser{
		Role:      data.Role,
		Nip:       data.Nip,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
	}
}
func ModeldepartmentToReqDepart(data model.Department) request.RequestDepartment {
	return request.RequestDepartment{
		NameDepartment: data.NamaDepartment,
		CreatedAt:      data.CreatedAt,
		UpdateAt:       data.UpdatedAt,
	}
}

func ModelUpdatedepartmentToReqDepart(data model.Department) request.RequestDepartment {
	return request.RequestDepartment{
		UpdateAt:       data.UpdatedAt,
		NameDepartment: data.NamaDepartment,
	}
}

func ReqDepartmentTomodelDepart(data request.RequestDepartment) model.Department {
	return model.Department{
		NamaDepartment: data.NameDepartment,
	}
}

func ReqDepartmentTomodelDepartudated(data request.RequestDepartment) model.Department {
	return model.Department{

		NamaDepartment: data.NameDepartment,
	}
}
func ModelsdepartmentToReqDepart(data *model.Department) request.RequestDepartment {

	return request.RequestDepartment{
		Id:             int(data.ID),
		NameDepartment: data.NamaDepartment,
		CreatedAt:      data.CreatedAt,
		UpdateAt:       data.UpdatedAt,
	}
}
func ListDepartmentModelToReq(data []model.Department) (datareq []request.RequestDepartment) {
	for _, val := range data {
		datareq = append(datareq, ModelsdepartmentToReqDepart(&val))
	}
	return datareq
}

func ReqDepartmentToRespondepart(data request.RequestDepartment) respon.ResponseDeparment {
	return respon.ResponseDeparment{
		Id:             data.Id,
		NameDepartment: data.NameDepartment,
		CreatedAt:      data.CreatedAt,
		UpdateAt:       data.UpdateAt,
	}
}
func RequstDepartmentToRespondepart(data *request.RequestDepartment) respon.ResponseDeparment {
	return respon.ResponseDeparment{
		Id:             data.Id,
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
		datareq = append(datareq, RequstDepartmentToRespondepart(&val))
	}
	return datareq
}
