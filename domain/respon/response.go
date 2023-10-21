package respon

import "time"

type ResponseUser struct {
	Id        int
	Role      string    `json:"role"`
	Nip       string    `json:"nip"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdat"`
	DeletedAt time.Time
	UpdateAt  time.Time
}

type ResponseDeparment struct {
	NameDepartment string    `json:"nama_departments"`
	CreatedAt      time.Time `json:"createdat"`
	UpdateAt       time.Time `json:"updatedat"`
}
