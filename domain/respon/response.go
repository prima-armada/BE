package respon

import "time"

type ResponseUser struct {
	Id        string
	Role      string    `json:"role"`
	Nip       string    `json:"nip"`
	Password  string    `json:"password"`
	Name      string    `json:"Name"`
	CreatedAt time.Time `json:"createdat"`
	DeletedAt time.Time
	UpdateAt  time.Time
}
