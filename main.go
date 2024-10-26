package imap

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/emersion/go-imap/client"
	"github.com/scorify/schema"
)

type Schema struct {
	Server   string `key:"server"`
	Port     int    `key:"port" default:"143"`
	Username string `key:"username"`
	Password string `key:"password"`
	Mailbox  string `key:"mailbox" default:"INBOX"`
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

	deadline, ok := ctx.Deadline()
	if !ok {
		return fmt.Errorf("context deadline is not set")
	}

	dialer := &net.Dialer{
		Deadline: deadline,
	}

	var imapClient *client.Client
	connStr := fmt.Sprintf("%s:%d", conf.Server, conf.Port)

	if conf.Secure {
		imapClient, err = client.DialWithDialerTLS(
			dialer,
			connStr,
			&tls.Config{InsecureSkipVerify: true},
		)
	} else {
		imapClient, err = client.DialWithDialer(
			dialer,
			connStr,
		)
	}
	if err != nil {
		return err
	}
	defer imapClient.Logout()

	err = imapClient.Login(conf.Username, conf.Password)
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}

	_, err = imapClient.Select(conf.Mailbox, true)
	if err != nil {
		return fmt.Errorf("failed opening mailbox %q: %w", conf.Mailbox, err)
	}

	return nil
}
