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
type Position struct {
	gorm.Model
	UserId      uint
	LevelKosong string
	Department  string
}
type Department struct {
	gorm.Model
	NamaDepartment string `gorm:"size:191" `
}
type SoalInterview struct {
	gorm.Model
	Kategori    string
	Description string
}

type Submission struct {
	gorm.Model
	IdDepartment      uint
	UserPengajuan     uint
	Jumlah            string
	Alasan            string
	TanggalKebutuhan  time.Time `gorm:"default:null"`
	KodePengajuan     string
	PosisiKosong      string
	Pencharian        string
	MaksimalGaji      float64
	StatusPengajuan   string
	IdVerifikasi      uint
	Idpersetujuan     uint
	IdEvaluasi        uint
	Golongan          string
	Durasi            string
	TanggalVerifikasi time.Time `gorm:"default:null"`
	TanggalPengajuan  time.Time `gorm:"default:null"`
	TanggalDisetujui  time.Time `gorm:"default:null"`
	TanggalEvaluasi   time.Time `gorm:"default:null"`
}
type GetUsersSubmission struct {
	Id                uint
	UserPengajuan     string
	NamaDepartment    string
	Alasan            string
	Pencharian        string
	KodePengajuan     string
	TanggalKebutuhan  string
	NamaEvaluasi      string
	NamaVerifikasi    string
	NamaPersetujuan   string
	StatusPengajuan   string
	TanggalVerifikasi string
	TanggalEvaluasi   string
	PosisiKosong      string
	TanggalPengajuan  time.Time
	TanggalDisetujui  string
}
type ReqGetManager struct {
	Id                uint
	Nama              string
	NamaDepartment    string
	Jumlah            string
	Alasan            string
	KodePengajuan     string
	StatusPengajuan   string
	PosisiKosong      string
	TanggalKebutuhan  string
	Pencharian        string
	Golongan          string
	TanggalPengajuan  time.Time
	TanggalVerifikasi string
	TanggalEvaluasi   string
	TanggalDisetujui  string
}
type ReqGetDireksi struct {
	Id               uint
	Nama             string
	NamaDepartment   string
	Jumlah           string
	Alasan           string
	KodePengajuan    string
	StatusPengajuan  string
	TanggalKebutuhan string
	PosisiKosong     string
	Pencharian       string
	Golongan         string
	TanggalPengajuan time.Time
	TanggalDisetujui string
}
type ReqGetAdmin struct {
	Id                uint
	UserPengajuan     string
	KodePengajuan     string
	PosisiKosong      string
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
	KodePengajuan     string
	Alasan            string
	StatusPengajuan   string
	TanggalKebutuhan  string
	PosisiKosong      string
	Pencharian        string
	Golongan          string
	TanggalPengajuan  time.Time
	TanggalVerifikasi string
}

type FormulirKandidat struct {
	gorm.Model
	NamaManager          string
	CuricullumVitae      string
	PosisiLamar          string
	KodePengajuan        string
	DepartementManager   string
	NamaKandidat         string
	ContactNumber        string
	ContactYangDihubungi string
	NomorContactDarurat  string
	InformasiJob         string
	NipRefrensi          string
	JenjangPendidikan    string
	NamaRefrensi         string
	Alamat               string
	Pengalaman           string
	AdminId              uint
}
type InterviewKandidat struct {
	gorm.Model
	NamaUser            string
	DepartementUser     string
	DepartementKandidat string
	KodePengajuan       string
	IdSoal              uint
	KategoriSoal        string
	NamaKandidat        string
	Nilai               float64
	Kriteria            string
	TanggalWwawancara   time.Time
	UserId              uint
	Role                string
}
type InterviewFPT struct {
	gorm.Model
	NamaUser            string
	DepartementUser     string
	DepartementKandidat string
	KodePengajuan       string
	IdSoal              uint
	KategoriSoal        string
	NamaKandidat        string
	Bobot               float64
	Nilai               float64
	TanggalWwawancara   time.Time
	UserId              uint
	Role                string
}
type DetailProses struct {
	gorm.Model
	IDAdmin               uint
	NilaiAdmin            float64
	NilaiManager          float64
	KandidatDepartment    string
	NamaKandidat          string
	TotalNilai            float64
	NilaiDireksi          float64
	NilaiDireksiFtp2      float64
	NilaiDireksiFtp3      float64
	NilaiInterviewDireksi float64
	NamaInterviewDireksi  string
	NamaDireksi           string
	NamaDireksi2          string
	NamaDireksi3          string
	InterviewDireksi      uint
	IdDireksi             uint
	IdDireksi2            uint
	IdDireksi3            uint
	KodePengajuan         string
	IdManager             uint
	NamaManager           string
	NamaAdmin             string
	Status                string
	CuricullumVitae       string
}
type SoalFPT struct {
	gorm.Model
	Kategori    string
	Description string
	Bobot       float64
}
