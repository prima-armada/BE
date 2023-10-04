package request

import "time"

type RequestUser struct {
	Id        string `json:"id" form:"id" validate:"required,min=5"`
	Role      string `json:"roles" form:"roles" validate:"required,min=5"`
	Nip       string `json:"nip" form:"nip" validate:"required,min=5"`
	Password  string `json:"password" form:"password" validate:"required,min=5"`
	Username  string `json:"username" form:"username" validate:"required,min=5"`
	Name      string `json:"nama" form:"nama" validate:"required,min=5"`
	IdHc      string
	IdManager string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}