package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
	"github.com/fatih/color"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil

}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	if len(*todos) == 0 {
		empty := color.New(color.Faint, color.Italic)
		empty.Println("✨ No tasks yet")
		return
	}

	tbl := table.New(os.Stdout)

	header := color.New(color.FgCyan, color.Bold)

	tbl.SetHeaders(
		header.Sprint("ID"),
		header.Sprint("Task"),
		header.Sprint("Done"),
		header.Sprint("Created"),
		header.Sprint("Completed"),
	)

	for index, t := range *todos {
		status := color.New(color.FgHiRed).Sprint("Pending")
		completedAt := "—"

		if t.Completed {
			status = color.New(color.FgGreen, color.Bold).Sprint("Done")

			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format("02 Jan 15:04")
			}
		}

		tbl.AddRow(
			strconv.Itoa(index+1),
			t.Title,
			status,
			t.CreatedAt.Format("02 Jan 15:04"),
			completedAt,
		)
	}

	fmt.Println()
	tbl.Render()
	fmt.Println()
}
