import json
import time

class Task:
    def __init__(self, task_id, description, status="todo", created_at=None, updated_at=None):
        self.id = task_id
        self.description = description
        self.status = status
        self.created_at = created_at or time.strftime("%Y-%m-%d %H:%M:%S")
        self.updated_at = updated_at or self.created_at

    def to_dict(self):
        return {
            "id": self.id,
            "description": self.description,
            "status": self.status,
            "created_at": self.created_at,
            "updated_at": self.updated_at,
        }

    @staticmethod
    def from_dict(data):
        return Task(
            data["id"],
            data["description"],
            data["status"],
            data["created_at"],
            data["updated_at"],
        )
