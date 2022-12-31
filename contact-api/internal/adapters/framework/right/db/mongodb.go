package db

import (
	"github.com/sm/contact-api/config"
	"github.com/sm/contact-api/internal/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (dbAdapter *Adapter) FindAll() ([]models.Contact, error) {
	var contacts []models.Contact
	err := dbAdapter.db.C(COLLECTION).Find(bson.M{}).All(&contacts)
	return contacts, err
}

func (dbAdapter *Adapter) FindByID(id string) (models.Contact, error) {
	var contact models.Contact
	err := dbAdapter.db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&contact)
	return contact, err
}

func (dbAdapter *Adapter) Insert(contact models.Contact) (error) {
	err := dbAdapter.db.C(COLLECTION).Insert(&contact)
	return err
}

func (dbAdapter *Adapter) Update(contact models.Contact) (error) {
	err := dbAdapter.db.C(COLLECTION).UpdateId(contact.ID, &contact)
	return err
}

func (dbAdapter *Adapter) Delete(contact models.Contact) (error) {
	err := dbAdapter.db.C(COLLECTION).Remove(&contact)
	return err
}