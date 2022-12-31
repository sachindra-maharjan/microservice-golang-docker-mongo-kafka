package api

import (
	"github.com/sm/contact-api/internal/models"
	"github.com/sm/contact-api/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
} 

func(app Application) AllContacts() ([]models.Contact, error) {
 return nil, nil
}

func(app Application) FindByID(string) (models.Contact, error) {
	return models.Contact{}, nil
}

func(app Application) FindByEmail(string) (models.Contact, error) {
	return models.Contact{}, nil
}

func(app Application) CreateContact(models.Contact) (error) {
	return nil
}

func(app Application) UpdateContact(models.Contact) (models.Contact, error) {
	return models.Contact{}, nil
}

func(app Application) DeleteContact(string) (error) {
	return nil
}
