# Task Tracker CLI

## Description

Task Tracker is a Command Line Interface (CLI) application that helps users manage and track tasks efficiently. The application allows adding, updating, deleting, and listing tasks, as well as marking tasks with different statuses such as "todo," "in-progress," and "done."

## Features

- **Add Task**: Add a new task with a description.
- **Update Task**: Modify the description of an existing task.
- **Delete Task**: Remove a task by its ID.
- **Mark Task**: Update the status of a task to "in-progress" or "done."
- **List Tasks**:
  - List all tasks.
  - List tasks by their status: "todo," "in-progress," or "done."
  - Filter tasks by creation date range using `--start-date` and `--end-date`.
- **Task Properties**:
  - Unique ID
  - Description
  - Status (todo, in-progress, done)
  - Creation Date
  - Last Updated Date

## Installation

### Prerequisites

- Go 1.23 or later installed.
- Git installed.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/D-CetinEren/backend-projects.git
   ```
2. Navigate to the project directory:
   ```bash
   cd backend-projects/go/task-tracker
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Build the project:
   ```bash
   go build -o task-tracker
   ```
5. Run the CLI:
   ```bash
   ./task-tracker
   ```

## Usage

Below are examples of how to use the CLI commands:

### Add a Task

```bash
./task-tracker add "Buy groceries"
```

### Update a Task

```bash
./task-tracker update 1 "Buy groceries and cook dinner"
```

### Delete a Task

```bash
./task-tracker delete 1
```

### Mark a Task

- Mark as in-progress:
  ```bash
  ./task-tracker mark-in-progress 1
  ```
- Mark as done:
  ```bash
  ./task-tracker mark-done 1
  ```

### List Tasks

- List all tasks:
  ```bash
  ./task-tracker list
  ```
- List tasks by status:
  ```bash
  ./task-tracker list --status todo
  ./task-tracker list --status in-progress
  ./task-tracker list --status done
  ```
- List tasks created after a specific date:
  ```bash
  ./task-tracker list --start-date 2025-01-01
  ```
- List tasks created before a specific date:
  ```bash
  ./task-tracker list --end-date 2025-01-10
  ```
- List tasks within a date range:
  ```bash
  ./task-tracker list --start-date 2025-01-01 --end-date 2025-01-10
  ```

## Directory Structure

```plaintext
📂 task-tracker/
│
├── cmd/            # Command definitions for CLI
│   ├── root.go
│   ├── add.go
│   ├── update.go
│   ├── delete.go
│   ├── list.go
│   ├── mark.go
│
├── internal/       # Internal packages for task logic
│   └── task/
│       ├── task.go
│       └── storage.go
│
├── logs/           # Log files directory
│   └── task-tracker.log
│
├── go.mod          # Go module file
├── go.sum          # Go dependencies
└── main.go         # Entry point
```

## Logging

Logs are stored in the `logs/task-tracker.log` file. Each operation is logged with a timestamp and relevant details.

## To-Do List

Here are some potential features to enhance the project:

- [ ] Unit Tests
- [X] UUID for task IDs
- [X] Filtering by date

## License

This project is licensed under the MIT License. See the LICENSE file for details.

https://roadmap.sh/projects/task-tracker

