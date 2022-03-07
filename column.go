package main

import "strconv"

type Column struct {
	ID                   string
	status               string
	amountOfElevators    int
	amountOfFloorsServed int
	servedFloorList      []int
	isBasement           bool
	elevatorsList        []*Elevator
	callButtonList       []*CallButton
}

func NewColumn(_id string, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	column := &Column{ID: _id, status: "online", amountOfElevators: _amountOfElevators, servedFloorList: _servedFloors}
	column.elevatorsList = make([]*Elevator, 0)
	column.callButtonList = make([]*CallButton, 0)
	column.amountOfFloorsServed = len(_servedFloors)
	column.createElevators()
	column.createCallButtons()
	return column
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return elevator
}

type BestElevatorInformations struct {
	bestElevator *Elevator
	bestScore    int
	referenceGap int
}

func checkIfElevatorIsBetter(scoreToCheck int, newElevator *Elevator,
	bestElevatorInformations *BestElevatorInformations,
	floor int) *BestElevatorInformations {
	if scoreToCheck < bestElevatorInformations.bestScore {
		bestElevatorInformations.bestScore = scoreToCheck
		bestElevatorInformations.bestElevator = newElevator
		bestElevatorInformations.referenceGap = Abs(newElevator.currentFloor - floor)
	} else if bestElevatorInformations.bestScore == scoreToCheck {
		gap := Abs(newElevator.currentFloor - floor)
		if bestElevatorInformations.referenceGap > gap {
			bestElevatorInformations.bestElevator = newElevator
			bestElevatorInformations.referenceGap = gap
		}
	}
	return bestElevatorInformations
}

func (c *Column) findElevator(requestedFloor int, requestedDirection string) *Elevator {
	bestElevatorInformations := &BestElevatorInformations{nil, 6, 10000000}
	if requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				bestElevatorInformations = checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
			}
			if elevator.currentFloor == 1 && elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			}
			if 1 > elevator.currentFloor && elevator.direction == "up" {
				bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
			}
			if 1 < elevator.currentFloor && elevator.direction == "down" {
				bestElevatorInformations = checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
			} else {
				bestElevatorInformations = checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {

			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				bestElevatorInformations = checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, requestedFloor)
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				bestElevatorInformations = checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInformations = checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, requestedFloor)
			} else {
				bestElevatorInformations = checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, requestedFloor)
			}
		}
	}
	return bestElevatorInformations.bestElevator
}

func (c *Column) createElevators() {
	for i := 0; i < c.amountOfElevators; i++ {
		elevator := NewElevator(strconv.Itoa(elevatorID + (i + 1)))
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}
}

func (c *Column) createCallButtons() {
	if c.isBasement {
		buttonFloor := -1
		for i := 0; i < c.amountOfFloorsServed; i++ {
			callButton := NewCallButton(buttonFloor, "up")
			c.callButtonList = append(c.callButtonList, callButton)
			buttonFloor--
		}
	} else {
		buttonFloor := 1
		for i := 0; i < c.amountOfFloorsServed; i++ {
			callButton := NewCallButton(buttonFloor, "down")
			c.callButtonList = append(c.callButtonList, callButton)
			buttonFloor++
		}
	}
}
