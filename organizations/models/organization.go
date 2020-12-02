package models

import (
	"encoding/json"
)

// Organization represents the fields associated with an organization in Opus.
type Organization struct {
	ID      string `gorm:"primary_key" json:"id"`
	Name    string `json:"name" validate:"required,min=3,max=255"`
	Version int    `json:"version"`
}

// OrganizationCreate defines the data required for creating an Organization.
type OrganizationCreate struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

// Serialize implements the Serializer interface for event publishing
func (o Organization) Serialize() ([]byte, error) {
	return json.Marshal(o)
}
