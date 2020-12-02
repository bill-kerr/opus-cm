package models

import (
	"encoding/json"
	"opus-cm/organizations/database"

	"github.com/gofrs/uuid"
)

// Organization represents the fields associated with an organization in Opus.
type Organization struct {
	ID      string `gorm:"primary_key" json:"id"`
	Name    string `gorm:"unique" json:"name" validate:"required,min=3,max=255"`
	Version int    `json:"version"`
}

// NewOrganization creates a new organization with some initialized fields.
func NewOrganization(name string) Organization {
	ID, _ := uuid.NewV4()
	return Organization{
		ID: ID.String(),
		Name: name,
		Version: 0,
	}
}

// OrganizationCreate defines the data required for creating an Organization.
type OrganizationCreate struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

// SerializeEvent implements the EventSerializer interface for event publishing
func (o Organization) SerializeEvent() ([]byte, error) {
	return json.Marshal(o)
}

// FindOne returns the first record in the database matching the provided condition. If none is found, returns an error.
func FindOne(condition Organization) (Organization, error) {
	db := database.GetDB()
	var org Organization
	err := db.Where(condition).First(&org).Error
	return org, err
}

// Save persists the organization to the database. If the operation fails, it returns an error.
func (o *Organization) Save() error {
	db := database.GetDB()
	tx := db.Create(o)
	return tx.Error
}