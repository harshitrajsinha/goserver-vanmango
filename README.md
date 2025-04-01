## VANMANGO Server

🚀 A **Van Management System** built using **Golang** with **PostgreSQL**, enabling advanced CRUD operations following standard REST API practices.

## ✨Features

- 🔗 **Dependency Injection** for modularity
- 🔒 **JWT Authentication** for security
- ⌛ **API Versioning** for backward compatibility
- 🛠 Robust **error handling** for debugging
- 💾 **Persistant storage** using PostgreSQL

## 📌 Tech Stack

- 🌐 **Server:** Golang
- 💾 **Database:** PostgreSQL
- 🚀 **Deployment:** Neon

## 🔌API Endpoints

### Van

| Method | Endpoint          | Description            |
| ------ | ----------------- | ---------------------- |
| GET    | `/api/v1/van/:id` | Get van using ID       |
| GET    | `/api/v1/vans`    | Get all vans           |
| POST   | `/api/v1/van`     | Create a new van       |
| PUT    | `/api/v1/van/:id` | Update an van          |
| PATCH  | `/api/v1/van/:id` | Update sub-part of van |
| DELETE | `/api/v1/van/:id` | Delete an van          |

### Engine

| Method | Endpoint             | Description               |
| ------ | -------------------- | ------------------------- |
| GET    | `/api/v1/engine/:id` | Get engine using ID       |
| GET    | `/api/v1/engines`    | Get all engines           |
| POST   | `/api/v1/engine`     | Create a new engine       |
| PUT    | `/api/v1/engine/:id` | Update an engine          |
| PATCH  | `/api/v1/engine/:id` | Update sub-part of engine |
| DELETE | `/api/v1/engine/:id` | Delete an engine          |

## 📂 Project Structure

```
goserver/
├── driver       # Database initialization
│ └── driver.go
├── handler      # API route handling
│ ├── v1
│ │ ├── engine.go
│ │ └── van.go
│ ├── login.go
│ └── utils.go
├── middleware   # API authorization
│ └── auth.go
├── models
│ ├── engine.go
│ ├── login.go
│ └── van.go
├── service
│ ├── engine.go
│ ├── interface.go
│ └── van.go
├── store           # CRUD operation
│ ├── engine.go
│ ├── interface.go
│ ├── schema.sql
│ └── van.go
├── tmp
│ ├── runner-build.exe
│ └── runner-build.exe~
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── goserver-vanmango.exe
├── main.go
├── pending.md
└── README.md
```

<!-- <details>
  <summary>Get all Vans</summary>

   ### Request

```curl
/api/v1/engines
```

### Response

```go
    {
    "code": 200,
    "data": [
        {
            "id": "e1f86b1a-0873-4c19-bae2-fc60329d0140",
            "displacement_in_cc": 2000,
            "no_of_cylinders": 4,
            "material": "aluminium"
        },
        {
            "id": "f4a9c66b-8e38-419b-93c4-215d5cefb318",
            "displacement_in_cc": 1600,
            "no_of_cylinders": 4,
            "material": "iron"
        },
        ]
    }
```

</details> -->
