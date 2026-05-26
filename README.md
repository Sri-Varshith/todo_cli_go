# Todo CLI

A clean and lightweight terminal-based todo manager built with Go.

## Features

- Add tasks
- Edit existing tasks
- Delete tasks
- Toggle task completion
- Persistent local storage using JSON
- Clean terminal table UI

---

## Installation

Clone the repository:

```bash
git clone https://github.com/Sri-Varshith/todo_cli_go
cd todo_cli_go
```

Build the executable:

```bash
go build -o mytodo
```

Run the application:

```bash
./mytodo -list
```

---

## Alternative Method

Install globally using Go:

```bash
go install
```

Then run from anywhere:

```bash
mytodo -list
```

---

## Usage

### Add a task

```bash
./mytodo -add "Learn Go"
```

### List tasks

```bash
./mytodo -list
```

### Toggle task completion

```bash
./mytodo -toggle 1
```

### Edit a task

```bash
./mytodo -edit 1:Build CLI project
```

### Delete a task

```bash
./mytodo -del 1
```

---

## Tech Stack

- Go
- `flag` package for CLI parsing
- `encoding/json` for persistence
- `github.com/aquasecurity/table` for terminal tables
- `github.com/fatih/color` for colored output

---

## Project Structure

```text
.
├── main.go
├── todos.go
├── storage.go
├── cmdflags.go
├── go.mod
└── tasks.json
```

---

## Future Improvements

- Priority levels
- Due dates
- Search and filtering
- Interactive TUI version using Bubble Tea
- Better command help output
- Sorting tasks

---

## License

MIT License
