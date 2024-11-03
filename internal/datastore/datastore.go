package datastore

import "sync"

type Service interface {
	GetPatientOnUID(uid string) (Patient, bool, error)
	PutPatientOnUID(patient Patient) error
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
func (ds *inMemoryDataStore) PutPatientOnUID(patient Patient) error {
	ds.Lock()
	defer ds.Unlock()

	ds.patients[patient.UID] = patient

	return nil
}

func (ds *inMemoryDataStore) GetPatientOnUID(patientUID string) (Patient, bool, error) {
	ds.Lock()
	defer ds.Unlock()

	patient, found := ds.patients[patientUID]

	return patient, found, nil
}
