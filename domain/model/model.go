package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role     string
	Nip      string
	Password string
	Username string
	Nama     string
	Bagian   string
}

type Department struct {
	gorm.Model
	NamaDepartment string `gorm:"size:191" `
}
type Submission struct {
	gorm.Model
	IdDepartment      uint
	UserPengajuan     uint
	Jumlah            string
	Alasan            string
	TanggalKebutuhan  time.Time `gorm:"default:null"`
	Pencharian        string
	MaksimalGaji      float64
	StatusPengajuan   string
	IdVerifikasi      uint
	Idpersetujuan     uint
	IdEvaluasi        uint
	Golongan          string
	TanggalVerifikasi time.Time `gorm:"default:null"`
	TanggalPengajuan  time.Time `gorm:"default:null"`
	TanggalDisetujui  time.Time `gorm:"default:null"`
	TanggalEvaluasi   time.Time `gorm:"default:null"`
}
type ReqGetManager struct {
	Id               uint
	Nama             string
	NamaDepartment   string
	Jumlah           string
	Alasan           string
	StatusPengajuan  string
	TanggalKebutuhan string
	Pencharian       string
	Golongan         string
	TanggalPengajuan time.Time
}
type ReqGetDireksi struct {
	Id               uint
	Nama             string
	NamaDepartment   string
	Jumlah           string
	Alasan           string
	StatusPengajuan  string
	TanggalKebutuhan string
	Pencharian       string
	Golongan         string
	TanggalPengajuan time.Time
	TanggalDisetujui string
}
type ReqGetAdmin struct {
	Id                uint
	UserPengajuan     string
	NamaDepartment    string
	Jumlah            string
	Alasan            string
	Pencharian        string
	TanggalKebutuhan  string
	MaksimalGaji      float64
	NamaEvaluasi      string
	NamaVerifikasi    string
	NamaPersetujuan   string
	StatusPengajuan   string
	Golongan          string
	TanggalVerifikasi string
	TanggalEvaluasi   string
	TanggalPengajuan  time.Time
	TanggalDisetujui  string
}
type ReqGetPresident struct {
	Id                uint
	Nama              string
	NamaDepartment    string
	Jumlah            string
	Alasan            string
	StatusPengajuan   string
	TanggalKebutuhan  string
	Pencharian        string
	Golongan          string
	TanggalPengajuan  time.Time
	TanggalVerifikasi string
}
