package kriteria

import (
	"errors"
	"par/domain/request"
	"time"
)

func KriteriaKandidat(test string) (nilai int) {
	if test == "sangat baik" {
		nilai = 4
		return nilai
	} else if test == "baik" {
		nilai = 3
		return nilai
	} else if test == "cukup" {
		nilai = 2
		return nilai
	} else if test == "kurang" {
		nilai = 1
		return nilai
	}
	nilai = 0
	return nilai
}
func CekSTATUS(Nilai float64) (status string, err error) {
	if Nilai > float64(4) {
		status = "nilai lebih dari 4 tidak ada"
		return status, errors.New("lebih dari 4 tidak ada")
	} else if Nilai >= float64(3) && Nilai <= 4 {
		status = "lolos ke tahap interview user"

		return status, nil
	} else if Nilai > float64(2) && Nilai < float64(3) {
		status = "tidak lolos ke tahap interview user"
		return status, nil
	}

	status = "anda tidak  lolos ke tahap interview user"
	return status, nil
}
func CekSTATUSformanager(Nilai float64) (status string, err error) {
	if Nilai > float64(4) {
		status = "nilai lebih dari 4 tidak ada"
		return status, errors.New("lebih dari 4 tidak ada")
	} else if Nilai >= float64(3) && Nilai <= 4 {
		status = "harap menunggu konfirmasi hr"

		return status, nil
	} else if Nilai >= float64(2) && Nilai < float64(3) {
		status = "anda tidak  lolos ke tahap selanjutnya"
		return status, nil
	}

	status = "anda tidak  lolos ke tahap selanjutnya"
	return status, nil
}
func CekSTATUSfordireksi(Nilai float64) (status string, err error) {
	if Nilai > 100 {
		return "", errors.New("nilai lebih dari 100 salah input dalam penilaiai ftp")
	} else if Nilai >= 70 && Nilai <= 100 {
		return "harap menunggu konfirmasi hr", nil
	}
	return "anda tidak lolos", nil
}

func Timer(data []request.ReqGetAdmin, tanggalnow time.Time, id uint) (tanggal time.Time, durasi time.Duration) {
	for _, val := range data {
		if uint(id) == val.Id {
			if val.PosisiKosong == "staff" {
				tanggal = tanggalnow.AddDate(0, 0, 30)
				durasi = tanggal.Sub(tanggalnow)
				return tanggal, durasi

			} else if val.PosisiKosong == "spv" {
				tanggal = tanggalnow.AddDate(0, 0, 60)
				durasi = tanggal.Sub(tanggalnow)
				return tanggal, durasi

			} else if val.PosisiKosong == "assisten manager" {
				tanggal = tanggalnow.AddDate(0, 0, 90)
				durasi = tanggal.Sub(tanggalnow)
				return tanggal, durasi

			} else if val.PosisiKosong == "manager" {
				tanggal = tanggalnow.AddDate(0, 0, 120)
				durasi = tanggal.Sub(tanggalnow)
				return tanggal, durasi

			} else if val.PosisiKosong == "vicepresident" {
				tanggal = tanggalnow.AddDate(0, 0, 150)
				durasi = tanggal.Sub(tanggalnow)
				return tanggal, durasi

			}
		}
	}
	return tanggal, durasi
}
