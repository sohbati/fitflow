package model

import (
	"time"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Image represents a gym image
type Image struct {
	URL       string `json:"url"`
	Type      string `json:"type"`      // logo, interior, exterior, trainer
	IsPrimary bool   `json:"is_primary"`
}

// Images represents a slice of images
type Images []Image

// Value implements the driver.Valuer interface
func (i Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}

// Scan implements the sql.Scanner interface
func (i *Images) Scan(value interface{}) error {
	if value == nil {
		*i = nil
		return nil
	}
	
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-byte value into Images")
	}
	
	return json.Unmarshal(bytes, i)
}
