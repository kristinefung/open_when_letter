package service

import (
	"open_when_letter/internal/model"
	"open_when_letter/internal/repository"

	"github.com/google/uuid"
)

type LetterBoxService interface {
	CreateLetterBox(title, description, coverImageURL, ownerID string) (*model.LetterBox, error)
}

type letterBoxService struct {
	repo repository.LetterBoxRepository
}

func NewLetterBoxService(repo repository.LetterBoxRepository) LetterBoxService {
	return &letterBoxService{repo: repo}
}

func (s *letterBoxService) CreateLetterBox(title, description, coverImageURL, ownerID string) (*model.LetterBox, error) {
	inviteCode := "BOX-" + uuid.New().String()[:8]

	box := &model.LetterBox{
		Title:         title,
		Description:   description,
		CoverImageURL: coverImageURL,
		OwnerUserID:   ownerID,
		InviteCode:    inviteCode,
	}

	if err := s.repo.Create(box); err != nil {
		return nil, err
	}

	return box, nil
}

// func (s *letterBoxService) GetLetterBoxByInviteCode(code string) (*model.LetterBox, error) {
//     return s.repo.FindByInviteCode(code)
// }
