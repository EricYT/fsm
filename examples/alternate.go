// +build ignore

package main

import (
	"fmt"

	"github.com/EricYT/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"idle",
		fsm.Events{
			{Name: "scan", Src: []string{"idle"}, Dst: "scanning"},
			{Name: "working", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"scanning"}, Dst: "scanning"},
			{Name: "situation", Src: []string{"idle"}, Dst: "idle"},
			{Name: "finish", Src: []string{"scanning"}, Dst: "idle"},
		},
		fsm.Callbacks{
			"scan": func(e *fsm.Event) error {
				fmt.Println("after_scan: " + e.FSM.Current())
				return nil
			},
			"working": func(e *fsm.Event) error {
				fmt.Println("working: " + e.FSM.Current())
				return nil
			},
			"situation": func(e *fsm.Event) error {
				fmt.Println("situation: " + e.FSM.Current())
				return nil
			},
			"finish": func(e *fsm.Event) error {
				fmt.Println("finish: " + e.FSM.Current())
				return nil
			},
		},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event("scan")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("1:" + fsm.Current())

	err = fsm.Event("working")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2:" + fsm.Current())

	err = fsm.Event("situation")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("3:" + fsm.Current())

	err = fsm.Event("finish")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("4:" + fsm.Current())

}
