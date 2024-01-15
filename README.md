# Bookstore API

## Overview
Implemented bookstore API Restful service
It built within Go, Gin, mysql, docker.

## Getting Started

### Prerequisites

- Go (version 1.19 or later) I have used 1.21.4
- Docker and Docker Compose
- MySQL Client (for local database interaction) -- optional



## Folder Structure

Here's an overview of the project's folder structure:

- `./main.go`: Entry point for this project.
- `/app` : All initialization functions are here (it will initialize database + migration, dependencies injection, routes, middleware and swagger) 
- `/internal`: Private application and library code.
    - `/domain`      : contains interfaces for repository and service layer
    -  `/dto`  : contains the struct to accept api requests  
    - `/handlers`: Contains api handlers that connect the routes to business logic.
    - `/helper` : added helper functions here
    - `/middleware`: Added API authentication middleware to check token
    - `/repositories`: Layer for database interaction.
    - `/services`: Business logic layer.
    - `/models`: Data models representing database tables.
    - `/utils` : Utility functions added here
- `/migrations`: Database migration file.
- `/tests`: Added integration tests. 
- `/docs`: Swagger documentation auto-generated files.
- `docker-compose.yml`: Docker Compose configuration for setting up the development environment.

### Setting Up the Development Environment

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/bookstore-api.git
   cd bookstore-api
2. **Start database and service with docker compose**
   ```bash 
   docker-compose up --build


### Running Tests

to run the tests, use the following command
   ```bash 
   go test ./tests/...
   ```

### Using the API
The API is now accessible at http://localhost:8080.
Use Swagger UI to interact with the API at http://localhost:8080/swagger/index.html.


### APIs flow 
Following is the API details 
- user registration API to create new user
- once user account gets created user need to login to get the authentication token
- user can fetch all the available books 
- for create order and get orders user need to provide valid auth token in headers 

### Assumptions 
- I have not implemented different currency support, Just assumed the book prices in the USD
- email verification is not implemented 
- Did not considered book inventory and order cancellation
- Payment integration is not implemented :)


