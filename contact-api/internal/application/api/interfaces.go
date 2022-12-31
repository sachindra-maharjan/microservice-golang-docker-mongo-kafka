package api

import (
	. "github.com/sm/contact-api/internal/models"
)


type ContactApi interface {
	AllContacts() ([]Contact, error)
	FindByID(string) (Contact, error)
	CreateContact(Contact) (error)
	UpdateContact(Contact) (error)
	DeleteContact(Contact) (error)
}