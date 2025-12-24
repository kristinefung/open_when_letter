package repository

import (
	"open_when_letter/internal/model"

	"gorm.io/gorm"
)

type LetterBoxRepository interface {
	Create(letterBox *model.LetterBox) error
}

type letterBoxRepository struct {
	db *gorm.DB
}

func NewLetterBoxRepository(db *gorm.DB) LetterBoxRepository {
	return &letterBoxRepository{db: db}
}

func (r *letterBoxRepository) Create(box *model.LetterBox) error {
	return r.db.Create(box).Error
}

func (r *letterBoxRepository) FindByID(id string) (*model.LetterBox, error) {
	var box model.LetterBox
	err := r.db.First(&box, "id = ?", id).Error
	return &box, err
}

func (r *letterBoxRepository) FindByInviteCode(code string) (*model.LetterBox, error) {
	var box model.LetterBox
	err := r.db.Where("invite_code = ?", code).First(&box).Error
	return &box, err
}
