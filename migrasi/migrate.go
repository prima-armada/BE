package migrasi

import (
	"par/domain/model"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Department{})
	db.AutoMigrate(&model.Submission{})
	db.AutoMigrate(&model.FormulirKandidat{})
	db.AutoMigrate(&model.SoalInterview{})
	db.AutoMigrate(&model.InterviewKandidat{})
	db.AutoMigrate(&model.DetailProses{})
	db.AutoMigrate(&model.Position{})
	db.AutoMigrate(&model.SoalFPT{})
	db.AutoMigrate(&model.InterviewFPT{})
}
