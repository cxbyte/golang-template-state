package main

import (
	"./states"
)

type Order struct {
	state OrderState
}

type OrderState struct {
	state states.IState
}

// Process order
func (o OrderState) Process() {
	o.state.Process();
}

// Close order
func (o OrderState) Close() {
	o.state.Close();
}

func main()  {
	o := Order{state: OrderState{states.NewState{}}};
	o.state.Close();
	o.state.Process();

	o = Order{state: OrderState{states.ProcessState{}}};
	o.state.Process();
	o.state.Close();

	o = Order{state: OrderState{states.CloseState{}}};
	o.state.Process();
	o.state.Close();
}