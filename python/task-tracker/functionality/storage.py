import json
import os
from task import Task

TASKS_FILE = "tasks.json"

def read_tasks():
    if not os.path.exists(TASKS_FILE):
        return []
    with open(TASKS_FILE, "r") as f:
        try:
            data = json.load(f)
            return [Task.from_dict(task) for task in data]
        except json.JSONDecodeError:
            return []

def write_tasks(tasks):
    with open(TASKS_FILE, "w") as f:
        json.dump([task.to_dict() for task in tasks], f, indent=2)
