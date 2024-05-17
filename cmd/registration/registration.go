package main

import (
	"log"

	"github.com/xebia/go-exercise/internal/registration"
	"github.com/xebia/go-exercise/internal/registration/datastore"
	"github.com/xebia/go-exercise/internal/registration/email"
	"github.com/xebia/go-exercise/internal/registration/protobuf"
)

func main() {
	patientStore := datastore.New()
	emailSender := email.New()

	service := registration.NewRegistrationService(patientStore, emailSender)
	err := protobuf.StartGrpcServer(protobuf.DefaultPort, service)
	if err != nil {
		log.Fatalf("Error starting registration server: %s", err)
	}
}
