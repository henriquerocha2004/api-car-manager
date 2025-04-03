package client

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type Reader interface {
	FindById(ctx context.Context, id ulid.ULID) (*Client, error)
	FindAll(ctx context.Context, criteria any) ([]*Client, error)
}

type Writer interface {
	Create(ctx context.Context, client *Client) error
	Update(ctx context.Context, client *Client) error
	Delete(ctx context.Context, id ulid.ULID) error
}

type Repository interface {
	Reader
	Writer
}
