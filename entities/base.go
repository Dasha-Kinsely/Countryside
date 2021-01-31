package entities

import "time"

type Base struct {
	CreatedAt time.Time `json:"creation" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated" gorm:"default:CURRENT_TIMESTAMP"`
}
