import argparse
from functionality import storage

def delete_task(task_id):
    tasks = storage.read_tasks()
    tasks = [task for task in tasks if task.id != task_id]
    storage.write_tasks(tasks)
    print(f"Task {task_id} deleted successfully.")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Delete a task")
    parser.add_argument("id", type=int, help="Task ID")
    args = parser.parse_args()
    delete_task(args.id)
