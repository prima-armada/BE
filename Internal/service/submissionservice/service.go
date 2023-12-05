package submissionservice

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/helper"
	"par/validasi"
	"time"

	"github.com/go-playground/validator"
)

type ServiceSubmission struct {
	rd       repocontract.RepoDepartment
	rsm      repocontract.RepoSubmission
	ru       repocontract.RepoUser
	validate *validator.Validate
}

func NewServiceSubmission(rsm repocontract.RepoSubmission, rd repocontract.RepoDepartment, ru repocontract.RepoUser) servicecontract.ServiceSubmission {
	return &ServiceSubmission{
		rd:       rd,
		rsm:      rsm,
		ru:       ru,
		validate: validator.New(),
	}
}

// GetAllSubmissionUser implements servicecontract.ServiceSubmission.
func (ssm *ServiceSubmission) GetAllSubmissionUser(deparment string) ([]request.ReqGetUsers, error) {
	cekdepatmen, errdeparment := ssm.rd.NameDepartment(deparment)

	if errdeparment != nil {
		return []request.ReqGetUsers{}, errors.New("deparment not found")
	}

	datarepo, errrepo := ssm.rsm.GetAllSubmissionUser(cekdepatmen.NameDepartment)
	if errrepo != nil {
		return []request.ReqGetUsers{}, errrepo
	}
	return datarepo, nil
}
func (ssm *ServiceSubmission) AddSubmission(newSubmission request.ReqSubmission, iduser int, res time.Time) (request.ReqSubmission, error) {

	validerr := ssm.validate.Struct(newSubmission)
	if validerr != nil {

		return request.ReqSubmission{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	cekuser, erruser := ssm.ru.IdUserExist(iduser)
	dtauser, erruser := ssm.ru.AllUser()

	if erruser != nil {
		return request.ReqSubmission{}, erruser
	}
	if cekuser.Role == "direksi" && newSubmission.PosisiKosong == "vicepresident" {
		for _, val := range dtauser {
			if val.Role == "vicepresident" && val.Bagian == cekuser.Bagian {
				return request.ReqSubmission{}, errors.New("anda sudah punya vicepresident")
			}
		}
	}
	if cekuser.Role == "vicepresident" && newSubmission.PosisiKosong == "manager" {
		for _, val := range dtauser {
			if val.Role == "manager" && val.Bagian == cekuser.Bagian {
				return request.ReqSubmission{}, errors.New("anda sudah punya manager")
			}
		}
	}
	if cekuser.Role == "vicepresident" && newSubmission.PosisiKosong == "staff" {
		for _, val := range dtauser {
			if val.Role == "manager" && val.Bagian == cekuser.Bagian {
				return request.ReqSubmission{}, errors.New("manager anda yang berhak mengajukan untuk staff")
			}
		}
	}

	if erruser != nil {
		return request.ReqSubmission{}, erruser
	}

	newSubmission.IdPengajuan = cekuser.Id

	cekdepartment, errdepartment := ssm.rd.NameDepartment(cekuser.Bagian)
	newSubmission.NamaDepartment = cekdepartment.NameDepartment

	if errdepartment != nil {
		return request.ReqSubmission{}, errdepartment
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)
	newSubmission.IdDepartment = uint(cekdepartment.Id)
	newSubmission.TanggalPengajuan = now
	newSubmission.StatusPengajuan = "diajukan"
	randString := helper.FileName(8)
	newSubmission.KodePengajuan = cekdepartment.NameDepartment + randString
	fmt.Print("newsubmission", newSubmission)
	datarepo, errrepo := ssm.rsm.AddSubmission(newSubmission, res)

	if errrepo != nil {
		return request.ReqSubmission{}, errrepo
	}
	return datarepo, nil
}

func (ssm *ServiceSubmission) GetAllSubmissionAdmin() ([]request.ReqGetAdmin, error) {
	datarepo, errrepo := ssm.rsm.GetAllSubmissionAdmin()

	if errrepo != nil {

		return []request.ReqGetAdmin{}, errrepo
	}
	return datarepo, nil
}

func (ssm *ServiceSubmission) UpdateSubmissionAdmin(iduser int, idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error) {
	validerr := ssm.validate.Struct(update)
	if validerr != nil {

		return request.UpdateAdmin{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	cekuser, erruser := ssm.ru.IdUserExist(iduser)

	if erruser != nil {
		return request.UpdateAdmin{}, erruser
	}

	update.IdEvaluasi = cekuser.Id

	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)
	if update.StatusPengajuan == "disetujui" || update.StatusPengajuan == "diverifikasi" {
		return request.UpdateAdmin{}, errors.New("anda tidak mempunyai akses tersebut")
	}
	update.TanggalDievalusi = now
	cekdata, errdata := ssm.rsm.GetAllSubmissionAdmin()

	if errdata != nil {
		return request.UpdateAdmin{}, errdata
	}
	for _, val := range cekdata {

		if uint(idsubmission) == val.Id && val.StatusPengajuan == "dievaluasi" {
			return request.UpdateAdmin{}, errors.New("pengajuan sudah di evaluasi")
		}

		if uint(idsubmission) == val.Id && val.StatusPengajuan == "diverifikasi" {
			return request.UpdateAdmin{}, errors.New("pengajuan sudah verifikasi")
		}
		if uint(idsubmission) == val.Id && val.StatusPengajuan == "disetujui" {
			return request.UpdateAdmin{}, errors.New("pengajuan sudah setujui oleh direksi")
		}
	}

	datarepo, errrepo := ssm.rsm.UpdateSubmissionAdmin(idsubmission, update)

	if errrepo != nil {
		return request.UpdateAdmin{}, errrepo
	}
	return datarepo, nil
}

func (ssm *ServiceSubmission) UpdateSubmissionPresident(iduser int, idsubmission int, update request.UpdateVicePresident) (request.UpdateVicePresident, error) {
	validerr := ssm.validate.Struct(update)
	if validerr != nil {

		return request.UpdateVicePresident{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	// fmt.Print("id manager", idManager)
	cekuser, erruser := ssm.ru.IdUserExist(iduser)
	fmt.Print(cekuser.Bagian)
	dep, errdep := ssm.rd.NameDepartment(cekuser.Bagian)

	if errdep != nil {
		return request.UpdateVicePresident{}, errdep
	}

	if dep.NameDepartment == "" {
		return request.UpdateVicePresident{}, errors.New("department tidak ada")
	}
	cekdata, errdata := ssm.rsm.GetAllSubmissionAdmin()

	if errdata != nil {
		return request.UpdateVicePresident{}, errdata
	}
	for _, val := range cekdata {
		if uint(idsubmission) == val.Id && cekuser.Bagian != val.NamaDepartment {
			return request.UpdateVicePresident{}, errors.New("department tidak sama")
		}
		if uint(idsubmission) == val.Id && val.StatusPengajuan == "diajukan" {
			return request.UpdateVicePresident{}, errors.New("masih di ajukan")
		}

		if uint(idsubmission) == val.Id && val.StatusPengajuan == "diverifikasi" {
			return request.UpdateVicePresident{}, errors.New("pengajuan sudah verifikasi")
		}
		if uint(idsubmission) == val.Id && val.StatusPengajuan == "disetujui" {
			return request.UpdateVicePresident{}, errors.New("pengajuan sudah setujui oleh direksi ")
		}
	}

	if erruser != nil {
		return request.UpdateVicePresident{}, erruser
	}

	update.IdVerifikasi = cekuser.Id

	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)
	if update.StatusPengajuan == "disetujui" || update.StatusPengajuan == "dievaluasi" {
		return request.UpdateVicePresident{}, errors.New("anda tidak mempunyai akses tersebut atau anda harus verifikasi pengajuan manager")
	}
	update.TanggalVerifikasi = now

	datarepo, errrepo := ssm.rsm.UpdateSubmissionPresident(idsubmission, update)

	if errrepo != nil {
		return request.UpdateVicePresident{}, errrepo
	}
	return datarepo, nil
}

// UpdateSubmissionDireksi implements servicecontract.ServiceSubmission.
func (ssm *ServiceSubmission) UpdateSubmissionDireksi(iduser int, idsubmission int, update request.UpdateDireksi) (request.UpdateDireksi, error) {
	validerr := ssm.validate.Struct(update)
	if validerr != nil {

		return request.UpdateDireksi{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	cekuser, erruser := ssm.ru.IdUserExist(iduser)
	dep, errdep := ssm.rd.NameDepartment(cekuser.Bagian)

	if errdep != nil {
		return request.UpdateDireksi{}, errdep
	}

	if dep.NameDepartment == "" {
		return request.UpdateDireksi{}, errors.New("department tidak ada")
	}
	cekdata, errdata := ssm.rsm.GetAllSubmissionAdmin()

	if errdata != nil {
		return request.UpdateDireksi{}, errdata
	}
	for _, val := range cekdata {
		if uint(idsubmission) == val.Id && cekuser.Bagian != val.NamaDepartment {
			return request.UpdateDireksi{}, errors.New("department tidak sama")
		}
		if uint(idsubmission) == val.Id && val.StatusPengajuan == "diajukan" {
			return request.UpdateDireksi{}, errors.New("masih di ajukan")
		}
		if uint(idsubmission) == val.Id && val.TanggalEvaluasi == "dievaluasi" {
			return request.UpdateDireksi{}, errors.New("pengajuan masih di evaluasi belum diverifikasi")
		}
		if uint(idsubmission) == val.Id && val.StatusPengajuan == "disetujui" {
			return request.UpdateDireksi{}, errors.New("pengajuan sudah di setujui")
		}
	}
	if erruser != nil {
		return request.UpdateDireksi{}, erruser
	}

	update.IdSetujui = cekuser.Id

	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)

	if update.StatusPengajuan == "diverifikasi" || update.StatusPengajuan == "dievaluasi" {
		return request.UpdateDireksi{}, errors.New(" anda hanya melakukan persutujuan pengajuan")
	}
	update.TanggalDisetujui = now

	datarepo, errrepo := ssm.rsm.UpdateSubmissionDireksi(idsubmission, update)

	if errrepo != nil {
		return request.UpdateDireksi{}, errrepo
	}
	return datarepo, nil
}

// GetNamaManager implements servicecontract.ServiceSubmission.
func (ssm *ServiceSubmission) GetNamaManager(namamanager string) ([]request.ReqGetManager, error) {
	datarepo, errrepo := ssm.rsm.GetNamaManager(namamanager)

	if errrepo != nil {
		return []request.ReqGetManager{}, errrepo
	}
	return datarepo, nil
}

// CodeSubmission implements servicecontract.ServiceSubmission.
func (ssm *ServiceSubmission) CodeSubmission(kode string) ([]request.ReqGetManager, error) {
	datarepo, errrepo := ssm.rsm.CodeSubmissions(kode)

	if errrepo != nil {
		return []request.ReqGetManager{}, errrepo
	}
	return datarepo, nil
}
