import argparse
import sys

sys.path.insert(0, '/home/cheren/Desktop/backend-projects/python/task-tracker/functionality')

from storage import read_tasks, write_tasks

def delete_task(task_id):
    tasks = read_tasks()
    tasks = [task for task in tasks if task.id != task_id]
    write_tasks(tasks)
    print(f"Task {task_id} deleted successfully.")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Delete a task")
    parser.add_argument("id", type=int, help="Task ID")
    args = parser.parse_args()
    delete_task(args.id)
