package models

type Client struct {
	id       string
	password string
}

func (c Client) GetID() string {
	return c.id
}

func (c Client) GetPassword() string {
	return c.password
}
