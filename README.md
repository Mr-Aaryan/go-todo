# Todo CLI

A simple and intuitive command-line todo application built with Go, featuring interactive prompts and a clean interface.

## Features

- ✅ Add new todo items
- 📋 View all todos
- 🗑️ Delete completed todos
- 🎯 Interactive command-line interface
- 💾 Persistent storage

## Built With

- [Go](https://golang.org/) - Programming language
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [PromptUI](https://github.com/manifoldco/promptui) - Interactive prompts

## Installation

### Prerequisites

- Go 1.19 or higher

### Installation

```bash
git clone https://github.com/Mr-Aaryan/go-todo.git
cd todo-cli
go build
go install
```

## Usage

### View all todos

Display your current todo list:

```bash
todo
```

### Add a new todo

Add a new item to your todo list:

```bash
todo add
```

This will prompt you to enter the todo item description interactively.

### Delete a todo

Remove a completed or unwanted todo:

```bash
todo delete
```

This will display your current todos and allow you to select which one to delete.

## Examples

```bash
# View your todos
$ todo
📋 Your Todos:
1. Buy groceries
2. Walk the dog
3. Finish project report

# Add a new todo
$ todo add
✏️  New Task: Learn Go programming
✅ Todo added successfully!

# Delete a todo
$ todo delete
? 🗑️  Select a task to delete:
  👉 Buy groceries
    Walk the dog
    Finish project report
✅ Todo deleted successfully!
```

## File Structure

```
todo-cli/
├── cmd/
│   ├── root.go          # Root command setup
│   ├── add.go           # Add command implementation
│   └── delete.go        # Delete command implementation
├── internal/
│   └── storage/         # Data persistence layer
├── main.go              # Application entry point
├── go.mod               # Go module file
├── go.sum               # Go dependencies
└── README.md            # This file
```