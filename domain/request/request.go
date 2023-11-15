package request

import (
	"time"
)

type RequestUser struct {
	Id        int
	Role      string `json:"roles" form:"roles" validate:"required,min=5,alpha"`
	Nip       string `json:"nip" form:"nip" validate:"required,min=5,numeric,max=15"`
	Password  string `json:"password" form:"password" validate:"required,min=5,alphanum,max=15"`
	Username  string `json:"username" form:"username" validate:"required,min=5,alphanum,max=15"`
	Name      string `json:"nama" form:"nama" validate:"required,min=5,alpha,max=15"`
	Bagian    string `json:"bagian" form:"bagian" validate:"required,min=5,alpha,max=15" `
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
	UserId    int
}

type RequestDepartment struct {
	Id             int
	NameDepartment string `json:"nama" form:"nama" validate:"required,min=5"`
	CreatedAt      time.Time
	DeletedAt      time.Time
	UpdateAt       time.Time
}
type ReqSubmissionManager struct {
	IdPengajuan      int
	IdDepartment     uint
	Jumlah           string `json:"jumlah" form:"jumlah" validate:"required"`
	Alasan           string `json:"alasan" form:"alasan" validate:"required,min=5"`
	KodePengajuan    string
	StatusPengajuan  string
	TanggalKebutuhan string `json:"tanggal_kebutuhan" form:"tanggal_kebutuhan" validate:"required"`
	Pencaharian      string `json:"pencaharian" form:"pencaharian" validate:"required"`
	Golongan         string `json:"golongan" form:"golongan" validate:"required"`
	NamaDepartment   string
	TanggalPengajuan time.Time
}
type ReqGetManager struct {
	Id               uint
	Nama             string
	NamaDepartment   string
	Jumlah           string `json:"jumlah" form:"jumlah" validate:"required"`
	Alasan           string `json:"alasan" form:"alasan" validate:"required,min=5"`
	StatusPengajuan  string
	TanggalKebutuhan string `json:"tanggal_kebutuhan" form:"tanggal_kebutuhan" validate:"required"`
	Pencharian       string `json:"pencaharian" form:"pencaharian" validate:"required"`
	Golongan         string `json:"golongan" form:"golongan" validate:"required"`
	KodePengajuan    string
	TanggalPengajuan time.Time
}
type ReqGetDireksi struct {
	Id               uint
	Nama             string
	NamaDepartment   string
	KodePengajuan    string
	Jumlah           string `json:"jumlah" form:"jumlah" validate:"required"`
	Alasan           string `json:"alasan" form:"alasan" validate:"required,min=5"`
	StatusPengajuan  string
	TanggalKebutuhan string `json:"tanggal_kebutuhan" form:"tanggal_kebutuhan" validate:"required"`
	Pencharian       string `json:"pencaharian" form:"pencaharian" validate:"required"`
	Golongan         string `json:"golongan" form:"golongan" validate:"required"`
	TanggalPengajuan time.Time
	TanggalDisetujui string `json:"tanggal_disetujui" form:"tanggal_disetujui" validate:"required"`
}
type ReqGetPresident struct {
	Id                uint
	Nama              string
	NamaDepartment    string
	KodePengajuan     string
	Jumlah            string `json:"jumlah" form:"jumlah" validate:"required"`
	Alasan            string `json:"alasan" form:"alasan" validate:"required,min=5"`
	StatusPengajuan   string
	TanggalKebutuhan  string `json:"tanggal_kebutuhan" form:"tanggal_kebutuhan" validate:"required"`
	Pencharian        string `json:"pencaharian" form:"pencaharian" validate:"required"`
	Golongan          string `json:"golongan" form:"golongan" validate:"required"`
	TanggalPengajuan  time.Time
	TanggalVerifikasi string `json:"tanggal_disetujui" form:"tanggal_disetujui" validate:"required"`
}
type ReqGetAdmin struct {
	Id                uint
	UserPengajuan     string
	NamaDepartment    string
	Jumlah            string
	Alasan            string
	Pencharian        string
	KodePengajuan     string
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
type UpdateAdmin struct {
	IdEvaluasi       int
	StatusPengajuan  string    `json:"status" form:"status" validate:"required,min=5"`
	TanggalDievalusi time.Time `json:"tanggal_disetujui" form:"tanggal_disetujui"`
	MaksimalGaji     float64   `json:"gaji" form:"gaji" validate:"required"`
}

type UpdateVicePresident struct {
	IdVerifikasi      int
	StatusPengajuan   string    `json:"status" form:"status" validate:"required,min=5"`
	TanggalVerifikasi time.Time `json:"tanggal_verifikasi" form:"tanggal_verifikasi"`
}
type UpdateDireksi struct {
	IdSetujui        int
	StatusPengajuan  string    `json:"status" form:"status" validate:"required,min=5"`
	TanggalDisetujui time.Time `json:"tanggal_disetujui" form:"tanggal_disetujui"`
}
type ReqFormulirKandidat struct {
	Id                   uint
	NamaManager          string `json:"nama_manager" form:"nama_manager" validate:"required,min=5"`
	KodePengajuan        string `json:"kodepengajuan" form:"kodepengajuan" validate:"required,min=5"`
	DepartementManager   string
	NamaKandidat         string `json:"nama_kandidat  " form:"nama_kandidat" validate:"required,min=5"`
	ContactNumber        string `json:"contact_kandidat" form:"contact_kandidat" validate:"required,min=5"`
	ContactYangDihubungi string `json:"contact_kerabat" form:"contact_kerabat" validate:"required"`
	NomorContactDarurat  string `json:"nomor_kerabat" form:"nomor_kerabat" validate:"required,min=5"`
	InformasiJob         string `json:"info_job" form:"info_job" validate:"required,min=5"`
	NipRefrensi          string `json:"nip_ref" form:"nip_ref"`
	JenjangPendidikan    string `json:"jenjang_pendidikan" form:"jenjang_pendidikan" validate:"required"`
	NamaRefrensi         string `json:"nama_refrensi" form:"nama_refrensi"`
	DepartmentRefrensi   string `json:"department_refrensi" form:"department_refrensi"`
	Alamat               string `json:"alamat" form:"alamat"`
	Pengalaman           string `json:"pengalaman" form:"pengalaman"`
	AdminId              uint
}
type RequesSoal struct {
	Id          uint
	Kategori    string `json:"kategori" form:"kategori" validate:"required,min=5"`
	Description string `json:"deskripsi" form:"deskripsi" validate:"required,min=5"`
}
type ReqInterviewKandidat struct {
	Id                  uint
	NamaUser            string  `json:"nama_user" form:"nama_user"`
	KodePengajuan       string  `json:"kodepengajuan" form:"kodepengajuan" validate:"required,min=5"`
	IdSoal              uint    `json:"id_soal" form:"id_soal"`
	KategoriSoal        string  `json:"kategori" form:"kategori" validate:"required,min=5"`
	DepartementUser     string  `json:"department_user" form:"department_manager"`
	NamaKandidat        string  `json:"nama_kandidat" form:"nama_kandidat"`
	Nilai               float64 `json:"nilai" form:"nilai"`
	Kriteria            string  `json:"kriteria" form:"kriteria" validate:"required,min=4"`
	TanggalWwawancara   string  `json:"tanggal" form:"tanggal"`
	UserId              uint
	DepartementKandidat string
	Role                string
}
