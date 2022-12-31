package ports

import (
	"github.com/sm/contact-api/internal/models"
)

type APIPort interface {
	AllContacts() ([]models.Contact, error)
	FindByID(string) (models.Contact, error)
	CreateContact(models.Contact) (error)
	UpdateContact(models.Contact) (error)
	DeleteContact(models.Contact) (error)
}