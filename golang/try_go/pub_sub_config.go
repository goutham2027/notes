package main

import (
	"os"
)

type PubSubConfig struct {
	GCP_PROJECT          string
	PUB_SUB_TOPIC        string
	PUB_SUB_SUBSCRIPTION string
}

func GetPubSubConfig() PubSubConfig {
	GCP_PROJECT := os.Getenv("GCP_PROJECT")
	PUB_SUB_TOPIC := os.Getenv("PUB_SUB_TOPIC")

	var c PubSubConfig
	c.GCP_PROJECT = GCP_PROJECT
	c.PUB_SUB_TOPIC = PUB_SUB_TOPIC
	return c

}
