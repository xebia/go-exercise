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
	// TODO add Patient type structure
}
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

func (ds *inMemoryDataStore) CreatePatient(patient Patient) error {
	ds.Lock()
	defer ds.Unlock()

	uid := uuid.New().String()
	ds.patients[uid] = patient

	return nil
}
