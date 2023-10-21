package model

import "time"

type User struct {
	Id        int
	Role      string
	Nip       string `gorm:"size:191; primaryKey"`
	Password  string
	Username  string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
	Admins    []Admin   `gorm:"foreignKey:Nip;references:Nip"`
	Managers  []Manager `gorm:"foreignKey:Nip;references:Nip"`
}

type Manager struct {
	Id        int    `gorm:"primaryKey"`
	Nip       string `gorm:"size:191" `
	Nama      string
	Bagian    string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}

type Admin struct {
	Id        int    `gorm:"primaryKey"`
	Nip       string `gorm:"size:191" `
	Nama      string
	Bagian    string
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}

type Department struct {
	Id             int    `gorm:"primaryKey"`
	NamaDepartment string `gorm:"size:191" `
	CreatedAt      time.Time
	DeletedAt      time.Time
	UpdateAt       time.Time
}
