package implement

import (
	"fmt"
)
type FakeDatabase struct {
}
func (db FakeDatabase) Save(id int32) {
	fmt.Println("Save ", id)
}