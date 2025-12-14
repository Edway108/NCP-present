package main

import log "github.com/sirupsen/logrus"

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"session":      "1ce3f6v",
		"product_type": "ticket",
		"quantity":     3,
	}).Info("New order created")
}
