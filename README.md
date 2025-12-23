# Go API Template

A production-ready, modular Go API template featuring **PostgreSQL**, **SQLC**, **Chi Router**, and **Docker** support. Includes a starter implementation for authentication and user management.

## Tech Stack

- **Language**: Go 1.24
- **Router**: [Chi v5](https://github.com/go-chi/chi) - Lightweight, idiomatic, and composable router.
- **Database**: PostgreSQL 17
- **Data Access**: [SQLC](https://sqlc.dev/) - Type-safe Go code generation from SQL.
- **Driver**: [pgx/v5](https://github.com/jackc/pgx) - High-performance PostgreSQL driver.
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations run via Docker.
- **Authentication**: Pre-configured setup (currently using JWT & Bcrypt as a starter).
- **Configuration**: [godotenv](https://github.com/joho/godotenv) - 12-factor app configuration.

## 📂 Project Structure

```
.
├── cmd/
│   └── api/            # Application entry point
├── internal/
│   ├── auth/           # JWT and Password hashing logic
│   ├── config/         # Configuration loading (.env)
│   ├── db/             # Generated SQLC code (Do not edit manually)
│   ├── handlers/       # HTTP Handlers (Controllers)
│   ├── middleware/     # HTTP Middleware (Auth, Logging)
│   └── server/         # Server setup and routing
├── migrations/         # Database migration files (.sql)
├── sql/                # SQL queries for SQLC generation
├── tests/              # Integration tests
├── docker-compose.yml  # Docker services (App, DB, Migrations)
├── Makefile            # Development commands
└── sqlc.yaml           # SQLC configuration
```

## 🛠️ Getting Started

### Prerequisites

- Docker & Docker Compose
- Go 1.24+ (optional, for local development)
- Make (optional)

### Running the Application

1.  **Start everything** (App, DB, Migrations):
    ```bash
    make up
    # OR
    docker compose up --build
    ```

    The API will be available at `http://localhost:8080`.

2.  **Stop the application**:
    ```bash
    make down
    ```

## Testing

Integration tests are located in the `tests/` directory. They spin up a test server and connect to the running Docker database.

```bash
make test
```

*Note: Ensure the Docker container is running (`make up`) before running tests.*

## 💻 Development Workflow

### Database Migrations

Migrations are handled by `golang-migrate`. To add a new migration, create a new pair of `.up.sql` and `.down.sql` files in the `migrations/` directory.

Example: `migrations/000002_add_posts.up.sql`

The migrations are automatically applied when the `migrate` container starts (defined in `docker-compose.yml`).

### Working with SQLC

1.  Write your SQL query in `sql/queries.sql` (or any `.sql` file in `sql/`).
2.  Annotate it with the function name and return type (e.g., `-- name: GetUser :one`).
3.  Generate the Go code:
    ```bash
    make sqlc
    # OR
    sqlc generate
    ```

### Environment Variables

Copy `.env.example` to `.env` to configure local settings.

```bash
cp .env.example .env
```

| Variable       | Description                                      | Default |
| :------------- | :----------------------------------------------- | :------ |
| `PORT`         | Port to listen on                                | `8080`  |
| `DATABASE_URL` | PostgreSQL connection string                     | `...`   |
| `JWT_SECRET`   | Secret key for signing tokens                    | `...`   |

## Example API Endpoints

The template includes a basic user management system to demonstrate the architecture.

### Auth
- `POST /auth/register` - Register a new user.
- `POST /auth/login` - Login and receive a token.

### Users
- `GET /users?email={email}` - Get user details (Protected, requires `Authorization: Bearer <token>`).
