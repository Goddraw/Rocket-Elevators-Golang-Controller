package main

type Door struct {
	ID     int
	status string
}

func NewDoor(id int) *Door {
	door := &Door{ID: id, status: "closed"}
	return door
}
