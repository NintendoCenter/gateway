package main

import "github.com/caarlos0/env"

type Config struct {
	CollectionAddr string `env:"COLLECTION_GRPC_ADDR" envDefault:"localhost:9092"`
	WebPort        string `env:"PORT" envDefault:"8080"`
}


func NewConfig() (*Config, error) {
	cfg := new(Config)
	return cfg, env.Parse(cfg)
}