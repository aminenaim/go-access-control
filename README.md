# go-access-control

Exploring access control implementations in Go.

## Getting started

This project demonstrates a scope-based access control implementation for a RESTful API built with Go and the Gin web framework. The core functionality is defined in the `services/core` package, while the API routes and handlers are defined in the `services/gateway` package.

### Core Package

The `core` package defines the following structs and types:

-   `HttpMethod`: a string representing the HTTP methods (GET, POST, PUT, DELETE, etc.).
-   `Role`: a struct representing a user role with a name and a list of scopes (permissions).
-   `Route`: a struct defining an API route with its path, HTTP method, handler function, and required scope.

-   `AuthHandler`: a struct that handles authentication and authorization based on user roles and scopes.

The `AuthHandler` struct contains the following fields:

-   `Routes`: a slice of API routes to be protected.
-   `Roles`: a slice of user roles and their associated scopes.
-   `IdentityProvider`: an implementation of the `IdentityProvider` interface, which provides user information and roles.

The `AuthMiddleware` function in the `AuthHandler` struct is a Gin middleware that checks if the user has the required scope for the requested API route.

If the user is not authenticated or doesn't have the required scope, the middleware returns a 401 Unauthorized or 403 Forbidden response.

### Gateway Package

The `gateway` package contains the main entry point of the application. It defines two sample user roles (`User` and `Admin`) with different scopes, and a list of API routes (`routes`) with their respective paths, methods, handlers, and required scopes.

The `main` function sets up the Gin server, creates an instance of the `AuthHandler` with the defined routes and roles, and registers the `AuthMiddleware` as a global middleware for the server. It then adds the routes to the server using the `AddRoutes` function and starts the server on port 8080.

## Usage

1. Start the server by running `go run services/gateway/main.go`.
2. The API will be available at `http://localhost:8080/api/v1`.
3. You can test the API routes with tools like Postman or curl (a basic test.sh is provided [here](test.sh)), passing the required authentication headers or credentials based on the user roles and scopes defined in the code.

> 🚩 Note: This is a simplified example to demonstrate scope-based access control in Go. In a production environment, you would likely use a more robust authentication and authorization mechanism, such as JSON Web Tokens (JWT) or OAuth2, and integrate with a real identity provider.
