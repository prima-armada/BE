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
		Name:      data.Nama,
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
		Bagian:    data.Bagian,
		Nama:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}
func ListreqlUserToRes(data []request.RequestUser) (datareq []respon.ResponseUser) {
	for _, val := range data {

		datareq = append(datareq, ReqtoResponUser(val))
	}
	return datareq
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

func RequestsubmissionTomodel(data request.ReqSubmission) model.Submission {
	return model.Submission{
		IdDepartment:  data.IdDepartment,
		UserPengajuan: uint(data.IdPengajuan),
		Jumlah:        data.Jumlah,
		Alasan:        data.Alasan,
		// TanggalKebutuhan: tanggal,
		Pencharian:       data.Pencaharian,
		StatusPengajuan:  data.StatusPengajuan,
		PosisiKosong:     data.PosisiKosong,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
		KodePengajuan:    data.KodePengajuan,
	}

}
func ModelsubmissionToRequest(data model.Submission) request.ReqSubmission {
	return request.ReqSubmission{
		IdDepartment: data.IdDepartment,
		IdPengajuan:  int(data.UserPengajuan),
		Jumlah:       data.Jumlah,
		Alasan:       data.Alasan,
		// TanggalKebutuhan: tanggal,
		PosisiKosong:     data.PosisiKosong,
		Pencaharian:      data.Pencharian,
		StatusPengajuan:  data.StatusPengajuan,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
		KodePengajuan:    data.KodePengajuan,
	}

}
func ReqsubmisionToRespon(data request.ReqSubmission) respon.ResponSubmission {
	return respon.ResponSubmission{
		IdDepartment:     data.IdDepartment,
		IdPengajuan:      data.IdPengajuan,
		Jumlah:           data.Jumlah,
		Alasan:           data.Alasan,
		TanggalKebutuhan: data.TanggalKebutuhan,
		Pencaharian:      data.Pencaharian,
		StatusPengajuan:  data.StatusPengajuan,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
		KodePengajuan:    data.KodePengajuan,
	}

}

func GetModelMnagerToReq(data model.ReqGetManager) request.ReqGetManager {
	return request.ReqGetManager{
		Id:                data.Id,
		Nama:              data.Nama,
		NamaDepartment:    data.NamaDepartment,
		Jumlah:            data.Jumlah,
		Alasan:            data.Alasan,
		StatusPengajuan:   data.StatusPengajuan,
		TanggalKebutuhan:  data.TanggalKebutuhan,
		Pencharian:        data.Pencharian,
		Golongan:          data.Golongan,
		TanggalPengajuan:  data.TanggalPengajuan,
		KodePengajuan:     data.KodePengajuan,
		TanggalVerifikasi: data.TanggalVerifikasi,
		TanggalEvaluasi:   data.TanggalEvaluasi,
		TanggalDisetujui:  data.TanggalDisetujui,
	}
}
func GetusersubmissionToReq(data model.GetUsersSubmission) request.ReqGetUsers {
	return request.ReqGetUsers{
		Id:                data.Id,
		UserPengajuan:     data.UserPengajuan,
		NamaDepartment:    data.NamaDepartment,
		Alasan:            data.Alasan,
		Pencharian:        data.Pencharian,
		KodePengajuan:     data.KodePengajuan,
		TanggalKebutuhan:  data.TanggalKebutuhan,
		NamaEvaluasi:      data.NamaEvaluasi,
		NamaVerifikasi:    data.NamaVerifikasi,
		NamaPersetujuan:   data.NamaPersetujuan,
		StatusPengajuan:   data.StatusPengajuan,
		TanggalVerifikasi: data.TanggalVerifikasi,
		TanggalEvaluasi:   data.TanggalEvaluasi,
		PosisiKosong:      data.PosisiKosong,
		TanggalPengajuan:  data.TanggalPengajuan,
		TanggalDisetujui:  data.TanggalDisetujui,
	}
}
func ListModeltoReqsubmission(data []model.GetUsersSubmission) (datareq []request.ReqGetUsers) {
	for _, val := range data {
		datareq = append(datareq, GetusersubmissionToReq(val))
	}
	return datareq
}
func ListModeltoReqmanager(data []model.ReqGetManager) (datareq []request.ReqGetManager) {
	for _, val := range data {
		datareq = append(datareq, GetModelMnagerToReq(val))
	}
	return datareq
}
func GetReqMnagerToRes(data request.ReqGetManager) respon.ReSponGetManager {
	return respon.ReSponGetManager{
		IdPengajuan:       int(data.Id),
		NamaManager:       data.Nama,
		NamaDepartment:    data.NamaDepartment,
		Jumlah:            data.Jumlah,
		Alasan:            data.Alasan,
		StatusPengajuan:   data.StatusPengajuan,
		TanggalKebutuhan:  data.TanggalKebutuhan,
		Pencaharian:       data.Pencharian,
		Golongan:          data.Golongan,
		TanggalPengajuan:  data.TanggalPengajuan,
		KodePengajuan:     data.KodePengajuan,
		TanggalVerifikasi: data.TanggalVerifikasi,
		TanggalEvaluasi:   data.TanggalEvaluasi,
		TanggalDisetujui:  data.TanggalDisetujui,
	}
}
func ListReqltoResmanager(data []request.ReqGetManager) (datares []respon.ReSponGetManager) {
	for _, val := range data {
		datares = append(datares, GetReqMnagerToRes(val))
	}
	return datares
}
func GetReqDireksiToRes(data request.ReqGetDireksi) respon.ReSponGetDireksi {
	return respon.ReSponGetDireksi{
		IdPengajuan:      int(data.Id),
		NamaManager:      data.Nama,
		NamaDepartment:   data.NamaDepartment,
		Jumlah:           data.Jumlah,
		Alasan:           data.Alasan,
		StatusPengajuan:  data.StatusPengajuan,
		TanggalKebutuhan: data.TanggalKebutuhan,
		KodePengajuan:    data.KodePengajuan,
		Pencaharian:      data.Pencharian,
		Golongan:         data.Golongan,
		TanggalPengajuan: data.TanggalPengajuan,
		TanggalDisetujui: data.TanggalDisetujui,
	}
}
func ListReqltoResDireksi(data []request.ReqGetDireksi) (datares []respon.ReSponGetDireksi) {
	for _, val := range data {
		datares = append(datares, GetReqDireksiToRes(val))
	}
	return datares
}
func GetModeldireksiToReq(data model.ReqGetDireksi) request.ReqGetDireksi {
	return request.ReqGetDireksi{
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
		TanggalDisetujui: data.TanggalDisetujui,
		KodePengajuan:    data.KodePengajuan,
	}
}
func ListModeltoReqDireksi(data []model.ReqGetDireksi) (datareq []request.ReqGetDireksi) {
	for _, val := range data {
		datareq = append(datareq, GetModeldireksiToReq(val))
	}
	return datareq
}

func GetReqadminToRes(data request.ReqGetAdmin) respon.ReSponGetAdmin {
	return respon.ReSponGetAdmin{
		IdPengajuan:       int(data.Id),
		NamaPengajuan:     data.UserPengajuan,
		NamaDepartment:    data.NamaDepartment,
		Jumlah:            data.Jumlah,
		Alasan:            data.Alasan,
		Pencaharian:       data.Pencharian,
		TanggalKebutuhan:  data.TanggalKebutuhan,
		MaksimalGaji:      data.MaksimalGaji,
		NamaEvaluasi:      data.NamaEvaluasi,
		NamaVerifikasi:    data.NamaVerifikasi,
		NamaPersetujuan:   data.NamaPersetujuan,
		StatusPengajuan:   data.StatusPengajuan,
		Golongan:          data.Golongan,
		TanggalVerifikasi: data.TanggalVerifikasi,
		TanggalEvaluasi:   data.TanggalEvaluasi,
		TanggalPengajuan:  data.TanggalPengajuan,
		TanggalDisetujui:  data.TanggalDisetujui,
		KodePengajuan:     data.KodePengajuan,
		PosisiKosong:      data.PosisiKosong,
	}
}
func ListReqltoResAdmin(data []request.ReqGetAdmin) (datares []respon.ReSponGetAdmin) {
	for _, val := range data {
		datares = append(datares, GetReqadminToRes(val))
	}
	return datares
}
func GetModeladminToReq(data model.ReqGetAdmin) request.ReqGetAdmin {
	return request.ReqGetAdmin{
		Id:                data.Id,
		UserPengajuan:     data.UserPengajuan,
		NamaDepartment:    data.NamaDepartment,
		Jumlah:            data.Jumlah,
		Alasan:            data.Alasan,
		StatusPengajuan:   data.StatusPengajuan,
		Pencharian:        data.Pencharian,
		TanggalKebutuhan:  data.TanggalKebutuhan,
		MaksimalGaji:      data.MaksimalGaji,
		NamaEvaluasi:      data.NamaEvaluasi,
		NamaVerifikasi:    data.NamaVerifikasi,
		NamaPersetujuan:   data.NamaPersetujuan,
		Golongan:          data.Golongan,
		TanggalVerifikasi: data.TanggalVerifikasi,
		TanggalEvaluasi:   data.TanggalEvaluasi,
		TanggalPengajuan:  data.TanggalPengajuan,
		TanggalDisetujui:  data.TanggalDisetujui,
		KodePengajuan:     data.KodePengajuan,
		PosisiKosong:      data.PosisiKosong,
	}
}
func ListModeltoReqadmin(data []model.ReqGetAdmin) (datareq []request.ReqGetAdmin) {
	for _, val := range data {
		datareq = append(datareq, GetModeladminToReq(val))
	}
	return datareq
}
func ReqadminTomodelsubmissionudated(data request.UpdateAdmin) model.Submission {
	return model.Submission{

		MaksimalGaji:    data.MaksimalGaji,
		StatusPengajuan: data.StatusPengajuan,
		IdEvaluasi:      uint(data.IdEvaluasi),
		TanggalEvaluasi: data.TanggalDievalusi,
	}
}
func ModelsubmissionToReqadminudated(data model.Submission) request.UpdateAdmin {
	return request.UpdateAdmin{

		MaksimalGaji:     data.MaksimalGaji,
		StatusPengajuan:  data.StatusPengajuan,
		IdEvaluasi:       int(data.IdEvaluasi),
		TanggalDievalusi: data.TanggalEvaluasi,
	}
}
func ReqsubmissionToResadminupated(data request.UpdateAdmin) respon.ResponUpdateAdmin {
	return respon.ResponUpdateAdmin{

		MaksimalGaji:     data.MaksimalGaji,
		StatusPengajuan:  data.StatusPengajuan,
		IdEvaluasi:       int(data.IdEvaluasi),
		TanggalDievalusi: data.TanggalDievalusi,
	}
}
func ReqpresidentTomodelsubmissionudated(data request.UpdateVicePresident) model.Submission {
	return model.Submission{

		StatusPengajuan:   data.StatusPengajuan,
		IdVerifikasi:      uint(data.IdVerifikasi),
		TanggalVerifikasi: data.TanggalVerifikasi,
	}
}
func ModelsubmissionToReqpresidentupdated(data model.Submission) request.UpdateVicePresident {
	return request.UpdateVicePresident{

		StatusPengajuan:   data.StatusPengajuan,
		IdVerifikasi:      int(data.IdVerifikasi),
		TanggalVerifikasi: data.TanggalVerifikasi,
	}
}
func ReqtoResponPresident(data request.UpdateVicePresident) respon.ResponUpdateVicePresident {
	return respon.ResponUpdateVicePresident{

		StatusPengajuan:   data.StatusPengajuan,
		IdVerifikasi:      int(data.IdVerifikasi),
		TanggalVerifikasi: data.TanggalVerifikasi,
	}
}
func ReqdireksiTomodelsubmissionudated(data request.UpdateDireksi) model.Submission {
	return model.Submission{
		TanggalKebutuhan: data.TanggalKebutuhan,
		Durasi:           data.Durasi,
		StatusPengajuan:  data.StatusPengajuan,
		Idpersetujuan:    uint(data.IdSetujui),
		TanggalDisetujui: data.TanggalDisetujui,
	}
}
func ModelDireksiToreq(data model.Submission) request.UpdateDireksi {
	return request.UpdateDireksi{
		TanggalKebutuhan: data.TanggalKebutuhan,
		Durasi:           data.Durasi,
		StatusPengajuan:  data.StatusPengajuan,
		IdSetujui:        int(data.Idpersetujuan),
		TanggalDisetujui: data.TanggalDisetujui,
	}
}
func ReqDireksiTores(data request.UpdateDireksi) respon.ResponUpdateDireksi {
	return respon.ResponUpdateDireksi{
		TanggalKebutuhan:   data.TanggalKebutuhan,
		Durasi:             data.Durasi,
		StatusPengajuan:    data.StatusPengajuan,
		IdPersetujuan:      data.IdSetujui,
		TanggalPersetujuan: data.TanggalDisetujui,
	}
}

func GetmodelpresidentToReq(data model.ReqGetPresident) request.ReqGetPresident {
	return request.ReqGetPresident{
		Id:                data.Id,
		Nama:              data.Nama,
		NamaDepartment:    data.NamaDepartment,
		Jumlah:            data.Jumlah,
		Alasan:            data.Alasan,
		StatusPengajuan:   data.StatusPengajuan,
		TanggalKebutuhan:  data.TanggalKebutuhan,
		Pencharian:        data.Pencharian,
		Golongan:          data.Golongan,
		TanggalPengajuan:  data.TanggalPengajuan,
		TanggalVerifikasi: data.TanggalVerifikasi,
		KodePengajuan:     data.KodePengajuan,
	}
}
func ListmodelltoReqPresident(data []model.ReqGetPresident) (datares []request.ReqGetPresident) {
	for _, val := range data {
		datares = append(datares, GetmodelpresidentToReq(val))
	}
	return datares
}

func GetReqpresidentToRes(data request.ReqGetPresident) respon.ReSponGetPresident {
	return respon.ReSponGetPresident{
		IdPengajuan:       int(data.Id),
		NamaManager:       data.Nama,
		NamaDepartment:    data.NamaDepartment,
		Jumlah:            data.Jumlah,
		Alasan:            data.Alasan,
		StatusPengajuan:   data.StatusPengajuan,
		TanggalKebutuhan:  data.TanggalKebutuhan,
		Pencaharian:       data.Pencharian,
		Golongan:          data.Golongan,
		TanggalPengajuan:  data.TanggalPengajuan,
		Tanggalverifikasi: data.TanggalVerifikasi,
	}
}
func ListReqltoResPresident(data []request.ReqGetPresident) (datares []respon.ReSponGetPresident) {
	for _, val := range data {
		datares = append(datares, GetReqpresidentToRes(val))
	}
	return datares
}
func ReqtoResponKandidat(data request.ReqFormulirKandidat) respon.ResFormulirKandidat {
	return respon.ResFormulirKandidat{
		Id:                   data.Id,
		NamaManager:          data.NamaManager,
		KodePengajuan:        data.KodePengajuan,
		DepartementManager:   data.DepartementManager,
		NamaKandidat:         data.NamaKandidat,
		ContactNumber:        data.ContactNumber,
		ContactYangDihubungi: data.ContactYangDihubungi,
		NomorContactDarurat:  data.NomorContactDarurat,
		InformasiJob:         data.InformasiJob,
		NipRefrensi:          data.NipRefrensi,
		JenjangPendidikan:    data.JenjangPendidikan,
		NamaRefrensi:         data.NamaRefrensi,
		Alamat:               data.Alamat,
		Pengalaman:           data.Pengalaman,
		AdminId:              data.AdminId,
	}
}
func ReqtomodelKandidat(data request.ReqFormulirKandidat) model.FormulirKandidat {
	return model.FormulirKandidat{
		NamaManager:          data.NamaManager,
		KodePengajuan:        data.KodePengajuan,
		DepartementManager:   data.DepartementManager,
		NamaKandidat:         data.NamaKandidat,
		ContactNumber:        data.ContactNumber,
		ContactYangDihubungi: data.ContactYangDihubungi,
		NomorContactDarurat:  data.NomorContactDarurat,
		InformasiJob:         data.InformasiJob,
		NipRefrensi:          data.NipRefrensi,
		JenjangPendidikan:    data.JenjangPendidikan,
		NamaRefrensi:         data.NamaRefrensi,
		Alamat:               data.Alamat,
		Pengalaman:           data.Pengalaman,
		AdminId:              data.AdminId,
		CuricullumVitae:      data.CV,
		PosisiLamar:          data.PosisiLamar,
	}
}
func ModeltoReqKandidat(data *model.FormulirKandidat) request.ReqFormulirKandidat {
	return request.ReqFormulirKandidat{
		Id:                   data.ID,
		NamaManager:          data.NamaManager,
		KodePengajuan:        data.KodePengajuan,
		DepartementManager:   data.DepartementManager,
		NamaKandidat:         data.NamaKandidat,
		ContactNumber:        data.ContactNumber,
		ContactYangDihubungi: data.ContactYangDihubungi,
		NomorContactDarurat:  data.NomorContactDarurat,
		InformasiJob:         data.InformasiJob,
		NipRefrensi:          data.NipRefrensi,
		JenjangPendidikan:    data.JenjangPendidikan,
		NamaRefrensi:         data.NamaRefrensi,
		Alamat:               data.Alamat,
		Pengalaman:           data.Pengalaman,
		AdminId:              data.AdminId,
		CV:                   data.CuricullumVitae,
		PosisiLamar:          data.PosisiLamar,
	}
}
func ListKandidattoreq(data []model.FormulirKandidat) (datareq []request.ReqFormulirKandidat) {
	for _, val := range data {
		datareq = append(datareq, ModeltoReqKandidat(&val))
	}
	return datareq
}
func Listtoreqresponkandidat(data []request.ReqFormulirKandidat) (datares []respon.ResFormulirKandidat) {
	for _, val := range data {
		datares = append(datares, ReqtoResponKandidat(val))
	}
	return datares
}
func Reqsoaltorespon(data request.RequesSoal) respon.ResponSoal {
	return respon.ResponSoal{
		Id:          data.Id,
		Kategori:    data.Kategori,
		Description: data.Description,
	}
}
func ReqsoalTomodel(data request.RequesSoal) model.SoalInterview {
	return model.SoalInterview{

		Kategori:    data.Kategori,
		Description: data.Description,
	}
}
func Modelsoaltoreq(data *model.SoalInterview) request.RequesSoal {
	return request.RequesSoal{
		Id:          data.ID,
		Kategori:    data.Kategori,
		Description: data.Description,
	}
}

func ListReqDtoressoal(data []request.RequesSoal) (datares []respon.ResponSoal) {
	for _, val := range data {
		datares = append(datares, Reqsoaltorespon(val))
	}
	return datares
}

func Listmodelotreqsoal(data []model.SoalInterview) (datareq []request.RequesSoal) {
	for _, val := range data {
		datareq = append(datareq, Modelsoaltoreq(&val))
	}
	return datareq
}
func Reqinterviewtomodel(data request.ReqInterviewKandidat, tanggal time.Time) model.InterviewKandidat {
	return model.InterviewKandidat{

		NamaUser:            data.NamaUser,
		DepartementUser:     data.DepartementUser,
		DepartementKandidat: data.DepartementKandidat,
		KodePengajuan:       data.KodePengajuan,
		IdSoal:              data.IdSoal,
		KategoriSoal:        data.KategoriSoal,
		NamaKandidat:        data.NamaKandidat,
		Nilai:               data.Nilai,
		Kriteria:            data.Kriteria,
		TanggalWwawancara:   tanggal,
		UserId:              data.UserId,
		Role:                data.Role,
	}

}
func ModelinterviewToRequest(data *model.InterviewKandidat, tanggal string) request.ReqInterviewKandidat {
	return request.ReqInterviewKandidat{
		Id:                  data.ID,
		NamaUser:            data.NamaUser,
		DepartementUser:     data.DepartementUser,
		DepartementKandidat: data.DepartementKandidat,
		KodePengajuan:       data.KodePengajuan,
		IdSoal:              data.IdSoal,
		KategoriSoal:        data.KategoriSoal,
		NamaKandidat:        data.NamaKandidat,
		Nilai:               data.Nilai,
		Kriteria:            data.Kriteria,
		TanggalWwawancara:   tanggal,
		UserId:              data.UserId,
		Role:                data.Role,
	}

}
func ModelinterviewToRequest2(data *model.InterviewKandidat) request.ReqInterviewKandidat {
	return request.ReqInterviewKandidat{
		Id:                  data.ID,
		NamaUser:            data.NamaUser,
		DepartementUser:     data.DepartementUser,
		DepartementKandidat: data.DepartementKandidat,
		KodePengajuan:       data.KodePengajuan,
		IdSoal:              data.IdSoal,
		KategoriSoal:        data.KategoriSoal,
		NamaKandidat:        data.NamaKandidat,
		Nilai:               data.Nilai,
		UserId:              data.UserId,
	}

}
func ModelinterviewToReq(data *model.InterviewKandidat) request.ReqInterviewKandidat {
	return request.ReqInterviewKandidat{
		Id:                  data.ID,
		NamaUser:            data.NamaUser,
		DepartementUser:     data.DepartementUser,
		KategoriSoal:        data.KategoriSoal,
		DepartementKandidat: data.DepartementKandidat,
		KodePengajuan:       data.KodePengajuan,
		NamaKandidat:        data.NamaKandidat,
		Nilai:               data.Nilai,
		UserId:              data.UserId,
	}

}
func Listmodelotreqinterview(data []model.InterviewKandidat) (datareq []request.ReqInterviewKandidat) {
	for _, val := range data {
		datareq = append(datareq, ModelinterviewToReq(&val))
	}
	return datareq
}
func Reqprosesadmintomodel(data request.ReqDetailProsesAdmin) model.DetailProses {
	return model.DetailProses{

		IDAdmin:            data.IDAdmin,
		NilaiAdmin:         data.NilaiAdmin,
		NamaKandidat:       data.NamaKandidat,
		TotalNilai:         data.TotalNilai,
		KodePengajuan:      data.KodePengajuan,
		NamaAdmin:          data.NamaAdmin,
		Status:             data.Status,
		KandidatDepartment: data.KandidatDepartment,
	}

}
func Modelprosesadmintoreq(data *model.DetailProses) request.ReqDetailProsesAdmin {
	return request.ReqDetailProsesAdmin{
		Id:                 data.ID,
		IDAdmin:            data.IDAdmin,
		NilaiAdmin:         data.NilaiAdmin,
		NamaKandidat:       data.NamaKandidat,
		TotalNilai:         data.TotalNilai,
		KodePengajuan:      data.KodePengajuan,
		NamaAdmin:          data.NamaAdmin,
		Status:             data.Status,
		KandidatDepartment: data.KandidatDepartment,
	}

}
func GetallProsessforAdmin(data *model.DetailProses) request.ReqDetailProses {
	return request.ReqDetailProses{
		Id:                 data.ID,
		IDAdmin:            data.IDAdmin,
		NilaiAdmin:         data.NilaiAdmin,
		NamaKandidat:       data.NamaKandidat,
		TotalNilai:         data.TotalNilai,
		IdManager:          data.IdManager,
		KodePengajuan:      data.KodePengajuan,
		NamaAdmin:          data.NamaAdmin,
		Status:             data.Status,
		KandidatDepartment: data.KandidatDepartment,
		CuricullumVitae:    data.CuricullumVitae,
		NamaDireksi:        data.NamaDireksi,
		NilaiDireksi:       data.NilaiDireksi,
		NilaiManager:       data.NilaiManager,
		NamaManager:        data.NamaManager,
	}

}
func Listmodelotreqdetail(data []model.DetailProses) (datareq []request.ReqDetailProses) {
	for _, val := range data {
		datareq = append(datareq, GetallProsessforAdmin(&val))
	}
	return datareq
}
func Reqdetailmanager(data request.ReqDetailProsesManager) model.DetailProses {
	return model.DetailProses{

		NilaiManager:       data.NilaiManager,
		KandidatDepartment: data.KandidatDepartment,
		NamaKandidat:       data.NamaKandidat,
		TotalNilai:         data.TotalNilai,
		KodePengajuan:      data.KodePengajuan,
		IdManager:          data.IdManager,
		NamaManager:        data.NamaManager,
		Status:             data.Status,
	}
}
func Modeldetailmanagertoreq(data *model.DetailProses) request.ReqDetailProsesManager {
	return request.ReqDetailProsesManager{
		Id:                 data.ID,
		NilaiManager:       data.NilaiManager,
		KandidatDepartment: data.KandidatDepartment,
		NamaKandidat:       data.NamaKandidat,
		TotalNilai:         data.TotalNilai,
		KodePengajuan:      data.KodePengajuan,
		IdManager:          data.IdManager,
		NamaManager:        data.NamaManager,
		Status:             data.Status,
	}
}
func Listmodelotreqdetailmanager(data []model.DetailProses) (datareq []request.ReqDetailProsesManager) {
	for _, val := range data {
		datareq = append(datareq, Modeldetailmanagertoreq(&val))
	}
	return datareq
}
func Reqposisitomodel(data request.ReqPosisi) model.Position {
	return model.Position{

		UserId:      data.UserId,
		LevelKosong: data.LevelKosong,
		Department:  data.Department,
	}

}
func ModelPosisitoreq(data *model.Position) request.ReqPosisi {
	return request.ReqPosisi{
		Id:          data.ID,
		UserId:      data.UserId,
		LevelKosong: data.LevelKosong,
		Department:  data.Department,
	}

}
func ReqfptlTomodel(data request.RequesSoalFpt) model.SoalFPT {
	return model.SoalFPT{

		Kategori:    data.Kategori,
		Description: data.Description,
		Bobot:       data.Bobot,
	}
}
func ModelfptlToreq(data *model.SoalFPT) request.RequesSoalFpt {
	return request.RequesSoalFpt{
		Id:          data.ID,
		Kategori:    data.Kategori,
		Description: data.Description,
		Bobot:       data.Bobot,
	}
}
func Listmodelotreqfpt(data []model.SoalFPT) (datareq []request.RequesSoalFpt) {
	for _, val := range data {
		datareq = append(datareq, ModelfptlToreq(&val))
	}
	return datareq
}
func Reqinterviewfpttomodel(data request.ReqInterviewfpt, tanggal time.Time) model.InterviewFPT {
	return model.InterviewFPT{

		NamaUser:            data.NamaUser,
		DepartementUser:     data.DepartementUser,
		DepartementKandidat: data.DepartementKandidat,
		KodePengajuan:       data.KodePengajuan,
		IdSoal:              data.IdSoal,
		KategoriSoal:        data.KategoriSoal,
		NamaKandidat:        data.NamaKandidat,
		Nilai:               data.Nilai,
		TanggalWwawancara:   tanggal,
		UserId:              data.UserId,
		Bobot:               data.Bobot,
		Role:                data.Role,
	}

}
func Modelfpttoreq(data *model.InterviewFPT, tanggal string) request.ReqInterviewfpt {
	return request.ReqInterviewfpt{
		Id:                  data.ID,
		NamaUser:            data.NamaUser,
		DepartementUser:     data.DepartementUser,
		DepartementKandidat: data.DepartementKandidat,
		KodePengajuan:       data.KodePengajuan,
		IdSoal:              data.IdSoal,
		KategoriSoal:        data.KategoriSoal,
		NamaKandidat:        data.NamaKandidat,
		Nilai:               data.Nilai,
		TanggalWwawancara:   tanggal,
		Bobot:               data.Bobot,
		UserId:              data.UserId,
		Role:                data.Role,
	}

}
func Modelfpttoreqfpt(data *model.InterviewFPT) request.ReqInterviewfpt {
	return request.ReqInterviewfpt{
		Id:                  data.ID,
		NamaUser:            data.NamaUser,
		DepartementUser:     data.DepartementUser,
		DepartementKandidat: data.DepartementKandidat,
		KodePengajuan:       data.KodePengajuan,
		IdSoal:              data.IdSoal,
		KategoriSoal:        data.KategoriSoal,
		NamaKandidat:        data.NamaKandidat,
		Nilai:               data.Nilai,
		TanggalWwawancara:   data.TanggalWwawancara.String(),
		Bobot:               data.Bobot,
		UserId:              data.UserId,
		Role:                data.Role,
	}

}
func Listmodelotreqinterviewfpt(data []model.InterviewFPT) (datareq []request.ReqInterviewfpt) {
	for _, val := range data {
		datareq = append(datareq, Modelfpttoreqfpt(&val))
	}
	return datareq
}
func ReqdetailDireksi(data request.ReqDetailProsesDireksi) model.DetailProses {
	return model.DetailProses{

		NamaKandidat:  data.NamaKandidat,
		NilaiDireksi:  data.NilaiDireksi,
		KodePengajuan: data.KodePengajuan,
		IdDireksi:     data.IdDireksi,
		NamaDireksi:   data.NamaDireksi,
		Status:        data.Status,
	}
}
func ModeldetailDireksi(data *model.DetailProses) request.ReqDetailProsesDireksi {
	return request.ReqDetailProsesDireksi{

		NamaKandidat:  data.NamaKandidat,
		NilaiDireksi:  data.NilaiDireksi,
		KodePengajuan: data.KodePengajuan,
		IdDireksi:     data.IdDireksi,
		NamaDireksi:   data.NamaDireksi,
		Status:        data.Status,
	}
}
