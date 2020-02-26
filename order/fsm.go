//author: richard
package order

import (
	"errors"
	"fmt"
	"mms/api"
)

//fsm(s, e) -> (action, s')
type fsm struct {
	states map[string]*event
}

type event struct {
	e map[int]func (*api.Order) error
	state     string
	nextState string
}

func (s *fsm) CreateState(state string,nextState string) *event {
	key := fmt.Sprintf("%s:%s", state, nextState)
	if s.states == nil {
		s.states = make(map[string]*event)
		s.states[key] = &event{state: state, nextState:nextState}
	}
	if _, ok := s.states[key]; !ok {
		s.states[key] = &event{state: state, nextState:nextState}
	}
	return s.states[key]
}

func (s *fsm) QueryAction(state, nextState string, e int) (func(*api.Order) error, error) {
	key := fmt.Sprintf("%s:%s", state, nextState)
	if _, ok := s.states[key]; !ok {
		return nil, errors.New(fmt.Sprintf("fsm don't support %s", key))
	}
	event := s.states[key]
	if _, ok := event.e[e]; !ok {
		return nil, errors.New(fmt.Sprintf("fsm don't support %d", e))
	}
	return event.e[e], nil
}

func (s *event) CreateEvent(e int, f func (*api.Order) error) *event {
	if s.e == nil {
		s.e = make(map[int]func (*api.Order) error)
	}
	s.e[e] = f
	return s
}

