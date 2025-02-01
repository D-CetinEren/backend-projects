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

- Python3  installed.
- Git installed.

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/D-CetinEren/backend-projects.git
   ```
2. Navigate to the project directory:
   ```bash
   cd backend-projects/python/task-tracker
   ```
3. Run the CLI:
   ```bash
   python3 task-tracker
   ```

## Usage

Below are examples of how to use the CLI commands:

### Add a Task

```bash
python3 task-tracker add "Buy groceries"
```

### Update a Task

```bash
python3 task-tracker update 1 "Buy groceries and cook dinner"
```

### Delete a Task

```bash
python3 task-tracker delete 1
```

### Mark a Task

- Mark as in-progress:
  ```bash
  python3 task-tracker mark-in-progress 1
  ```
- Mark as done:
  ```bash
  python3 task-tracker mark-done 1
  ```

## Directory Structure
```plaintext
📂 task-tracker/
│
├── commands/
│   ├── add.py
│   ├── update.py
│   ├── delete.py
│   └── mark.py
│
├── functionality/
│   ├── __init__.py
│   ├── storage.py
│   └── task.py
│
├── main.py
└── README.md
```

## To-Do List

Here are some potential features to enhance the project:

- [ ] UUID for task IDs
- [ ] Filtering by date

https://roadmap.sh/projects/task-tracker

