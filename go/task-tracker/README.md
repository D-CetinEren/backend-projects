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
- **Configuration Management**: Flexible configuration using YAML files
- **Structured Error Handling**: Custom error types for better error management
- **Input Validation**: Comprehensive validation for task operations
- **Graceful Shutdown**: Proper signal handling and cleanup
- **Logging**: Detailed logging with timestamps and relevant details
- **Unit Tests**: Comprehensive test cases for all commands

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

## Configuration

Task Tracker supports flexible configuration using YAML files. Configuration allows you to customize settings such as storage paths and log file locations.

### Default Configuration

If no configuration file is provided, Task Tracker uses default settings:

- **Storage Path**: Current directory (`./tasks.json`)

- **Log Path**: `./logs/task-tracker.log`

### Custom Configuration

1. **Create a configuration file** (e.g., `config.yaml`) in the project directory:
   ```yaml
   storage_path: "./data/tasks.json"
   log_path: "./logs/task-tracker.log"
   ```

2. **Run the CLI with the configuration file:**

   Modify the `main.go` if necessary to accept a configuration file path or place the `config.yaml` in the default configuration directory.

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
│   ├── add_test.go        # Unit tests for add command
│   ├── update_test.go     # Unit tests for update command
│   ├── delete_test.go     # Unit tests for delete command
│   ├── list_test.go       # Unit tests for list command
│   ├── mark_test.go       # Unit tests for mark command
│
├── internal/       # Internal packages for task logic
│   ├── config/     # Configuration handling
│   │   ├── config.go
│   │   └── config_test.go
│   └── task/
│       ├── task.go
│       ├── storage.go
│       └── task_test.go    # Unit tests for task package
│
├── logs/           # Log files directory
│   └── task-tracker.log
│
├── configs/        # Configuration files (YAML)
│   └── config.yaml
│
├── go.mod          # Go module file
├── go.sum          # Go dependencies
├── main.go         # Entry point
├── README.md       # Project documentation
└── LICENSE         # License information
```

## Logging


Logs are stored in the `logs/task-tracker.log` file. Each operation is logged with a timestamp and relevant details, aiding in debugging and monitoring the application's behavior.

## To-Do List

Here are some potential features to enhance the project:

- [ ] Add more advanced filters (e.g., by priority, tags)
- [ ] Implement a web interface for managing tasks
- [ ] Support exporting tasks to CSV or JSON file
- [ ] Add integration with notification systems
- [ ] Add database support instead of file storage
- [ ] Implement backup functionality
- [ ] Add user authentication and multi-user support
- [X] Unit Tests: Added test cases for all major commands
- [X] Filtering by date
- [X] Configuration management
- [X] Structured error handling
- [X] Input validation
- [X] Graceful shutdown

## License

This project is licensed under the MIT License. See the LICENSE file for details.

https://roadmap.sh/projects/task-tracker

