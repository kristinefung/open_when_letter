package model

import "time"

type LetterBox struct {
	ID            string    `gorm:"primaryKey;type:varchar(36)" json:"boxId"`
	OwnerUserID   string    `gorm:"type:varchar(36);not null" json:"-"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Description   string    `gorm:"type:text" json:"description,omitempty"`
	CoverImageURL string    `gorm:"type:varchar(512)" json:"coverImageUrl,omitempty"`
	InviteCode    string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"inviteCode"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
