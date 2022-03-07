package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	ID          int
	status      string
	calledFloor int
	direction   string
}

func NewCallButton(_floor int, _direction string) *CallButton { 
	return &CallButton{ID: 1, status: "online", calledFloor: _floor, direction: _direction}
}

