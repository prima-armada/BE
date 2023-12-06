package soalservice

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/validasi"

	"github.com/go-playground/validator"
)

type Servicessoal struct {
	rsl      repocontract.RepoSoal
	validate *validator.Validate
}

func NewServiceSoal(rsl repocontract.RepoSoal) servicecontract.ServiceSoal {
	return &Servicessoal{
		rsl:      rsl,
		validate: validator.New(),
	}
}

func (ssl *Servicessoal) AddSoal(newksoal request.RequesSoal) (request.RequesSoal, error) {
	validerr := ssl.validate.Struct(newksoal)
	if validerr != nil {

		return request.RequesSoal{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	// Mengecek apakah kategori sudah ada
	kategoriExist, _ := ssl.rsl.KategoriSoal(newksoal.Kategori)
	if kategoriExist.Kategori == newksoal.Kategori {
		return request.RequesSoal{}, errors.New("kategori sudah ada")
	}
	datarepo, errrepo := ssl.rsl.AddSoal(newksoal)

	if errrepo != nil {
		return request.RequesSoal{}, errrepo
	}
	return datarepo, nil
}

// AllSoal implements servicecontract.ServiceSoal.
func (ssl *Servicessoal) AllSoal() (data []request.RequesSoal, err error) {
	datarepo, errepo := ssl.rsl.AllSoal()
	// fmt.Print(datarepo, "soal")
	if datarepo == nil {
		return []request.RequesSoal{}, errors.New("data tidak ada")
	}
	if errepo != nil {
		return datarepo, err
	}
	return datarepo, nil
}

// KategoriSoal implements servicecontract.ServiceSoal.
func (ssl *Servicessoal) KategoriSoal(kategori string) (data request.RequesSoal, err error) {

	data, err = ssl.rsl.KategoriSoal(kategori)

	if err != nil {
		return data, err
	}
	return data, nil
}

// Updatedsoal implements servicecontract.ServiceSoal.
func (ssl *Servicessoal) Updatedsoal(id int, update request.RequesSoal) (data request.RequesSoal, err error) {
	validerr := ssl.validate.Struct(update)
	if validerr != nil {

		return request.RequesSoal{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	data, err = ssl.rsl.Updatedsoal(id, update)

	if err != nil {
		return data, err
	}
	return data, nil
}

// DeleteSoal implements servicecontract.ServiceSoal.
func (ssl *Servicessoal) DeleteSoal(id int) error {

	fmt.Print(id)
	_, err := ssl.rsl.DeletedSoal(id)

	if err != nil {
		return err
	}
	return nil
}
