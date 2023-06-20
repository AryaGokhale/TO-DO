# API Server in Golang
This is a REST API server implemented in Golang that provides various endpoints for user signup, login, and note management. It utilizes the Gin HTTP web framework and JWT tokens for session identification.

# Endpoints
The server implements the following endpoints:

Signup user: Allows users to create an account by providing their credentials.
Login user: Enables users to authenticate themselves and obtain a JWT token for subsequent requests.
Create notes: Allows authenticated users to create new notes.
Read notes: Retrieves existing notes for an authenticated user.
Delete notes: Enables users to delete their existing notes.

# Tech Stack
The API server is built using the following technologies:

Gin: A lightweight HTTP web framework for building RESTful APIs in Golang. It provides routing, middleware, and other essential functionalities.
JWT Token: JSON Web Tokens are used for creating unique tokens that identify each user session (sessionID). These tokens are used for authentication and authorization purposes.
Postman: REST API testing is performed using Postman, a popular API development and testing tool.
Swagger: Swagger will be used for further testing and documentation of the API. It provides a user-friendly interface to explore and interact with the API endpoints.

The docker-compose.yml file is provided to simplify the process of running the server in a container using Docker Compose. It defines a service called app, which is built from the Docker image specified in the Dockerfile. The service maps port 8080 of the container to port 8080 of the host machine.
