package submissionservice

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
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

func (ssm *ServiceSubmission) AddSubmissionManager(newSubmission request.ReqSubmissionManager, idManager int, res time.Time) (request.ReqSubmissionManager, error) {
	validerr := ssm.validate.Struct(newSubmission)
	if validerr != nil {

		return request.ReqSubmissionManager{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	fmt.Print("id manager", idManager)
	cekuser, erruser := ssm.ru.IdUserExist(idManager)

	if erruser != nil {
		return request.ReqSubmissionManager{}, erruser
	}

	newSubmission.IdPengajuan = cekuser.Id
	fmt.Print("cekid", newSubmission.IdPengajuan)
	cekdepartment, errdepartment := ssm.rd.NameDepartment(cekuser.Bagian)

	if errdepartment != nil {
		return request.ReqSubmissionManager{}, errdepartment
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")

	//set timezone,
	now := time.Now().In(loc)
	newSubmission.IdDepartment = uint(cekdepartment.Id)
	newSubmission.TanggalPengajuan = now
	newSubmission.StatusPengajuan = "diajukan"

	fmt.Print("newsubmission", newSubmission)
	datarepo, errrepo := ssm.rsm.AddSubmissionManager(newSubmission, res)

	if errrepo != nil {
		return request.ReqSubmissionManager{}, errrepo
	}
	return datarepo, nil
}

func (ssm *ServiceSubmission) GetAllSubmissionManager(id int) ([]request.ReqGetManager, error) {
	datarepo, errrepo := ssm.rsm.GetAllSubmissionManager(id)
	// fmt.Print("service", datarepo)
	if errrepo != nil {
		return []request.ReqGetManager{}, errrepo
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

func (ssm *ServiceSubmission) GetAllSubmissionDireksi(deparment string) ([]request.ReqGetDireksi, error) {
	datarepo, errrepo := ssm.rsm.GetAllSubmissionDireksi(deparment)
	// fmt.Print("service", datarepo)
	if errrepo != nil {
		return []request.ReqGetDireksi{}, errrepo
	}
	return datarepo, nil
}

func (ssm *ServiceSubmission) UpdateSubmissionAdmin(iduser int, idsubmission int, update request.UpdateAdmin) (request.UpdateAdmin, error) {
	validerr := ssm.validate.Struct(update)
	if validerr != nil {

		return request.UpdateAdmin{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	// fmt.Print("id manager", idManager)
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
	// fmt.Print("id manager", idManager)
	cekuser, erruser := ssm.ru.IdUserExist(iduser)

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

// GetAllSubmissionPresident implements servicecontract.ServiceSubmission.
func (ssm *ServiceSubmission) GetAllSubmissionPresident(deparment string) ([]request.ReqGetPresident, error) {
	datarepo, errrepo := ssm.rsm.GetAllSubmissionPresident(deparment)
	// fmt.Print("service", datarepo)
	if errrepo != nil {
		return []request.ReqGetPresident{}, errrepo
	}
	return datarepo, nil
}
