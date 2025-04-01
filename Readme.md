## VanMango - Van Management System based on Golang

## Features

- Security - Authentication
- API Versioning

## Project Structure

goserver
├── driver
│ └── driver.go
├── handler
│ ├── v1
│ │ ├── engine.go
│ │ └── van.go
│ ├── login.go
│ └── utils.go
├── middleware
│ └── auth.go
├── models
│ ├── engine.go
│ ├── login.go
│ └── van.go
├── service
│ ├── engine.go
│ ├── interface.go
│ └── van.go
├── store
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
