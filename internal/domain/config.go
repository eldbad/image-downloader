package domain

import (
	"fmt"
	"net/url"
)

type Config struct {
	url   *url.URL
	flags []string
}

func (c Config) Url() *url.URL {
	return c.url
}

func (c Config) Flags() []string {
	return c.flags
}

func NewConfig(args []string) (*Config, error) {
	if len(args) != 2 {
		fmt.Println("Please, do learn how to return an error. Thanks")
	}

	url, err := url.Parse(args[1])

	return &Config{url, nil}, err
}
