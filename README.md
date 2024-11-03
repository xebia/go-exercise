# Humanitech

## Scenario:

We are building a new platform for patients to gain insights in their medical history. The platform is called
humanitech. To do so, we have to allow
patients to register with our platform. When registering, we will validate that the email address provided belongs to
the patient in question. We will do so by sending the patient a randomly generated pincode which they'll have to enter
to verify their email.

## Expectation:

1. Implement the an api to register patients
2. Create a new api to complete the registration flow
3. Extra: What more apis do we need, get creative :)

You can make use of the following building blocks 👇

### Building blocks:

- PatientStorer - interface + implementation
- EmailSender - interface + implementation
- Server boilerplate + testsuite
- [api schema](./humanitech.yaml) for implementation or inspiration
- Makefile ready to go with handy commands

### Registration flow

![alt text](docs/registration-flow.png)

### Sequence Diagram

![alt text](docs/sequence-diagram.png)