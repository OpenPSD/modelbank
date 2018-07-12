package entities

import "time"

type financialInstitutionProcessor interface {
	onboardPerson(name string,
		address string,
		dateOfBirth time.Time,
		cityOfBirth string,
		countryOfBirth string) *AccountHolder

	onboardOrganisation(name string,
		address string,
		tin string) *AccountHolder

	// receivePaymentInstruction
	// receives incoming paymentinstructions from other banks and stores them
	// for further processing.
	receivePaymentInstruction(paymentInstruction PaymentInstruction)

	// sendPaymentInstruction
	// sends outgoing paymentinstructions to other banks.
	sendPaymentInstruction(paymentInstruction PaymentInstruction)

	// processIncomingPaymentInstructions
	// books all received incoming paymentinstructions to corresponding accounts' ledgers.
	// Incoming paymentinstructions can either be received from other financial institutions or
	// created by accountholders wihin the same bank.
	processIncomingPaymentInstructions()

	// createOutgoingPaymentInstructions
	// takes all existing paymenttransactions created by accountholders and creates
	// outgoing paymentInstructions. Books all payment transactions on accounts's ledgers.
	createOutgoingPaymentInstructions()

	GetAccounts() []Account
	GetAccount(iban string) Account
}
