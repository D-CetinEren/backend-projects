import argparse
from functionality import storage
from functionality import task


def add_task(description):
    tasks = storage.read_tasks()
    new_task = task.Task(len(tasks) + 1, description)
    tasks.append(new_task)
    storage.write_tasks(tasks)
    print(f"Task added successfully (ID: {new_task.id})")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Add a new task")
    parser.add_argument("description", type=str, help="Task description")
    args = parser.parse_args()
    add_task(args.description)
