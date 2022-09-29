package client

import "github.com/TovarischSuhov/go-template/internal/domain"

type Client interface {
	GetTopicsList() (*domain.Response, error)
}
