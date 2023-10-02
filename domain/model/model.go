package model

import "time"

type User struct {
	Id        string
	Role      string
	Nip       string
	Password  string
	Name      string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}
