package main

// import "fmt"

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("tasks.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)

}
