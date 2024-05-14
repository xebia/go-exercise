package datastore

import (
	"sync"
)

type inMemoryPatientStore struct {
	sync.Mutex
	patients map[string]Patient
}

func New() PatientStorer {
	return &inMemoryPatientStore{
		patients: map[string]Patient{},
	}
}
func (ps *inMemoryPatientStore) PutPatientOnUid(patient Patient) error {
	ps.Lock()
	defer ps.Unlock()

	ps.patients[patient.UID] = patient

	return nil
}

func (ps *inMemoryPatientStore) GetPatientOnUid(patientUID string) (Patient, bool, error) {
	ps.Lock()
	defer ps.Unlock()

	patient, found := ps.patients[patientUID]

	return patient, found, nil
}
