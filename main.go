package main

// import "fmt"

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("tasks.json")
	storage.Load(&todos)
	todos.add("task1")
	todos.add("task2")
	todos.toggle(0)
	todos.print()
	storage.Save(todos)

}
