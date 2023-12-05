package interviewservice

import (
	"errors"
	"fmt"
	"par/domain/contract/repocontract"
	"par/domain/contract/servicecontract"
	"par/domain/request"
	"par/kriteria"
	"par/validasi"
	"time"

	"github.com/go-playground/validator"
)

type Serviceinterview struct {
	ri       repocontract.RepoInterview
	rd       repocontract.RepoDepartment
	rk       repocontract.RepoKandidat
	ru       repocontract.RepoUser
	rs       repocontract.RepoSoal
	rsm      repocontract.RepoSubmission
	rdt      repocontract.RepoProcess
	rft      repocontract.RepoSoalFpt
	validate *validator.Validate
}

func NewServiceinterview(ri repocontract.RepoInterview, rk repocontract.RepoKandidat, rd repocontract.RepoDepartment, ru repocontract.RepoUser, rs repocontract.RepoSoal, rsm repocontract.RepoSubmission, rdt repocontract.RepoProcess, rft repocontract.RepoSoalFpt) servicecontract.Serviceinterview {
	return &Serviceinterview{
		ri:       ri,
		rd:       rd,
		rk:       rk,
		ru:       ru,
		rs:       rs,
		rsm:      rsm,
		rdt:      rdt,
		rft:      rft,
		validate: validator.New(),
	}
}

// AddInterview implements servicecontract.Serviceinterview.
func (si *Serviceinterview) AddInterview(newinterview request.ReqInterviewKandidat) (request.ReqInterviewKandidat, error) {
	validerr := si.validate.Struct(newinterview)
	if validerr != nil {

		return request.ReqInterviewKandidat{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}

	repouser, erruser := si.ru.IdUserExist(int(newinterview.UserId))

	if erruser != nil {
		return request.ReqInterviewKandidat{}, erruser
	}
	newinterview.NamaUser = repouser.Name

	newinterview.DepartementUser = repouser.Bagian
	newinterview.Role = repouser.Role

	newinterview.UserId = uint(repouser.Id)

	kriteriakandidat := kriteria.KriteriaKandidat(newinterview.Kriteria)

	newinterview.Nilai = float64(kriteriakandidat)

	getcodedannama, errcode := si.rk.GetCodedannamaKandidat(newinterview.KodePengajuan, newinterview.NamaKandidat)

	if errcode != nil {
		return request.ReqInterviewKandidat{}, errcode
	}
	newinterview.NamaKandidat = getcodedannama.NamaKandidat
	newinterview.DepartementKandidat = getcodedannama.DepartementManager

	datapengajuan, errPengajuan := si.rsm.GetAllSubmissionUser(repouser.Bagian)
	getuser, errusers := si.ru.AllUser()

	if errusers != nil {
		return request.ReqInterviewKandidat{}, erruser
	}
	if errPengajuan != nil {
		return request.ReqInterviewKandidat{}, errPengajuan
	}
	if newinterview.Role == "vicepresident" {
		for _, val := range datapengajuan {
			if newinterview.KodePengajuan == val.KodePengajuan {
				if val.PosisiKosong == "staff" {
					for _, val := range getuser {
						if val.Role == "manager" && newinterview.DepartementUser == val.Bagian {
							return request.ReqInterviewKandidat{}, errors.New("anda sudah punya manager,harap manager anda yang isi	")
						}
					}
				}
				if val.PosisiKosong == "manager" {
					for _, val := range getuser {
						if val.Role == "manager" && newinterview.DepartementUser == val.Bagian {
							return request.ReqInterviewKandidat{}, errors.New("anda sudah punya manager")
						}
					}
				}
			}
		}
	}
	if newinterview.Role == "direksi" {
		for _, val := range datapengajuan {
			if newinterview.KodePengajuan == val.KodePengajuan {
				if val.PosisiKosong == "manager" {
					for _, val := range getuser {
						if val.Role == "vicepresident" && newinterview.DepartementUser == val.Bagian {
							return request.ReqInterviewKandidat{}, errors.New("anda sudah punya manager,harap vicepresident anda yang isi")
						}
					}
				}
				if val.PosisiKosong == "vicepresident" {
					for _, val := range getuser {
						if val.Role == "vicepresident" && newinterview.DepartementUser == val.Bagian {
							return request.ReqInterviewKandidat{}, errors.New("anda sudah punya vicepresident")
						}
					}
				}
			}
		}
	}
	for _, val := range datapengajuan {
		if newinterview.KodePengajuan == val.KodePengajuan {
			if val.StatusPengajuan != "disetujui" {
				return request.ReqInterviewKandidat{}, errors.New(" kode pengajuan " + newinterview.KodePengajuan + " belum disetujui ")
			}
		}
	}
	dataproses, errproses := si.rdt.GetallDetail()

	if errproses != nil {
		return request.ReqInterviewKandidat{}, errproses
	}
	if repouser.Role == "manager" || repouser.Role == "vicepresident" || repouser.Role == "direksi" {
		for _, val := range dataproses {
			if newinterview.KodePengajuan == val.KodePengajuan {
				if val.Status != "lolos ke tahap interview user" && newinterview.NamaKandidat == val.NamaKandidat {
					return request.ReqInterviewKandidat{}, errors.New(" maaf kode pengajuan " + newinterview.KodePengajuan + val.Status)
				}
			}
		}
	}
	if repouser.Role == "manager" || repouser.Role == "vicepresident" || repouser.Role == "direksi" {
		if repouser.Bagian != newinterview.DepartementKandidat {
			return request.ReqInterviewKandidat{}, errors.New("hanya boleh interview sesama department")
		}
	}

	res, errConvtime := time.Parse("02/01/2006", newinterview.TanggalWwawancara)
	if errConvtime != nil {
		return request.ReqInterviewKandidat{}, errConvtime
	}
	soal, errsoal := si.rs.KategoriSoal(newinterview.KategoriSoal)

	if errsoal != nil {
		return request.ReqInterviewKandidat{}, errsoal
	}
	if soal.Kategori == "" {
		return request.ReqInterviewKandidat{}, errors.New("kategori tidak ada")
	}
	cekinterview, errcek := si.ri.CekallInterview(int(newinterview.UserId), newinterview.KodePengajuan, newinterview.NamaKandidat)
	// fmt.Print(cekinterview)
	for _, val := range cekinterview {
		if newinterview.KodePengajuan == val.KodePengajuan {
			if newinterview.NamaKandidat == val.NamaKandidat {
				if newinterview.UserId == val.UserId && soal.Kategori == val.KategoriSoal {
					return request.ReqInterviewKandidat{}, errors.New("kategori soal tidak boleh sama ")
				}
			}
		}
	}

	if errcek != nil {
		return request.ReqInterviewKandidat{}, errcek
	}

	newinterview.IdSoal = soal.Id

	datarepo, errrepo := si.ri.AddInterview(newinterview, res)

	if errrepo != nil {
		return request.ReqInterviewKandidat{}, errrepo
	}
	return datarepo, nil
}

func (si *Serviceinterview) AddInterviewfpt(newinterview request.ReqInterviewfpt) (request.ReqInterviewfpt, error) {
	validerr := si.validate.Struct(newinterview)
	if validerr != nil {

		return request.ReqInterviewfpt{}, errors.New(validasi.ValidationErrorHandle(validerr))
	}
	if newinterview.Nilai < 10 && newinterview.Nilai > 100 {
		return request.ReqInterviewfpt{}, errors.New("hanya 10 sampai 100")
	}

	repouser, erruser := si.ru.IdUserExist(int(newinterview.UserId))

	if erruser != nil {
		return request.ReqInterviewfpt{}, erruser
	}
	newinterview.NamaUser = repouser.Name

	newinterview.DepartementUser = repouser.Bagian
	newinterview.Role = repouser.Role

	newinterview.UserId = uint(repouser.Id)
	ceksoal, errsoal := si.rft.KategoriSoal(newinterview.KategoriSoal)
	fmt.Print("ini bobot", ceksoal.Bobot)
	if errsoal != nil {
		return request.ReqInterviewfpt{}, errsoal
	}
	newinterview.IdSoal = ceksoal.Id
	newinterview.Bobot = ceksoal.Bobot

	var nilai float64

	nilai = (ceksoal.Bobot / 100) * newinterview.Nilai

	newinterview.Nilai = nilai
	getcodedannama, errcode := si.rk.GetCodedannamaKandidat(newinterview.KodePengajuan, newinterview.NamaKandidat)

	if errcode != nil {
		return request.ReqInterviewfpt{}, errcode
	}
	newinterview.NamaKandidat = getcodedannama.NamaKandidat
	newinterview.DepartementKandidat = getcodedannama.DepartementManager
	fmt.Print("tes", newinterview.DepartementKandidat)
	dataproses, errproses := si.rdt.GetallDetail()
	fmt.Print("ini data proses", dataproses)
	if errproses != nil {
		return request.ReqInterviewfpt{}, errproses
	}
	if repouser.Role != "direksi" {
		return request.ReqInterviewfpt{}, errors.New("hanya untuk direksi")
	}
	if repouser.Role == "direksi" {
		for _, val := range dataproses {
			if newinterview.KodePengajuan == val.KodePengajuan {

				if val.Status != "anda lolos ke tahap fpt" && newinterview.NamaKandidat == val.NamaKandidat {
					return request.ReqInterviewfpt{}, errors.New(" maaf kode pengajuan " + newinterview.KodePengajuan + val.Status)
				}
			}
		}
	}
	cekkandidat, errkandidat := si.rk.GetCodeKandidat(newinterview.KodePengajuan)
	if errkandidat != nil {
		return request.ReqInterviewfpt{}, errkandidat
	}
	for _, val := range cekkandidat {
		if newinterview.KodePengajuan == val.KodePengajuan && newinterview.NamaKandidat == val.NamaKandidat {
			if val.PosisiLamar == "staff" {
				return request.ReqInterviewfpt{}, errors.New("maaf fpt hanya untuk proses manager dan vicepresident")
			}
		}
	}
	res, errConvtime := time.Parse("02/01/2006", newinterview.TanggalWwawancara)
	if errConvtime != nil {
		return request.ReqInterviewfpt{}, errConvtime
	}
	cekinterview, errinterview := si.ri.GetallInterviewftp(newinterview.NamaKandidat, newinterview.KodePengajuan)
	if errinterview != nil {
		return request.ReqInterviewfpt{}, errinterview
	}

	for _, val := range cekinterview {
		if newinterview.KodePengajuan == val.KodePengajuan && newinterview.KategoriSoal == val.KategoriSoal && newinterview.NamaKandidat == val.NamaKandidat {
			return request.ReqInterviewfpt{}, errors.New("kategori tidak boleh sama")
		}
	}
	datarepo, errrepo := si.ri.AddInterviewfpt(newinterview, res)

	if errrepo != nil {
		return request.ReqInterviewfpt{}, errrepo
	}
	return datarepo, nil
}
func (si *Serviceinterview) GetallInterview(userid int, kode string, nama string) (data []request.ReqInterviewKandidat, err error) {
	data, err = si.ri.GetallInterview(userid, kode, nama)

	if err != nil {
		return data, err
	}
	return data, nil
}
