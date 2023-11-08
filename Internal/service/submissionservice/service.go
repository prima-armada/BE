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
