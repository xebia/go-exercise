package datastore

//go:generate mockgen -source=dataStorer.go -destination=dataStorerMocks.go -package=datastorer PatientStorer

type PatientStorer interface {
	GetPatientOnUid(uid string) (Patient, bool, error)
	PutPatientOnUid(patient Patient) error
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
