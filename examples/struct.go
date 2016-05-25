// +build ignore

package main

import (
	"errors"
	"fmt"

	"github.com/EricYT/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) error { return d.enterState(e) },
		},
	)

	return d
}

func (d *Door) enterState(e *fsm.Event) error {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
	return errors.New("crash")
}

func main() {
	door := NewDoor("heaven")

	err := door.FSM.Event("open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event("close")
	if err != nil {
		fmt.Println(err)
	}
}
