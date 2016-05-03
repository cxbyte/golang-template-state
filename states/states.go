package states

import "errors"

type IState interface {
	Process() (string, error)
	Close() (string, error)
}

type State struct {
	IState
}

func (s State) Process() (string, error) {
	return "", errors.New("Can't process order");
}

func (s State) Close() (string, error) {
	return "", errors.New("Can't close order");
}