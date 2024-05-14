package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MarcGrol/patientregistration/internal/registration/protobuf"
)

func main() {
	cliArgs := parseArgs()

	log.Printf("args: %+v", cliArgs)

	client, cleanup, err := protobuf.NewGrpcClient(protobuf.DefaultPort)
	if err != nil {
		log.Fatalf("*** Error creating motification-client: %v", err)
	}
	defer cleanup()
	log.Printf("Created registration-client")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if cliArgs.command == "start-registration" {
		patientUid, err := startRegistration(ctx, client, cliArgs.bsn, cliArgs.name, cliArgs.email)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Patient %s registered", patientUid)
	} else if cliArgs.command == "complete-registration" {
		err = completeRegistration(ctx, client, cliArgs.patientUid, cliArgs.pinCode)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Registration completed for patient %s", cliArgs.patientUid)
	} else {
		log.Fatalf("Unrecognized command %s", cliArgs.command)
	}
}

func startRegistration(ctx context.Context, client protobuf.RegistrationServiceClient, bsn int, name, email string) (string, error) {
	resp, err := client.RegisterPatient(ctx, &protobuf.RegisterPatientRequest{
		Patient: &protobuf.Patient{
			BSN:      fmt.Sprintf("%d", bsn),
			FullName: name,
			Address: &protobuf.Address{
				PostalCode:  "3731TB",
				HouseNumber: 79,
			},
			Contact: &protobuf.Contact{
				EmailAddress: email,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("Error registering client: %s", err)
	}
	return resp.PatientUid, nil
}

func completeRegistration(ctx context.Context, client protobuf.RegistrationServiceClient, patientUid string, pinCode int) error {
	// TODO
	return nil
}

type args struct {
	command    string
	email      string
	name       string
	bsn        int
	pinCode    int
	patientUid string
}

func parseArgs() args {

	help := flag.Bool("help", false, "This help text")
	command := flag.String("command", "start-registration", "Command (start-registration, complete-registration or brute-force-registration)")

	bsn := flag.Int("bsn", 12345, "Bsn number of patient")
	name := flag.String("name", "Michael Jordan", "Name of patient")
	email := flag.String("email", "eva.marc@hetnet.nl", "Email address of patient")

	patientUid := flag.String("patient-uid", "", "Uid of patient")
	pinCode := flag.Int("pin-code", -1, "Pincode to confirm registration")

	flag.Parse()

	if help != nil && *help {
		fmt.Fprintf(os.Stderr, "\nUsage:\n")
		fmt.Fprintf(os.Stderr, "\t[flags]\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(-1)
	}

	return args{
		command: *command,
		email:   *email,
		name:    *name,
		bsn:     *bsn,

		pinCode:    *pinCode,
		patientUid: *patientUid,
	}
}
