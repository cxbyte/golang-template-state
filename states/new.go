package states

type NewState struct {
	State
}

func (s NewState) Process() (string, error) {
	return "Process order", nil;
}