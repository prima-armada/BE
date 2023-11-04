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
	Token    string `json:"token"`
}
type ResponseDeparment struct {
	Id             int       `json:"id_departments"`
	NameDepartment string    `json:"nama_departments"`
	CreatedAt      time.Time `json:"createdat"`
	UpdateAt       time.Time `json:"updatedat"`
}
