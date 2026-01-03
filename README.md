# CloudSave 

CloudSave is a personal learning project focused on building a secure,
cloud-based file storage system using Go and web technologies.

**Note:** This project is currently under active development.

## About the Project

CloudSave is a web application where users can register,
authenticate, and securely store files in the cloud.

The project is primarily built to demonstrate:
- backend architecture in Go
- REST API design
- database integration
- authentication and security concepts

## Project Status

Under active development

Implemented so far:
- Project structure following Go best practices
- Environment-based configuration
- MySQL database connection
- User registration with hashed passwords
- REST API endpoint for user registration

Planned features:
- User login with JWT authentication
- File upload and download
- Access control and authorization
- Frontend integration

## Project Structure
```
cloudsave/
│
├── cmd/
│   └── server/                # Application entry point (starts the server)
│
├── internal/
│   ├── config/                # Application configuration and environment loading
│   ├── db/                    # Database connection and setup
│   ├── models/                # Data models (users, files, etc.)
│   ├── handlers/              # HTTP request handlers (controllers)
│   ├── middleware/            # Middleware (authentication, logging, etc.)
│   ├── services/              # Business logic layer
│   └── utils/                 # Helper and utility functions
│
├── frontend/                  # Frontend files (HTML, CSS, JavaScript)
│
├── storage/                   # Uploaded files (excluded from version control)
│
└── README.md                  # Project documentation
```