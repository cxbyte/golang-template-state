package main

import (
	"golang-template-state/states"
	"fmt"
)

type Order struct {
	state OrderState
}

type OrderState struct {
	state states.IState
}

// Process order
func (o OrderState) Process() (string, error) {
	return o.state.Process();
}

// Close order
func (o OrderState) Close() (string, error) {
	return o.state.Close();
}

func main()  {
	o := Order{state: OrderState{states.NewState{}}};
	message, err := o.state.Close();
	fmt.Println(err);
	message, err = o.state.Process();
	fmt.Println(message);

	o = Order{state: OrderState{states.ProcessState{}}};
	message, err = o.state.Process();
	fmt.Println(err);
	message, err = o.state.Close();
	fmt.Println(message);

	o = Order{state: OrderState{states.CloseState{}}};
	message, err = o.state.Process();
	fmt.Println(err);
	message, err = o.state.Close();
	fmt.Println(err);
}