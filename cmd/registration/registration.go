package main

import (
	"log"

	"github.com/MarcGrol/patientregistration/internal/registration"
	"github.com/MarcGrol/patientregistration/internal/registration/datastore"
	"github.com/MarcGrol/patientregistration/internal/registration/email"
	"github.com/MarcGrol/patientregistration/internal/registration/protobuf"
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
