# Github-user-activity

`Github-user-activity` is a CLI tool written in Go that allows you to fetch, filter, and display recent GitHub activity for a specified user. It supports features like caching, pagination, filtering by event types, and exporting results in different formats (JSON, YAML, or plain text).

## Features

- **Caching**: Reduces the number of API calls by storing responses locally.
- **Pagination**: Fetches more than 30 events by iterating over multiple pages.
- **Event Filtering**: Filters events by type (e.g., `PushEvent`, `IssuesEvent`, `WatchEvent`).
- **Flexible Output**: Supports output in plain text, JSON, or YAML format.
- **File Export**: Option to save the output to a file.

## Prerequisites

- **Go**: Version 1.23 or higher
- **GitHub Personal Access Token (PAT)**: Required for making authenticated API requests to avoid rate limiting.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/D-CetinEren/Github-user-activity.git
   cd Github-user-activity
   ```

2. Build the project:
   ```bash
   go build -o github-user-activity
   ```

3. Run the binary:
   ```bash
   ./github-user-activity
   ```

## Usage

### Fetch Activity

Basic command to fetch a user's activity:
```bash
./github-user-activity activity <username>
```

### Options

| Flag             | Description                                                                 |
|------------------|-----------------------------------------------------------------------------|
| `--cache-ttl`    | Cache time-to-live in minutes (default: 10).                                |
| `--max-pages`    | Maximum number of pages to fetch (default: 1).                              |
| `--event-type`   | Filter events by type (e.g., `PushEvent`, `IssuesEvent`).                    |
| `--output`       | Output format: `text` (default), `json`, or `yaml`.                         |
| `--output-file`  | File to write the output.                                                   |

### Examples

1. Fetch the latest activity for user `octocat`:
   ```bash
   ./github-user-activity activity octocat
   ```

2. Fetch activity with caching:
   ```bash
   ./github-user-activity activity octocat --cache-ttl 30
   ```

3. Fetch and filter events by type:
   ```bash
   ./github-user-activity activity octocat --event-type PushEvent
   ```

4. Save output to a file in JSON format:
   ```bash
   ./github-user-activity activity octocat --output json --output-file activity.json
   ```

5. Fetch multiple pages of activity:
   ```bash
   ./github-user-activity activity octocat --max-pages 5
   ```

## Directory Structure

```
Github-user-activity/
├── cmd/                  # CLI commands
│   ├── root.go           # Root command
│   └── activity.go       # 'activity' command implementation
├── internal/             # Internal packages
│   ├── api/              # GitHub API client and caching logic
│   ├── cache/            # Local caching implementation
│   ├── filters/          # Filtering logic
│   └── formatter/        # Formatting output
├── models/               # Data models for GitHub events
├── go.mod                # Go module file
├── go.sum                # Dependency checksum file
└── main.go               # Main entry point
```

---

## **To-Do List**

### **Basic Enhancements**
- [X] Handle GitHub API rate limits with user feedback.
- [X] Add filtering options for event types (e.g., push, issue, star).
- [X] Paginate results to fetch more than 30 events.

### **Intermediate Features**
- [X] Support JSON/YAML output for easy integration with other tools.
- [X] Implement local caching to reduce API calls.
- [X] Allow fetching activity for multiple users simultaneously.

### **Advanced Features**
- [ ] Add a repository insights command (e.g., `repos <username>`).
- [ ] Implement real-time activity monitoring (`--watch` mode).
- [ ] Support date filtering for activity (e.g., `--since`, `--until`).
- [ ] Integrate GitHub authentication via Personal Access Tokens.
- [ ] Add contribution analysis (e.g., total commits, pull requests).

---
## **Contributing**

Contributions are welcome! Feel free to submit a pull request or open an issue for suggestions and bug reports.

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---


