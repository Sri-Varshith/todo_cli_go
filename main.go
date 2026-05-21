package main

// import "fmt"

func main() {
	todos := Todos{}
	todos.add("task1")
	todos.add("task2")
	todos.toggle(0)
	todos.print()

}
