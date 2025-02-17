# Task CLI

A lightweight, efficient command-line task manager built with Golang and BoltDB. Manage your daily tasks with simple commands right from your terminal.

## Features

* Add new tasks with descriptions and optional due dates
* List all pending tasks in an organized format
* Mark single or multiple tasks as completed
* View completed tasks filtered by today's date
* Permanently delete tasks when no longer needed
* Persistent storage using BoltDB for reliable data management

## Installation

Ensure you have Go 1.16 or later installed on your system.

```bash
# Clone the repository
git clone https://github.com/yourusername/task-cli.git
cd task-cli

# Install dependencies
go mod tidy

# Build the CLI
go build -o task

# Optional: Add to your PATH for system-wide access
sudo mv task /usr/local/bin/

```

## Usage

### Add Task
To add a new task, run:
```bash
./task add "Finish Golang project"
```
### List Pending Tasks
To list all pending tasks, run:
```bash
./task list
```
### Mark Task as Completed
To mark a task as completed, run:
```bash
./task do 1 2
```
### View Completed Tasks (Today)
To view all tasks completed today, run:
```bash
./task do 1 2
```
### Delete a Task Permanently
To delete a task permanently, run:
```bash
./task rm 3
```

## Project Structure

```
task-cli/
├── cmd/
│   ├── add.go         # Handle task addition
│   ├── list.go        # Display task listings
│   ├── do.go          # Mark tasks complete
│   ├── completed.go   # Show completed tasks
│   ├── rm.go          # Remove tasks
│   └── root.go        # Root command configuration
├── db/
│   └── tasks.go       # Database operations
├── main.go            # Application entry point
├── go.mod            # Go module definition
├── go.sum            # Module dependency checksums
└── README.md         # Project documentation
```
## Future Scope

Here are some potential features to enhance the Task CLI:

- **Task Filtering**: Add flags to filter tasks based on criteria such as status, priority, or due date.

- **Task Due Dates**: Allow users to set due dates for tasks and list tasks that are due soon.

- **Task Search**: Implement a search command to find tasks by name, ID, or description.

- **Task Priority**: Allow users to set priorities for tasks (e.g., high, medium, low) and sort tasks based on priority.

## Contributing
Pull requests are welcome! Feel free to open an issue for suggestions or bugs.
