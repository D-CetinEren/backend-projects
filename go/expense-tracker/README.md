# Expense Tracker

A simple CLI-based application written in Go to help you manage your expenses efficiently. This application allows you to add, update, delete, and view expenses while also providing summaries for total or monthly expenses.  

## Features

- **Add Expenses**: Add an expense with a description, amount, and category.  
- **Update Expenses**: Update an existing expense.  
- **Delete Expenses**: Remove an expense by its ID.  
- **List Expenses**: View all expenses with details like ID, date, description, and amount.  
- **Summarize Expenses**: Get a summary of all expenses or filter by a specific month.  
- **Categories**: Assign categories to your expenses.  
- **Budget Warning** (Optional): Set monthly budgets and receive warnings when exceeding them.  
- **Export to CSV**: Export your expenses to a CSV file for easy sharing or backup.  

---

## Installation

1. Clone the repository:  
   ```bash
   git clone https://github.com/your-username/expense-tracker.git
   cd expense-tracker
   ```

2. Install dependencies:  
   ```bash
   go mod tidy
   ```

3. Build the application:  
   ```bash
   go build -o expense-tracker
   ```

4. Run the application:  
   ```bash
   ./expense-tracker
   ```

---

## Usage

### Adding an Expense
```bash
expense-tracker add --description "Lunch" --amount 20 --category "Food"
# Output: Expense added successfully (ID: 1)
```

### Listing Expenses
```bash
expense-tracker list
# Output:
# ID   Date        Description   Amount  Category
# 1    2025-01-24  Lunch         $20.00  Food
```

### Deleting an Expense
```bash
expense-tracker delete --id 1
# Output: Expense deleted successfully
```

### Viewing a Summary
```bash
expense-tracker summary
# Output: Total expenses: $20.00
```

### Viewing Monthly Summary
```bash
expense-tracker summary --month 1
# Output: Total expenses for January: $20.00
```

---

## Directory Structure

```
expense-tracker/
├── cmd/               # CLI commands using Cobra
│   ├── root.go        # Root command
│   ├── add.go         # Add expense command
│   ├── list.go        # List expenses command
│   ├── delete.go      # Delete expense command
│   ├── summary.go     # Summary command
├── internal/          
│   ├── models/        # Data models (e.g., Expense struct)
│   ├── repository/    # File handling and storage logic
│   ├── services/      # Business logic for managing expenses
├── data/              # JSON file storage for expenses
├── main.go            # Entry point of the application
├── go.mod             # Go module dependencies
├── go.sum             # Go module checksums
```

---

## Future Enhancements

- [ ] Add support for recurring expenses.  
- [ ] Support for more advanced filtering (e.g., by date range, category).  
- [ ] Implement encryption for sensitive data.  
- [ ] Allow database storage using SQLite for improved scalability.  
- [X] Add unit tests for better reliability.  

---