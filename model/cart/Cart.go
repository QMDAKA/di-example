package cart

import (
	_interface "di_example/interface"
)

type cart struct {
	database _interface.IDatabase
	logger _interface.ILogger
	email _interface.IEmailSender
}
func New(database _interface.IDatabase, logger _interface.ILogger, email _interface.IEmailSender) cart {
	c := cart {database, logger, email}
	return c
}
func (c cart) CheckOut(orderId, userId int32) {
	c.database.Save(userId)
	c.email.SendMail(orderId)
	c.logger.LogInfo("log checkout")
}
