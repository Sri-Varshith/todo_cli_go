package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task enter title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a new task enter id and title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Edit an id to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Edit an id to toggle a task")
	flag.BoolVar(&cf.List, "list", false, "Display all task")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error invalid format use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
