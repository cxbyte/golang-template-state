package states

import "fmt"

type NewState struct {
	State
}

func (s NewState) Process() {
	fmt.Println("Process order");
}