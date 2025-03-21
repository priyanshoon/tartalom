# Tartalom

Tartalom is a simple headless content management system (CMS) built with Go (Golang). It provides a backend system to manage content, focusing on simplicity and efficiency.

## Features

- **User Authentication**: Supports Google OAuth 2.0 for user authentication.
- **Content Management**: Allows users to create, read, update, and delete blog posts.
- **JWT Security**: Utilizes JSON Web Tokens (JWT) for securing API endpoints.
- **Database Integration**: Integrates with MySQL using GORM for data persistence.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [MySQL](https://www.mysql.com/) database
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/priyanshoon/tartalom.git
   cd tartalom
   ```

2. **Switch to the development branch**:

   ```bash
   git checkout dev
   ```

3. **Set up environment variables**:

   Create a `.env` file in the root directory and add the following variables:

   ```env
   GOOGLE_CLIENT_ID=your_google_client_id
   GOOGLE_CLIENT_SECRET=your_google_client_secret
   JWT_SECRET=your_jwt_secret
   DATABASE_USER=your_database_user
   DATABASE_PASS=your_database_password
   DATABASE_HOST=localhost
   DATABASE_PORT=3306
   DATABASE_NAME=tartalom_db
   ```

4. **Install dependencies**:

   ```bash
   go mod tidy
   ```

5. **Run the application**:

   ```bash
   go run cmd/api/main.go
   ```

   The server will start on `http://localhost:6969`.

## Project Structure

```
tartalom/
├── cmd/
│   └── api/
│       └── main.go          # Entry point of the application
├── config/
│   └── config.go            # Configuration handling
├── database/
│   ├── connect.go        # Database connection setup
│   └── database.go                # Database instance
├── handler/
│   ├── auth_handler.go              # Authentication handlers
│   ├── blog_handler.go              # Blog management handlers
│   └── api.go
|   └── user_handler.go
├── middleware/
│   └── auth.go              # Authentication middleware
├── model/
│   ├── blog.go              # Blog model
│   └── user.go              # User model
├── route/
│   ├── auth_route.go              # Authentication routes
│   ├── user_route.go              # User routes
│   └── blog_route.go              # Blog routes
├── utils/
│   └── password_generator.go          # Utility functions
├── .env                     # Environment variables
├── .gitignore
├── go.mod
└── README.md
```

## API Endpoints

### Authentication

- `GET /api/auth/login/google` - Redirects to Google OAuth 2.0 login.
- `GET /api/auth/google/callback` - Handles Google OAuth 2.0 callback.

### Blog Management

- `POST /api/blogs` - Create a new blog post.
- `GET /api/blogs` - Retrieve all blog posts.
- `GET /api/blogs/:id` - Retrieve a specific blog post by ID.
- `PUT /api/blogs/:id` - Update a specific blog post by ID.
- `DELETE /api/blogs/:id` - Delete a specific blog post by ID.

> **Note**: Blog management endpoints are protected and require a valid JWT token.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
