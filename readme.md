# Project Installation

1. Clone the project
   `git clone https://github.com/tegarpratama/bythen`

2. Running docker compose
   `docker-compose up --build -d`

3. Migrate & seed data
   `docker exec -it bythenai-app-1 sh`
   `go run migrate/migrate.go`

# API Endpoints

User Registration & Authentication:

- POST `/register` - Register a new user.
- POST `/login` - Login and receive a token for authentication.

Blog Posts:

- POST `/posts` - Create a new blog post.
- GET `/posts/{id}` - Get blog post details by ID.
- GET `/posts` - List all blog posts.
- PUT `/posts/{id}` - Update a blog post.
- DELETE `/posts/{id}` - Delete a blog post.

Comments:

- POST `/posts/{id}/comments` - Add a comment to a blog post.
- GET `/posts/{id}/comments` - List all comments for a blog post.
