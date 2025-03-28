package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateClient(t *testing.T) {
	client := new("Henrique", "Rocha", "PF", "48583696098", "2001-03-03")
	assert.Equal(t, client.name, "Henrique")
}

func TestShouldReturnErrorIfTryCreateClientWithInvalidDocument(t *testing.T) {
	client := new("Henrique", "Rocha", "PF", "48583696600", "2001-03-03")
	err := client.validate()
	assert.Error(t, err, "invalid document")
}

func TestShouldReturnErrorIfTryCreateClientWithIinvalidEntityType(t *testing.T) {
	client := new("Henrique", "Rocha", "AA", "48583696098", "2001-03-03")
	err := client.validate()
	assert.Error(t, err, "the entity type informed is invalid")
}
