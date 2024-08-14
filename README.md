# Simple Auth Service

This is a simple authentication service built with Go, providing three key endpoints: `/login`, `/refresh-token`, and `/profile`.

## Endpoints

- **POST /login**
    - Description: Authenticate a user and return a JWT token.
    - Request Body: `{"email": "user@example.com", "password": "password"}`
    - Response: JSON object containing access and refresh tokens.

- **POST /refresh-token**
    - Description: Refresh the JWT token using a refresh token.
    - Request Body: `{"refresh_token": "your_refresh_token_here"}`
    - Response: JSON object containing a new access token and refresh token.

- **GET /profile**
    - Description: Retrieve the profile information of the authenticated user.
    - Authentication: Requires a valid JWT token.
    - Response: JSON object containing user profile details.

## Setup

Follow these steps to set up and run the project:

1. **Clone the Repository**

   ```bash
   git clone https://github.com/okyirmawan/simple-auth-service.git
   cd your_repository

2. **Edit .env File**

    ```bash
    JWT_SECRET_KEY=your_secret_key_here
    DB_USERNAME=user
    DB_PASSWORD=password
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_DATABASE=db_name

3. **Create the Databasee**

    ```bash
   CREATE DATABASE db_name;

4. **Run Migrations**
    
    ```bash
    goose -dir database/migrations mysql "user:password@tcp(localhost:3306)/db_name" up

5. **Seed the Database**

    ```bash
    go run main.go -seed

6. **Run the Application**

    ```bash
    go run main.go

## Dependencies

- Go 1.22.6
- Gin - Web framework
- GORM - ORM for Go
- Goose - Database migrations tool