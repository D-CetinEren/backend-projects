import argparse
import subprocess
import sys

def main():
    parser = argparse.ArgumentParser(description="Task Tracker CLI")
    subparsers = parser.add_subparsers(dest="command", required=True)

    subparsers.add_parser("add", help="Add a new task")
    subparsers.add_parser("update", help="Update a task description")
    subparsers.add_parser("delete", help="Delete a task")
    subparsers.add_parser("mark", help="Mark a task as done or in-progress")

    args, unknown = parser.parse_known_args()
    
    if args.command == "add":
        subprocess.run([sys.executable, "commands/add.py"] + unknown)
    elif args.command == "update":
        subprocess.run([sys.executable, "commands/update.py"] + unknown)
    elif args.command == "delete":
        subprocess.run([sys.executable, "commands/delete.py"] + unknown)
    elif args.command == "mark":
        subprocess.run([sys.executable, "commands/mark.py"] + unknown)

if __name__ == "__main__":
    main()
