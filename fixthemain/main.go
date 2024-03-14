package main

import "fmt"

type Door struct {
	state string
}

func PrintStr(s string) {
	fmt.Println(s)
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSED"
	return true
}

func IsDoorOpen(Door Door) bool {
	PrintStr("Is the Door opened?")
	return Door.state == "OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed?")
	return ptrDoor.state == "CLOSED"
}

func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == "CLOSE" {
		CloseDoor(door)
	}
}
