package main

import (
	"di_example/implement"
	"di_example/interface"
	"di_example/model/cart"
)

func main() {
	var databaseImpl _interface.IDatabase
	databaseImpl = implement.FakeDatabase{}
	var emailImpl _interface.IEmailSender
	emailImpl = implement.Email{}
	var loggerImpl _interface.ILogger
	loggerImpl = implement.Logger{}
	cart := cart.New(databaseImpl, loggerImpl, emailImpl)
	cart.CheckOut(21, 21)
}
