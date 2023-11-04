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
	Nama     string
	Bagian   string
}

type Department struct {
	gorm.Model
	NamaDepartment string `gorm:"size:191" `
}
