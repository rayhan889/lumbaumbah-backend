# 🐬 Lumba Umbah (Backend) - A Laundry Service

## 📖 Description

This is a backend service for the Lumba Umbah (laundry) application. It provides a RESTful API for managing laundry requests, laundry types, and addresses. The service is built using the Go programming language and utilizes the Gin web framework for handling HTTP requests and responses. It also uses the Gorm ORM for database interactions.

## ✅ Features

- 🚀 Manage laundry requests, laundry types, and addresses
- 🛠️ Implement user authentication and authorization
- 📊 Track laundry request status and completion dates
- 🔐 Secure API endpoints with JWT authentication
- 💪 Implement unit tests for the service
- 📦 Package the service as a Docker image

## 🛠️ Tech Stack

- Go
- Gin
- Gorm
- JWT
- PostgreSQL
- Docker

## 🤔 Why do use layering pattern as my choice of architecture?

Layering pattern is a design pattern that allows me to separate the application into layers, each responsible for a specific aspect of the application. In this case, the layers are:

1. Hanlder Layer: This layer contains the endpoint handlers for the API. It handles the incoming requests and delegates the processing to the appropriate service layer.
2. Service Layer: This layer contains the business logic for the application.
3. Store(Repository) Layer: This layer contains the data access layer for the application. It interacts with the database using the Gorm ORM.

With this pattern...

- I can easily mock a repo or service when testing a layer above it and of course it makes easier to write unit tests.
- Each layer has a clear job, such as handling requests, processing data, or interacting with the database.
- This pattern also widely used in the industry, such as the microservices architecture.
