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
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlag() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo, specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo, specify index & title, id:new_title")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo, specify index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo, specify index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

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
		parts := strings.SplitAfterN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Invalid index for edit, error: %v", err)
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Delete != -1:
		todos.delete(cf.Delete)

	default:
		fmt.Printf("Invalid command")
	}
}
