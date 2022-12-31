package ports

import (
	"github.com/sm/contact-api/internal/models"
)

type DBPort interface {
	FindAll() ([]models.Contact, error)
	FindByID(string) (models.Contact, error)
	Insert(models.Contact) (error)
	Update(models.Contact) (error)
	Delete(models.Contact) (error)
}