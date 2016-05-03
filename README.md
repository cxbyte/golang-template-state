Example of State pattern on golang

Order{state: OrderState{states.NewState{}}};
Order{state: OrderState{states.ProcessState{}}}
Order{state: OrderState{states.CloseState{}}};

**Run**

go build -o state && ./state

**Tests**

go test