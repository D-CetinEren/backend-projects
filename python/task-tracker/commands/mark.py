import argparse
import time
import sys

sys.path.insert(0, '/home/cheren/Desktop/backend-projects/python/task-tracker/functionality')

from storage import read_tasks, write_tasks

def mark_task(task_id, status):
    tasks = read_tasks()
    for task in tasks:
        if task.id == task_id:
            task.status = status
            task.updated_at = time.strftime("%Y-%m-%d %H:%M:%S")
            write_tasks(tasks)
            print(f"Task {task_id} marked as {status}.")
            return
    print(f"Task with ID {task_id} not found.")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Mark a task as done or in progress")
    parser.add_argument("id", type=int, help="Task ID")
    parser.add_argument("status", choices=["done", "in-progress"], help="New task status")
    args = parser.parse_args()
    mark_task(args.id, args.status)
