package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int64           `gorm:"primarykey"`
	Metadata  json.RawMessage `json:"metadata" gorm:"type:jsonb;default: '{}'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BaseModelWithoutID struct {
	Metadata  json.RawMessage `json:"metadata" gorm:"type:jsonb;default: '{}'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *BaseModel) SetMetadata(metadata any) {
	data, _ := json.Marshal(metadata)
	m.Metadata = data
}
