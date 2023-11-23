package posisiservice

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/validasi"

	"github.com/go-playground/validator"
)

type ServicePosisi struct {
	ru       repocontract.RepoUser
	rs       repocontract.RepoSoal
	rp       repocontract.RepoPosisi
	validate *validator.Validate
}

func NewServiceposisi(rp repocontract.RepoPosisi, ru repocontract.RepoUser) servicecontract.ServicePosisi {
	return &ServicePosisi{
		rp: rp,
		ru: ru,

		validate: validator.New(),
	}
}

// AddPosisi implements servicecontract.ServicePosisi.
func (sp *ServicePosisi) AddPosisi(id int, newProcess request.ReqPosisi) (request.ReqPosisi, error) {
	validerr := sp.validate.Struct(newProcess)
	if validerr != nil {

		return request.ReqPosisi{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	repouser, erruser := sp.ru.IdUserExist(id)
	if erruser != nil {
		return request.ReqPosisi{}, erruser
	}

	getmanager, errmanager := sp.ru.GetAllManager(newProcess.LevelKosong)
	fmt.Print("get manager", getmanager)
	if errmanager != nil {
		return request.ReqPosisi{}, errmanager
	}
	if repouser.Role == "vicepresident" {
		for _, val := range getmanager {
			if repouser.Bagian == val.Bagian {
				fmt.Print("manager bagian", val.Bagian)
				return request.ReqPosisi{}, errors.New("anda sudah memiliki manager")
			}
		}
	}
	newProcess.UserId = uint(repouser.Id)
	newProcess.Department = repouser.Bagian
	cekposisi, errposisi := sp.rp.Getdetailposisi(int(newProcess.UserId))

	if errposisi != nil {
		return request.ReqPosisi{}, errposisi
	}

	if cekposisi.LevelKosong == newProcess.LevelKosong && cekposisi.Department == newProcess.Department && cekposisi.UserId == newProcess.UserId {
		return request.ReqPosisi{}, errors.New("anda sudah membuat posisi kosong")
	}
	datarepo, errrepo := sp.rp.AddPosisi(newProcess)

	if errrepo != nil {
		return request.ReqPosisi{}, errrepo
	}

	return datarepo, nil

}
