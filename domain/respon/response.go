package respon

import "time"

type ResponseUser struct {
	Id        int
	Role      string    `json:"role"`
	Nip       string    `json:"nip"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdat"`
	DeletedAt time.Time
	UpdateAt  time.Time
}
