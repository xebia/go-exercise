package testdata

import "github.com/xebia/go-exercise/internal/db"

var _ db.Service = (*Mock)(nil)

type Mock struct {
	CreatePatientFunc   func(p db.Patient) error
	GetPatientByUIDFunc func(uid string) (db.Patient, error)
}

func (m *Mock) CreatePatient(patient db.Patient) error {
	return m.CreatePatientFunc(patient)
}

func (m *Mock) GetPatientByUID(uid string) (db.Patient, error) {
	return m.GetPatientByUIDFunc(uid)
}
