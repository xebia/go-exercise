# patientregistration

## Scenario:

We are building a new platform for patients to gain insights in their medical history. To do so, we have to allow
patients to register with our platform. When registering, we will validate that the email address provided belongs to
the patient in question. We will do so by sending the patient a randomly generated pincode which they'll have to enter
to verify their email.

## Expectation:

1. Create the registration service to account for the aforementioned scenario.
2. Replace the inbound registration interface with a GRPC implementation

You can make use of the following building blocks 👇

### Building blocks:

- PatientStorer - interface + implementation
- EmailSender - interface + implementation
- Initial RegistrationService - interface + unit test
- protobuf contract with the initial start patient registration request

### Registration flow

![alt text](docs/registration-flow.png)

### Sequence Diagram

![alt text](docs/sequence-diagram.png)