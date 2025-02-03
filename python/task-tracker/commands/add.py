import argparse
import sys

sys.path.insert(0, '/home/cheren/Desktop/backend-projects/python/task-tracker/functionality')

from storage import read_tasks, write_tasks
from task import Task

def add_task(description):
    tasks = read_tasks()
    new_task = Task(len(tasks) + 1, description)
    tasks.append(new_task)
    write_tasks(tasks)
    print(f"Task added successfully (ID: {new_task.id})")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Add a new task")
    parser.add_argument("description", type=str, help="Task description")
    args = parser.parse_args()
    add_task(args.description)
