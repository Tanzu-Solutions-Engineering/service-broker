package main

import (
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/Pivotal-Field-Engineering/service-broker/broker"
	"github.com/pivotal-cf/brokerapi"
)

func main() {

	logger := lager.NewLogger("broker")
	workshopBroker := broker.Broker{}

	username := os.Getenv("SECURITY_USER_NAME")
	if username == "" {
		logger.Debug("SECURITY_USER_NAME is not set! Defaulting to \"pivotal\"")
		username = "pivotal"
	}
	password := os.Getenv("SECURITY_USER_PASSWORD")
	if password == "" {
		logger.Debug("SECURITY_USER_PASSWORD is not set! Defaulting to \"pivotal\"")
		password = "pivotal"
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logger.Debug("PORT is not set! Defaulting to \"8080\"")
		port = "8080"
	}

	brokerCreds := brokerapi.BrokerCredentials{
		Username: username,
		Password: password,
	}

	handler := brokerapi.New(workshopBroker, logger, brokerCreds)

	err := http.ListenAndServe(":"+port, handler)
	logger.Fatal("HTTP Service Failed", err, lager.Data{})
}
