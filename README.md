
# 📝 TaskAppSkillsRock


## 🚀 Running the Application

### With Docker
```bash
docker-compose up -d
```

### Locally
```bash
go run cmd/main.go
```

---

## ⚙️ Environment Variables

Set up your `.env` file with the following variables:

| Variable       | Description          | Example              |
|----------------|----------------------|----------------------|
| `APP_NAME`     | Name of the app      | `TaskAppSkillsRock`  |
| `APP_PORT`     | App server port      | `3000`               |
| `DB_HOST`      | Database host        | `localhost`          |
| `DB_PORT`      | Database port        | `5432`               |
| `DB_USER`      | Database user        | `postgres`           |
| `DB_PASSWORD`  | Database password    | `postgres`           |
| `DB_NAME`      | Database name        | `taskdb`             |

---

## 📚 API Documentation

### ✅ Get All Tasks

**Endpoint:**  
`GET /api/tasks`

**Response Example:**
```json
[
  {
    "id": 1,
    "title": "Task 1",
    "description": "Description 1",
    "status": "done",
    "created_at": "2023-06-01T10:00:00Z",
    "updated_at": "2023-06-01T10:00:00Z"
  },
  {
    "id": 2,
    "title": "Task 2",
    "description": "Description 2",
    "status": "in_progress",
    "created_at": "2023-06-02T10:00:00Z",
    "updated_at": "2023-06-02T10:00:00Z"
  }
]
```

---

### ➕ Create a Task

**Endpoint:**  
`POST /api/tasks`

**Request Body:**
```json
{
  "title": "New Task",
  "description": "Description of the new task",
  "status": "new"
}
```

---

### ✏️ Update a Task

**Endpoint:**  
`PUT /api/tasks/:id`

**Request Body:**
```json
{
  "title": "updated title",
  "description": "updated description",
  "status": "in_progress"
}
```

---

### ❌ Delete a Task

**Endpoint:**  
`DELETE /api/tasks/:id`

---
