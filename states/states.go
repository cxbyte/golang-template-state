package states

import (
	"fmt"
)

type IState interface {
	Process()
	Close()
}

type State struct {
	IState
}

func (s State) Process() {
	fmt.Println("Can't process order");
}

func (s State) Close() {
	fmt.Println("Can't close order");
}