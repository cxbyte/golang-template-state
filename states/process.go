package states

type ProcessState struct {
	State
}

func (s ProcessState) Close() (string, error) {
	return "Close order", nil;
}
