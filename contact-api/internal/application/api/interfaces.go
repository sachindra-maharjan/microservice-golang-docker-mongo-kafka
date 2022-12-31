package api

import (
	. "github.com/sm/contact-api/internal/models"
)


type ContactApi interface {
	AllContacts() ([]Contact, error)
	FindByID(string) (Contact, error)
	FindByEmail(string) (Contact, error)
	CreateContact(Contact) (error)
	UpdateContact(Contact) (Contact, error)
	DeleteContact(string) (error)
}