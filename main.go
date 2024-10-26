package imap

import (
	"context"

	"github.com/scorify/schema"
)

type Schema struct {
	Server   string `key:"server"`
	Port     int    `key:"port" default:"143"`
	Username string `key:"username"`
	Password string `key:"password"`
	Secure   bool   `key:"secure"`
}

func Validate(config string) error {
	conf := Schema{}

	err := schema.Unmarshal([]byte(config), &conf)
	if err != nil {
		return err
	}

	return nil
}

func Run(ctx context.Context, config string) error {
	conf := Schema{}

	err := schema.Unmarshal([]byte(config), &conf)
	if err != nil {
		return err
	}

	return nil
}
