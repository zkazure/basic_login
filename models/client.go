package models

type Client struct {
	id string
}

func (c Client) GetID() string {
	return c.id
}
