package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	cliArgs := parseArgs()

	log.Printf("args: %+v", cliArgs)

	// create http client and send requests

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
