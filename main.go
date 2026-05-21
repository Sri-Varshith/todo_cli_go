package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("task1")
	todos.add("task2")
	fmt.Println(todos)
	todos.delete(0)
	fmt.Println(todos)

}
