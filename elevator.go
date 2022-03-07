package main

import "sort"

type Elevator struct {
	ID                    string
	status                string
	currentFloor          int
	direction             string
	door                  Door
	floorRequestsList     []int
	completedRequestsList []int
}

func NewElevator(_elevatorID string) *Elevator {
	elev := &Elevator{ID: _elevatorID, status: "idle", currentFloor: 1,
		direction: "none", door: *NewDoor(1), floorRequestsList: make([]int, 0), completedRequestsList: make([]int, 0)}
	return elev
}

func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		destination := e.floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			e.sortFloorList()
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"

		e.completedRequestsList = append(e.completedRequestsList, e.floorRequestsList[0])
		e.floorRequestsList = e.floorRequestsList[1:]
	}
	e.status = "idle"
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Ints(e.floorRequestsList) // Ascending Order
	} else if e.direction == "down" {
		sort.Ints(e.floorRequestsList)
		e.floorRequestsList = Reverse(e.floorRequestsList) // Descending Order
	}
}

func (e *Elevator) addNewRequest(requestedFloor int) {
	if !contains(e.floorRequestsList, requestedFloor) {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)
	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}
}
