Rocket-Elevators-Golang-Controller
This my Go Residential Controller from the pseudocode. Enjoy!

This program controls a Column of elevators.

It sends an elevator when a user presses a call button on a floor and it takes
a user to its desired floor when a button is pressed from the inside of the elevator. 

The elevator selection is based on the current direction and status of every elevators. It priorizes the elevator going in the same direction the user wants to go and sends the closest one as well so that the user waits less time.

To be able to try the program, you need 

- To create column with the parameters: id, amount of Floors and amount of Elevators
- Set the status of the elevators (are they moving in a direction or idle?)
- To set which floor is requested and where each elevators are.
- Once this is done, you can decide the behaviour of each elevators and then start a scenario.


Here is an example of a potential test scenario for the program:

func scenario2() (*Column, *Elevator) {
	column := battery.columnsList[2]

	elevator1 := ElevatorDetails{1, "up", "stopped", []int{21}}
	elevator2 := ElevatorDetails{23, "up", "moving", []int{28}}
	elevator3 := ElevatorDetails{33, "down", "moving", []int{1}}
	elevator4 := ElevatorDetails{40, "down", "moving", []int{24}}
	elevator5 := ElevatorDetails{39, "down", "moving", []int{1}}

	column.setupElevators([]ElevatorDetails{elevator1, elevator2, elevator3, elevator4, elevator5})

	chosenColumn, chosenElevator := battery.assignElevator(36, "up")
	moveAllElevators(chosenColumn)
	return chosenColumn, chosenElevator
}