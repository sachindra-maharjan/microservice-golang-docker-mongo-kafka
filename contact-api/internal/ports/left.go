package ports

import (
	"github.com/sm/contact-api/internal/models"
)

type APIPort interface {
	AllContacts() ([]models.Contact, error)
	FindByID(string) (models.Contact, error)
	FindByEmail(string) (models.Contact, error)
	CreateContact(models.Contact) (error)
	UpdateContact(models.Contact) (models.Contact, error)
	DeleteContact(string) (error)
}