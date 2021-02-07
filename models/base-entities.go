package models

import "time"

type BaseModel struct {
	CreationAt time.Time
	UpdatedAt  time.Time
}
