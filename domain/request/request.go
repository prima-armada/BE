package request

import "time"

type RequestUser struct {
	Id        string `json:"id" form:"id" validate:"required,min=5"`
	Role      string `json:"roles" form:"roles" validate:"required,min=5"`
	Nip       string `json:"nip" form:"nip" validate:"required,min=5"`
	Password  string `json:"password" form:"password" validate:"required,min=5"`
	Name      string `json:"Name" form:"Name" validate:"required,min=5"`
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}
