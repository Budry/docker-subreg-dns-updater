package utils

import (
	"errors"
	"strings"
)

type Uri struct {
	SubDomain string
	Domain string
}

func NewUri(hostname string) (*Uri, error) {
	fragments := strings.Split(hostname, ".")

	if len(fragments) < 3 {
		return nil, errors.New("Invalid hostname format")
	}

	subDomain := strings.Join(fragments[:len(fragments)-2], ".")
	domain := strings.Join(fragments[len(fragments)-2:len(fragments)], ".")

	return &Uri{
		SubDomain: subDomain,
		Domain:    domain,
	}, nil
}