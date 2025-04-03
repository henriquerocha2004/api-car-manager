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

type Client struct {
	Id         ulid.ULID
	Name       string
	LastName   string
	EntityType entityType
	Document   document.Document
	BirthDay   *time.Time
	Addresses  []Address
	Phones     []Phone
}

func (c *Client) validate() error {
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

func (c *Client) checkEntity() error {
	if c.EntityType == individual || c.EntityType == legalEntity {
		return nil
	}

	return errors.New("the entity type informed is invalid")
}

func (c *Client) validateDocument() error {
	return c.Document.Validate()
}

func (c *Client) addAddress(street, neighborhood, city, state, zipCode string) {
	address := Address{
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		ZipCode:      zipCode,
	}

	c.Addresses = append(c.Addresses, address)
}

func (c *Client) addPhone(description, number string) {
	phone := Phone{
		Description: description,
		Number:      number,
	}

	c.Phones = append(c.Phones, phone)
}
