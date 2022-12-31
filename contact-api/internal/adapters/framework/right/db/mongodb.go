package db

import (
	"github.com/sm/contact-api/config"
	"github.com/sm/contact-api/internal/models"
	mgo "gopkg.in/mgo.v2"
)

type Adapter struct {
	db *mgo.Database
}

const (
	COLLECTION = "contacts"
)

func NewAdapter(dbConfg config.Config) (*Adapter, error) {
	session, err := mgo.Dial(dbConfg.Server)
	if err != nil {
		return nil, err
	}
	db := session.DB(dbConfg.Database)
	return &Adapter{db: db}, nil
}

func (db *Adapter) AllContacts() ([]models.Contact, error) {
	return nil, nil
}

func (db *Adapter) FindByID(string) (models.Contact, error) {
	return models.Contact{}, nil
}

func (db *Adapter) FindByEmail(string) (models.Contact, error) {
	return models.Contact{}, nil
}

func (db *Adapter) CreateContact(models.Contact) (error) {
	return nil
}

func (db *Adapter) UpdateContact(models.Contact) (models.Contact, error) {
	return models.Contact{}, nil
}

func (db *Adapter) DeleteContact(string) (error) {
	return nil
}