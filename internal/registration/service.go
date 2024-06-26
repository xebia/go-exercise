package registration

import (
	"context"
	"fmt"

	"github.com/xebia/go-exercise/internal/registration/datastore"
	"github.com/xebia/go-exercise/internal/registration/email"
	"github.com/xebia/go-exercise/internal/registration/protobuf"
)

type Service struct {
	patientStore datastore.PatientStorer
	emailSender  email.EmailSender
	protobuf.UnimplementedRegistrationServiceServer
}

func NewRegistrationService(patientStore datastore.PatientStorer, emailSender email.EmailSender) *Service {
	return &Service{
		patientStore: patientStore,
		emailSender:  emailSender,
	}
}

func (rs *Service) RegisterPatient(ctx context.Context, req *protobuf.RegisterPatientRequest) (*protobuf.RegisterPatientResponse, error) {
	// TODO
	return nil, fmt.Errorf("Not implemented")
}

// TODO add CompletePatientRegistration
