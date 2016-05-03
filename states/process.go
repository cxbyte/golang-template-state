package states

import "fmt"

type ProcessState struct {
	State
}

func (s ProcessState) Close() {
	fmt.Println("Close order");
}
