package request

import "time"

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
	StatusPengajuan  string
	TanggalKebutuhan string `json:"tanggal_kebutuhan" form:"tanggal_kebutuhan" validate:"required"`
	Pencaharian      string `json:"pencaharian" form:"pencaharian" validate:"required"`
	Golongan         string `json:"golongan" form:"golongan" validate:"required"`
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
	TanggalPengajuan time.Time
}
type ReqGetDireksi struct {
	Id               uint
	Nama             string
	NamaDepartment   string
	Jumlah           string `json:"jumlah" form:"jumlah" validate:"required"`
	Alasan           string `json:"alasan" form:"alasan" validate:"required,min=5"`
	StatusPengajuan  string
	TanggalKebutuhan string `json:"tanggal_kebutuhan" form:"tanggal_kebutuhan" validate:"required"`
	Pencharian       string `json:"pencaharian" form:"pencaharian" validate:"required"`
	Golongan         string `json:"golongan" form:"golongan" validate:"required"`
	TanggalPengajuan time.Time
	TanggalDisetujui string `json:"tanggal_disetujui" form:"tanggal_disetujui" validate:"required"`
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
type UpdateAdmin struct {
	IdEvaluasi       int
	StatusPengajuan  string    `json:"status" form:"status" validate:"required,min=5"`
	TanggalDievalusi time.Time `json:"tanggal_disetujui" form:"tanggal_disetujui"`
	MaksimalGaji     float64   `json:"gaji" form:"gaji" validate:"required"`
}
