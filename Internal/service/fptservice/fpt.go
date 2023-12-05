package fptservice

import (
	"errors"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/validasi"

	"github.com/go-playground/validator"
)

type Servicefpt struct {
	rsl      repocontract.RepoSoalFpt
	validate *validator.Validate
}

func NewServicefpt(rsl repocontract.RepoSoalFpt) servicecontract.ServiceSoalFpt {
	return &Servicefpt{
		rsl:      rsl,
		validate: validator.New(),
	}
}

func (sf *Servicefpt) AddSoal(newsoal request.RequesSoalFpt) (request.RequesSoalFpt, error) {
	validerr := sf.validate.Struct(newsoal)
	if validerr != nil {

		return request.RequesSoalFpt{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	kategoriExist, _ := sf.rsl.KategoriSoal(newsoal.Kategori)
	if kategoriExist.Kategori == newsoal.Kategori {
		return request.RequesSoalFpt{}, errors.New("kategori sudah ada")
	}
	datarepo, errrepo := sf.rsl.AddSoal(newsoal)

	if errrepo != nil {
		return request.RequesSoalFpt{}, errrepo
	}
	return datarepo, nil
}

// AllSoal implements servicecontract.ServiceSoalFpt.
func (sf *Servicefpt) AllSoal() (data []request.RequesSoalFpt, err error) {
	datarepo, errrepo := sf.rsl.AllSoal()
	if errrepo != nil {
		return []request.RequesSoalFpt{}, errrepo
	}
	return datarepo, nil
}

func (sf *Servicefpt) DeleteSoal(id int) error {
	// fmt.Print(id)
	_, err := sf.rsl.DeletedSoal(id)

	if err != nil {
		return err
	}
	return nil
}

// Updatedsoal implements servicecontract.ServiceSoalFpt.
func (sf *Servicefpt) Updatedsoal(id int, update request.RequesSoalFpt) (data request.RequesSoalFpt, err error) {
	validerr := sf.validate.Struct(update)
	if validerr != nil {

		return request.RequesSoalFpt{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	data, err = sf.rsl.Updatedsoal(id, update)

	if err != nil {
		return data, err
	}
	return data, nil
}

// KategoriSoal implements servicecontract.ServiceSoalFpt.
func (sf *Servicefpt) KategoriSoal(kategori string) (data request.RequesSoalFpt, err error) {

	data, err = sf.rsl.KategoriSoal(kategori)

	if err != nil {
		return data, err
	}
	return data, nil
}
