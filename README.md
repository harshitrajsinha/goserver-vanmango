## VanMango - Van Management System based on Golang

## Features

- Security - Authentication
- API Versioning

## Project Structure

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
