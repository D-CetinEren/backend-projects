# Github-user-activity

**Github-user-activity** is a command-line application (CLI) built with Go, designed to fetch and display the recent activity of a GitHub user using the GitHub API. This tool is helpful for developers and GitHub users who want a quick overview of any user's recent GitHub activities directly in their terminal.

---

## **Directory Structure**

```
github-user-activity/
├── cmd/
│   ├── root.go         # Root command
│   └── activity.go     # 'activity' subcommand
├── internal/
│   ├── api.go          # Handles GitHub API communication
│   ├── formatter.go    # Formats the output
│   └── activity.go     # Logic to display activities
├── main.go             # Entry point for the application
├── go.mod              # Go module file
└── go.sum              # Dependency file
```

---

## **Features**

- Fetches recent activity of a GitHub user using the GitHub API.
- Displays activity details in the terminal, such as:
  - Pushes to repositories
  - Starred repositories
  - Opened issues
- Handles errors gracefully, including invalid usernames and API failures.
- Built using the Cobra CLI library, adhering to SOLID principles.

---

## **Installation**

### **Prerequisites**
- Go 1.23 or later

### **Steps to Install**

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/github-user-activity.git
   cd github-user-activity
   ```

2. Build the CLI:
   ```bash
   go build -o github-user-activity
   ```

3. Run the application:
   ```bash
   ./github-user-activity activity <username>
   ```

---

## **Usage**

### **Commands**

#### 1. Display Help
```bash
./github-user-activity --help
```

#### 2. Fetch Recent Activity
```bash
./github-user-activity activity <username>
```
- Replace `<username>` with the desired GitHub username.

**Example**:
```bash
./github-user-activity activity octocat
```

**Output**:
```
Recent activity for GitHub user 'octocat':
- Pushed to repository 'octocat/Spoon-Knife'
- Starred repository 'octocat/Hello-World'
- Opened an issue in 'octocat/test-repo'
```

---

## **Development**

### **Testing**

- Write unit tests for each component using Go's `testing` package.
- Mock API responses for testing.

### **Contributing**

Contributions are welcome! Please submit a pull request or file an issue for feature requests and bug reports.

---

## **To-Do List**

### **Basic Enhancements**
- [X] Handle GitHub API rate limits with user feedback.
- [X] Add filtering options for event types (e.g., push, issue, star).
- [X] Paginate results to fetch more than 30 events.

### **Intermediate Features**
- [ ] Support JSON/YAML output for easy integration with other tools.
- [X] Implement local caching to reduce API calls.
- [ ] Allow fetching activity for multiple users simultaneously.

### **Advanced Features**
- [ ] Add a repository insights command (e.g., `repos <username>`).
- [ ] Implement real-time activity monitoring (`--watch` mode).
- [ ] Support date filtering for activity (e.g., `--since`, `--until`).
- [ ] Integrate GitHub authentication via Personal Access Tokens.
- [ ] Add contribution analysis (e.g., total commits, pull requests).

---

## **License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---


