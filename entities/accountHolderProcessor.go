package modelbank

import (
	"time"
)

type accountHolderProcessor interface {
	openAccount(accountHolder AccountHolder, currency string) *Account

	closeAccount()

	// intructPayment
	// This is a credit transfer initiated by accountholder. It
	// creates a paymenttransaction and stores it in financial institution's
	// context for further processing.
	intructPayment(crdtrName string,
		crdtrIban string,
		crdtBic string,
		crdtAmnt float64,
		crdtCrrncy string,
		purpose string,
		reqExecDate time.Time, dbtrAccnt *Account) *PaymentTransaction
}
