package implement

import (
	"fmt"
)
type Email struct {
}
func (e Email) SendMail(id int32) {
	fmt.Println("Send mail to", id)
}