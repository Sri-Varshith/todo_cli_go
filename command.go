package main

import "flag"

type CmdFlags struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
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

