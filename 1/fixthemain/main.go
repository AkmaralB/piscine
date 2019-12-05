package main

import (
	"github.com/01-edu/z01"
)

type Door struct {
	state string
}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"
}
func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	if ptrDoor.state == "CLOSE" {
		return true
	}
	return false
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	if ptrDoor.state == "CLOSE" {
		return true
	}
	return false
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
