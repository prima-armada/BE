package respon

import "time"

type ResponseUser struct {
	Role      string    `json:"role"`
	Nip       string    `json:"nip"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdat"`
	Nama      string    `json:"nama"`
	Bagian    string    `json:"bagian"`
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

type ResponSubmission struct {
	IdPengajuan      int       `json:"user"`
	IdDepartment     uint      `json:"department"`
	Jumlah           string    `json:"jumlah"`
	Alasan           string    `json:"alasan"`
	StatusPengajuan  string    `json:"status"`
	TanggalKebutuhan string    `json:"tanggal_kebutuhan"`
	Pencaharian      string    `json:"pencaharian"`
	Golongan         string    `json:"golongan"`
	TanggalPengajuan time.Time `json:"pengajuan_tanggal"`
	KodePengajuan    string    `json:"kode_pengajuan"`
}
type ReSponGetManager struct {
	IdPengajuan       int       `json:"id"`
	NamaManager       string    `json:"nama_manager"`
	NamaDepartment    string    `json:"department"`
	Jumlah            string    `json:"jumlah"`
	Alasan            string    `json:"alasan"`
	StatusPengajuan   string    `json:"status_pengajuan"`
	TanggalKebutuhan  string    `json:"tanggal_kebutuhan"`
	Pencaharian       string    `json:"pencarian"`
	Golongan          string    `json:"golongan"`
	TanggalPengajuan  time.Time `json:"tanggal_pengajuan"`
	KodePengajuan     string    `json:"kode_pengajuan"`
	TanggalVerifikasi string    `json:"tanggal_verifikasi"`
	TanggalDisetujui  string    `json:"tanggal_disetujui"`
	TanggalEvaluasi   string    `json:"tanggal_evalusi"`
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
	KodePengajuan    string    `json:"kode_pengajuan"`
}
type ReSponGetAdmin struct {
	IdPengajuan       int       `json:"id"`
	NamaPengajuan     string    `json:"nama_user"`
	NamaDepartment    string    `json:"department"`
	Jumlah            string    `json:"jumlah"`
	Alasan            string    `json:"alasan"`
	Pencaharian       string    `json:"pencarian"`
	TanggalKebutuhan  string    `json:"tanggal_kebutuhan"`
	MaksimalGaji      float64   `json:"gaji"`
	NamaEvaluasi      string    `json:"nama_admin"`
	NamaVerifikasi    string    `json:"nama_verifikasi"`
	NamaPersetujuan   string    `json:"nama_direksi"`
	StatusPengajuan   string    `json:"status_pengajuan"`
	Golongan          string    `json:"golongan"`
	TanggalVerifikasi string    `json:"tanggal_verifikasi"`
	TanggalEvaluasi   string    `json:"tanggal_evalusi"`
	TanggalPengajuan  time.Time `json:"tanggal_pengajuan"`
	TanggalDisetujui  string    `json:"tanggal_disetujui"`
	KodePengajuan     string    `json:"kode_pengajuan"`
	PosisiKosong      string    `json:"posisi"`
}
type ReSponGetPresident struct {
	IdPengajuan       int       `json:"id"`
	NamaManager       string    `json:"nama_manager"`
	NamaDepartment    string    `json:"department"`
	Jumlah            string    `json:"jumlah"`
	Alasan            string    `json:"alasan"`
	StatusPengajuan   string    `json:"status_pengajuan"`
	TanggalKebutuhan  string    `json:"tanggal_kebutuhan"`
	Pencaharian       string    `json:"pencarian"`
	Golongan          string    `json:"golongan"`
	TanggalPengajuan  time.Time `json:"tanggal_pengajuan"`
	Tanggalverifikasi string    `json:"tanggal_verifikasi"`
	KodePengajuan     string    `json:"kode_pengajuan"`
}
type ResponUpdateAdmin struct {
	IdEvaluasi       int
	StatusPengajuan  string    `json:"status" `
	TanggalDievalusi time.Time `json:"tanggal_eval" `
	MaksimalGaji     float64   `json:"gaji"`
}
type ResponUpdateVicePresident struct {
	IdVerifikasi      int
	StatusPengajuan   string    `json:"status"`
	TanggalVerifikasi time.Time `json:"verifikasi_tanggal"`
}
type ResponUpdateDireksi struct {
	IdPersetujuan      int       `json:"direksi"`
	StatusPengajuan    string    `json:"status"`
	TanggalPersetujuan time.Time `json:"persetujuan_tanggal"`
}

type ResFormulirKandidat struct {
	Id                   uint
	NamaManager          string `json:"manager"`
	KodePengajuan        string `json:"kode_pengajuan"`
	DepartementManager   string `json:"departement_manager"`
	NamaKandidat         string `json:"nama_kandidat"`
	ContactNumber        string `json:"contact_kandidat"`
	ContactYangDihubungi string `json:"contact_kerabat"`
	NomorContactDarurat  string `json:"nomor_kerabat"`
	InformasiJob         string `json:"info_job"`
	NipRefrensi          string `json:"nip_ref"`
	JenjangPendidikan    string `json:"jenjang_pendidikan"`
	NamaRefrensi         string `json:"nama_refrensi"`
	Alamat               string `json:"alamat"`
	Pengalaman           string `json:"pengalama"`
	AdminId              uint
}
type ResponSoal struct {
	Id          uint
	Kategori    string `json:"kategori"`
	Description string `json:"deskripsi"`
}
type ResPosisi struct {
	id          uint
	UserId      uint
	LevelKosong string
	Department  string
}
