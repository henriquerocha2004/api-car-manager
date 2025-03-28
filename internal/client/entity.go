package client

import (
	"errors"
	"time"

	"github.com/henriquerocha2004/api-car-manager/pkg/document"
	"github.com/oklog/ulid/v2"
)

type entityType string

const (
	individual  entityType = "PF"
	legalEntity entityType = "PJ"
)

type client struct {
	id         ulid.ULID
	name       string
	lastName   string
	entityType entityType
	document   document.Document
	birthDay   *time.Time
	addresses  []address
}

type address struct {
	street       string
	neighborhood string
	city         string
	state        string
	zipCode      string
}

func new(name, lastname, entity, docNumber, birthDay string) *client {
	birth, _ := time.Parse("2006-01-02", birthDay)
	c := &client{
		id:         ulid.Make(),
		name:       name,
		lastName:   lastname,
		entityType: entityType(entity),
		document:   document.Document(docNumber),
		birthDay:   &birth,
	}

	return c
}

func (c *client) validate() error {
	err := c.validateDocument()
	if err != nil {
		return err
	}

	err = c.checkEntity()
	if err != nil {
		return err
	}

	return nil
}

func (c *client) checkEntity() error {
	if c.entityType == individual || c.entityType == legalEntity {
		return nil
	}

	return errors.New("the entity type informed is invalid")
}

func (c *client) validateDocument() error {
	return c.document.Validate()
}

func (c *client) addAddress(street, neighborhood, city, state, zipCode string) error {
	address := address{
		street:       street,
		neighborhood: neighborhood,
		city:         city,
		state:        state,
		zipCode:      zipCode,
	}

	c.addresses = append(c.addresses, address)
}
