package models

import "encoding/json"

// Organization represents the fields associated with a company or organization in Opus.
type Organization struct {
	ID      string `gorm:"primary_key" json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
}

// OrganizationCreate defines the data required for creating an Organization.
type OrganizationCreate struct {
	Name string `json:"name"`
}

// Serialize implements the Serializer interface for event publishing
func (o Organization) Serialize() ([]byte, error) {
	return json.Marshal(o)
}