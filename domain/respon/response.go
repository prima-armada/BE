package respon

import "time"

type ResponseUser struct {
	Role      string    `json:"role"`
	Nip       string    `json:"nip"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdat"`
}
type LoginRespon struct {
	Role     string `json:"role"`
	Nip      string `json:"nip"`
	Username string `json:"username"`
	Bagian   string `json:"bagian"`
	Token    string `json:"token"`
}
type ResponseDeparment struct {
	Id             int       `json:"id_departments"`
	NameDepartment string    `json:"nama_departments"`
	CreatedAt      time.Time `json:"createdat"`
	UpdateAt       time.Time `json:"updatedat"`
}

type ResponSubmissionManager struct {
	IdPengajuan      int       `json:"user"`
	IdDepartment     uint      `json:"department"`
	Jumlah           string    `json:"jumlah"`
	Alasan           string    `json:"alasan"`
	StatusPengajuan  string    `json:"status"`
	TanggalKebutuhan string    `json:"tanggal_kebutuhan"`
	Pencaharian      string    `json:"pencaharian"`
	Golongan         string    `json:"golongan"`
	TanggalPengajuan time.Time `json:"pengajuan_tanggal"`
}
type ReSponGetManager struct {
	IdPengajuan      int       `json:"id"`
	NamaManager      string    `json:"user"`
	NamaDepartment   string    `json:"department"`
	Jumlah           string    `json:"jumlah"`
	Alasan           string    `json:"alasan"`
	StatusPengajuan  string    `json:"status"`
	TanggalKebutuhan string    `json:"tanggal_kebutuhan"`
	Pencaharian      string    `json:"pencarian"`
	Golongan         string    `json:"golongan"`
	TanggalPengajuan time.Time `json:"tanggal_pengajuan"`
}
type ReSponGetDireksi struct {
	IdPengajuan      int       `json:"id"`
	NamaManager      string    `json:"manager"`
	NamaDepartment   string    `json:"department"`
	Jumlah           string    `json:"jumlah"`
	Alasan           string    `json:"alasan"`
	StatusPengajuan  string    `json:"status"`
	TanggalKebutuhan string    `json:"tanggal_kebutuhan"`
	Pencaharian      string    `json:"pencarian"`
	Golongan         string    `json:"golongan"`
	TanggalPengajuan time.Time `json:"tanggal_pengajuan"`
	TanggalDisetujui string    `json:"tanggal_disetujui"`
}
type ReSponGetAdmin struct {
	IdPengajuan       int       `json:"id"`
	NamaPengajuan     string    `json:"nama_manager"`
	NamaDepartment    string    `json:"department"`
	Jumlah            string    `json:"jumlah"`
	Alasan            string    `json:"alasan"`
	Pencaharian       string    `json:"pencarian"`
	TanggalKebutuhan  string    `json:"tanggal_kebutuhan"`
	MaksimalGaji      float64   `json:"gaji"`
	NamaEvaluasi      string    `json:"nama_admin"`
	NamaVerifikasi    string    `json:"nama_vicepresident"`
	NamaPersetujuan   string    `json:"nama_direksi"`
	StatusPengajuan   string    `json:"status_pengajuan"`
	Golongan          string    `json:"golongan"`
	TanggalVerifikasi string    `json:"tanggal_verifikasi"`
	TanggalEvaluasi   string    `json:"tanggal_evalusi"`
	TanggalPengajuan  time.Time `json:"tanggal_pengajuan"`
	TanggalDisetujui  string    `json:"tanggal_disetujui"`
}
