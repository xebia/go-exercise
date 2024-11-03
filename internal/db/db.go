package db

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")

type Service interface {
	CreatePatient(patient Patient) error
	GetPatientByUID(uid string) (Patient, error)
}

type Patient struct {
	UID                string
	BSN                string
	FullName           string
	Address            StreetAddress
	Contact            Contact
	RegistrationPin    int
	FailedPinCount     int
	RegistrationStatus RegistrationStatus
}

type StreetAddress struct {
	PostalCode  string
	HouseNumber int
}

type Contact struct {
	EmailAddress string
}
type RegistrationStatus int

const (
	Unregistered RegistrationStatus = iota
	Pending
	Registered
	Blocked
)

type inMemoryDataStore struct {
	sync.Mutex
	patients map[string]Patient
}

func NewService() Service {
	return &inMemoryDataStore{
		patients: map[string]Patient{},
	}
}

func (ds *inMemoryDataStore) GetPatientByUID(uid string) (Patient, error) {
	ds.Lock()
	defer ds.Unlock()

	patient, ok := ds.patients[uid]
	if !ok {
		return Patient{}, ErrNotFound
	}

	return patient, nil
}
func (ds *inMemoryDataStore) GetPatientOnUID(patientUID string) (Patient, bool, error) {
	ds.Lock()
	defer ds.Unlock()

	patient, found := ds.patients[patientUID]

	return patient, found, nil
}

func (ds *inMemoryDataStore) CreatePatient(patient Patient) error {
	ds.Lock()
	defer ds.Unlock()

	uid := uuid.New().String()
	ds.patients[uid] = patient

	return nil
}
