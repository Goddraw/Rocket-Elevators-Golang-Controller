package main

import (
	"math"
)

//global variables
var elevatorID int = 1
var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Battery struct {
	columnID                  int
	ID                        int
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	status                    string
	columnsList               []*Column
	floorRequestButtonsList   []*FloorRequestButton
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {

	battery := &Battery{columnID: 1, ID: _id, amountOfColumns: _amountOfColumns, amountOfFloors: _amountOfFloors,
		amountOfBasements: _amountOfBasements, amountOfElevatorPerColumn: _amountOfElevatorPerColumn,
		status: "online",
	}
	battery.columnsList = make([]*Column, 0)
	battery.floorRequestButtonsList = make([]*FloorRequestButton, 0)
	if _amountOfBasements > 0 {
		battery.createBasementColumn(battery.amountOfColumns, battery.amountOfElevatorPerColumn)
		battery.createBasementFloorRequestButtons(battery.amountOfBasements)
		battery.amountOfColumns--
	}

	battery.createColumns()
	battery.createFloorRequestsButton()

	return battery
}

func (b *Battery) createBasementColumn(amountOfBasements int, amountOfElevatorsPerColumn int) {
	servedFloorList := make([]int, 0)
	floor := -1
	for i := 0; i < amountOfBasements; i++ {
		servedFloorList = append(servedFloorList, floor)
		floor--
	}
	column := NewColumn(string(alphabet[b.columnID-1]), b.amountOfElevatorPerColumn, servedFloorList, true)
	b.columnsList = append(b.columnsList, column)
	b.columnID++
}

func (b *Battery) createFloorRequestsButton() {
	for i := 0; i < b.amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(i+1, "up")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, floorRequestButton)
	}
}
func (b *Battery) createColumns() {
	var amountOfFloorsPerColumn = int(math.Ceil(float64(b.amountOfFloors) / float64(b.amountOfColumns)))
	floor := 1
	for i := 0; i < b.amountOfColumns; i++ {
		servedFloorList := make([]int, 0)
		for j := 0; j < amountOfFloorsPerColumn; j++ {
			if floor <= b.amountOfFloors {
				servedFloorList = append(servedFloorList, floor)
				floor++
			}
		}
		column := NewColumn(string(alphabet[b.columnID-1]), b.amountOfElevatorPerColumn, servedFloorList, false)
		b.columnsList = append(b.columnsList, column)
		b.columnID++
	}
}

func (b *Battery) createBasementFloorRequestButtons(amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < amountOfBasements; i++ {
		basementButton := NewFloorRequestButton(buttonFloor, "down")
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, basementButton)
		buttonFloor--
	}
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	var bestColumn *Column = nil
	for _, column := range b.columnsList {
		if contains(column.servedFloorList, _requestedFloor) {
			bestColumn = column
		}
	}
	return bestColumn
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {

	column := b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return column, elevator
}
