package model

import "time"

type User struct {
	Id        string
	Role      string
	Nip       string
	Password  string
	Username  string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}

type Manager struct {
	Id        string
	Nip       string
	Nama      string
	Bagian    string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}

type HumanCapital struct {
	Id        string
	Nip       string
	Nama      string
	Bagian    string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}
