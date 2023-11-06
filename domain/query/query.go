package query

import (
	"par/domain/model"
	"par/domain/request"
	"par/domain/respon"
	"time"
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

func ModeltoReq(data *model.User) request.RequestUser {
	return request.RequestUser{
		Id:        int(data.ID),
		Role:      data.Role,
		Nip:       data.Nip,
		Password:  data.Password,
		Username:  data.Username,
		Bagian:    data.Bagian,
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
		Bagian:   data.Bagian,
		Token:    token,
	}
}

func ListModelUserToReq(data []model.User) (datareq []request.RequestUser) {
	for _, val := range data {

		datareq = append(datareq, ModeltoReq(&val))
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
func ModeldepartmentToReqDepart(data *model.Department) request.RequestDepartment {
	return request.RequestDepartment{
		Id:             int(data.ID),
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

func RequestmanagerTomodel(data request.ReqSubmissionManager, tanggal time.Time) model.Submission {
	return model.Submission{
		IdDepartment:     data.IdDepartment,
		UserPengajuan:    uint(data.IdPengajuan),
		Jumlah:           data.Jumlah,
		Alasan:           data.Alasan,
		TanggalKebutuhan: tanggal,
		Pencharian:       data.Pencaharian,
		StatusPengajuan:  data.StatusPengajuan,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
	}

}
func ModelmanagerToRequest(data model.Submission, tanggal string) request.ReqSubmissionManager {
	return request.ReqSubmissionManager{
		IdDepartment:     data.IdDepartment,
		IdPengajuan:      int(data.UserPengajuan),
		Jumlah:           data.Jumlah,
		Alasan:           data.Alasan,
		TanggalKebutuhan: tanggal,
		Pencaharian:      data.Pencharian,
		StatusPengajuan:  data.StatusPengajuan,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
	}

}
func ReqmanagerToRespon(data request.ReqSubmissionManager) respon.ResponSubmissionManager {
	return respon.ResponSubmissionManager{
		IdDepartment:     data.IdDepartment,
		IdPengajuan:      data.IdPengajuan,
		Jumlah:           data.Jumlah,
		Alasan:           data.Alasan,
		TanggalKebutuhan: data.TanggalKebutuhan,
		Pencaharian:      data.Pencaharian,
		StatusPengajuan:  data.StatusPengajuan,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
	}

}

func GetModelMnagerToReq(data model.ReqGetManager) request.ReqGetManager {
	return request.ReqGetManager{
		Id:               data.Id,
		Nama:             data.Nama,
		NamaDepartment:   data.NamaDepartment,
		Jumlah:           data.Jumlah,
		Alasan:           data.Alasan,
		StatusPengajuan:  data.StatusPengajuan,
		TanggalKebutuhan: data.TanggalKebutuhan,
		Pencharian:       data.Pencharian,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
	}
}
func ListModeltoReqmanager(data []model.ReqGetManager) (datareq []request.ReqGetManager) {
	for _, val := range data {
		datareq = append(datareq, GetModelMnagerToReq(val))
	}
	return datareq
}
func GetReqMnagerToRes(data request.ReqGetManager) respon.ReSponGetManager {
	return respon.ReSponGetManager{
		IdPengajuan:      int(data.Id),
		NamaManager:      data.Nama,
		NamaDepartment:   data.NamaDepartment,
		Jumlah:           data.Jumlah,
		Alasan:           data.Alasan,
		StatusPengajuan:  data.StatusPengajuan,
		TanggalKebutuhan: data.TanggalKebutuhan,
		Pencaharian:      data.Pencharian,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
	}
}
func ListReqltoResmanager(data []request.ReqGetManager) (datares []respon.ReSponGetManager) {
	for _, val := range data {
		datares = append(datares, GetReqMnagerToRes(val))
	}
	return datares
}
