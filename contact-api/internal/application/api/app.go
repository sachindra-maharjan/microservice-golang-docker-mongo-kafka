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
	contacts, err :=	app.db.FindAll()
	if err != nil {
		return nil, err
	}
 return contacts, nil
}

func(app Application) FindByID(id string) (models.Contact, error) {
	contact, err := app.db.FindByID(id)
	if err != nil {
		return models.Contact{}, err
	}
	return contact, nil
}

func(app Application) CreateContact(contact models.Contact) (error) {
	err := app.db.Insert(contact)
	return err
}

func(app Application) UpdateContact(contact models.Contact) (error) {
	err := app.db.Update(contact)
	return err
}

func(app Application) DeleteContact(contact models.Contact) (error) {
	err := app.db.Delete(contact)
	return err
}
