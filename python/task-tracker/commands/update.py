import argparse
import time
import sys

sys.path.insert(0, '/home/cheren/Desktop/backend-projects/python/task-tracker/functionality')

from storage import read_tasks, write_tasks

def update_task(task_id, new_description):
    tasks = read_tasks()
    for task in tasks:
        if task.id == task_id:
            task.description = new_description
            task.updated_at = time.strftime("%Y-%m-%d %H:%M:%S")
            write_tasks(tasks)
            print(f"Task {task_id} updated successfully.")
            return
    print(f"Task with ID {task_id} not found.")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Update a task description")
    parser.add_argument("id", type=int, help="Task ID")
    parser.add_argument("description", type=str, help="New task description")
    args = parser.parse_args()
    update_task(args.id, args.description)
