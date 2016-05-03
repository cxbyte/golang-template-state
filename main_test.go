package main

import (
	"testing"
	"golang-template-state/states"
)

func TestNewStateProcess(t *testing.T)  {
	o := Order{state: OrderState{states.NewState{}}};

	_, err := o.state.Process();
	if err != nil  {
		t.Error("Expected nil, got ", err);
	}
}

func TestNewStateClose(t *testing.T)  {
	o := Order{state: OrderState{states.NewState{}}};

	_, err := o.state.Close();
	if err == nil  {
		t.Error("Expected Can't close order, got ", err);
	}
}