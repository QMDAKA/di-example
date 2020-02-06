package implement

import "fmt"

type Logger struct {
}
func (l Logger) LogInfo(info string) {
	fmt.Println("log info: ", info)
}