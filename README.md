# Bookstore API

## Table of Contents
1. [Project Description](#project-description)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Setup Instructions](#setup-instructions)
5. [Architecture Overview](#architecture-overview)
6. [Endpoints](#endpoints)
7. [Configuration](#configuration)
8. [Contributing](#contributing)
9. [License](#license)

---

## Project Description

The **Bookstore API** is a RESTful API built using Go that allows users to manage books, authors, and user accounts. It provides endpoints for creating, reading, updating, and deleting books, as well as managing user authentication and authorization.

This project follows a modular architecture with clear separation of concerns, making it easy to maintain and extend.

---

## Features

- **User Management**: Register, login, and manage user profiles.
- **Book Management**: Add, update, delete, and retrieve books.
- **Author Management**: Associate books with authors.
- **Authentication**: Secure endpoints with JWT-based authentication.
- **Pagination**: Support for paginated responses for large datasets.
- **Error Handling**: Comprehensive error handling for all API endpoints.

---

## Prerequisites

Before running the application, ensure you have the following installed:

- **Go**: Version 1.18 or higher.
- **MongoDB**: For storing data.
- **Docker (Optional)**: To run MongoDB in a container.
- **Postman or cURL**: For testing API endpoints.

---

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/bookstore-api.git
cd bookstore-api
```

## Setup Instructions

### 1. Clone the Repository

```
git clone https://github.com/yourusername/bookstore-api.git
cd bookstore-api
```

### 2. Install Dependencies

Run the following command to install the required Go packages:

```
go mod tidy
```

### 3. Set Up Environment Variables

Copy the example configuration file and modify it as needed:

```
cp .env.example .env
```

Edit `.env` to configure database connection details, JWT secret, and other settings.

Example `.env` file:

```
PORT=8080
MONGODB_URI=mongodb://localhost:27017
JWT_SECRET=your_jwt_secret_key
```

### 4. Start MongoDB

If you're using Docker, start MongoDB with the following command:

```
docker run --name mongodb -p 27017:27017 -d mongo
```

Alternatively, start MongoDB locally if it's already installed on your system.

### 5. Run the Application

Start the API server by running:

```
go run cmd/main.go
```

The server will start on `http://localhost:8080`.