package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role     string
	Nip      string
	Password string
	Username string
	Admins   []Admin   `gorm:"foreignKey:ID"`
	Managers []Manager `gorm:"foreignKey:ID"`
}

type Manager struct {
	gorm.Model
	Nip    string
	Nama   string
	Bagian string
}

type Admin struct {
	gorm.Model
	Nip    string
	Nama   string
	Bagian string
}

type Department struct {
	gorm.Model
	NamaDepartment string `gorm:"size:191" `
}
