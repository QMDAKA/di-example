package implement

import (
	"fmt"
)
type Database struct {
}
func (db Database) Save(id int32) {
	fmt.Println("Save ", id)
}