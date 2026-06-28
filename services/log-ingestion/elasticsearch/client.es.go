package elasticsearch

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type Client struct {
	ES *elasticsearch.Client
}

type Config struct {
	Addresses []string
	Username  string
	Password  string
	CloudId   string
}

func New(cfg Config) (*Client, error) {
	esCfg := elasticsearch.Config{
		Addresses: cfg.Addresses,
		Username:  cfg.Username,
		Password:  cfg.Password,
		CloudID:   cfg.CloudId,
	}

	// creating new client
	res, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		return &Client{}, fmt.Errorf("[ES ERROR] Error occured while creating new ES client connection: %w", err)
	}

	// Ping to client if cluster is recheable
	pingRes, err := res.Info()
	if err != nil {
		return &Client{}, fmt.Errorf("[ES ERROR] Cluster is unrecheable: %w", err)
	}
	defer pingRes.Body.Close()

	if pingRes.IsError() {
		return &Client{}, fmt.Errorf("[ES ERROR] ES Ping error: %w", err)
	}

	return &Client{ES: res}, nil
}
