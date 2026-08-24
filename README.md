# Mini LMS Platform — Backend API

A backend API for an online learning platform. Students can register, browse courses, enroll, watch lessons, take quizzes, track progress, and receive certificates upon completion. Admins manage all content and users.

---

## Tech Stack

| Layer | Technology |
|-------|------------|
| Language | Go 1.26+ |
| Framework | Gin |
| ORM | Ent |
| Database | MySQL 8 |
| Cache | Redis |
| Search | Elasticsearch 8 |
| Queue | RabbitMQ |
| Docs | Swagger (`/docs`) |
| Deploy | Docker Compose |

---

## Architecture

```
Request → Delivery → Handler → Usecase → Repository → MySQL
```

- **Delivery** — route registration, middleware setup
- **Handler** — request binding, validation, response
- **Usecase** — business logic
- **Repository** — database queries via Ent ORM
- **DI** — manual dependency injection at `di/di.go`

---

## Project Structure

```
lms-backend/
├── docker-compose.yml
├── .env.example
├── lms-api/                    # API server
│   ├── cmd/
│   │   ├── main.go
│   │   ├── app.go
│   │   └── seed/               # seed data
│   ├── ent/
│   │   └── schema/             # 12 entity schemas
│   └── internal/
│       ├── common/             # env, middleware, response, pagination
│       ├── delivery/           # route registration
│       │   └── admin/
│       ├── di/                 # dependency injection
│       ├── dto/                # request/response structs
│       ├── handler/            # HTTP handlers
│       │   └── admin/
│       ├── repository/         # DB interfaces + implementations
│       │   └── repository_impl/
│       └── usecase/            # business logic
│           └── usecase_impl/
└── lms-worker/                 # RabbitMQ consumers
    ├── cmd/
    ├── ent/
    └── internal/
        ├── common/
        ├── delivery/           # consumer registration
        ├── di/
        ├── handler/            # message handlers
        ├── repository/
        └── usecase/
```

---

## Getting Started

### Prerequisites

- Go 1.26+
- Docker & Docker Compose

### 1. Clone and configure

```bash
git clone <repo-url>
cd lms-backend

cp .env.example .env
cp lms-api/.env.example lms-api/.env
cp lms-worker/.env.example lms-worker/.env
```

Edit `lms-api/.env`:

```env
IS_PRODUCTION=false
HOST=0.0.0.0
PORT=8080

DATABASE_URL=root:password@tcp(localhost:3310)/lms-backend?parseTime=True

REDIS_ADDR=localhost:6379
REDIS_PASS=password

ELASTIC_ADDRS=http://localhost:9200
ELASTIC_USER=
ELASTIC_PASSWORD=
ELASTIC_CERT_FINGERPRINT=

RABBIT_MQ_URL=amqp://user:password@localhost:5672/

SECRET_ACCESS_TOKEN=your-secret-key
EXPIRES_AT_ACCESS_TOKEN=1h

SECRET_REFRESH_TOKEN=your-refresh-secret-key
EXPIRES_AT_REFRESH_TOKEN=24h
```

### 2. Start services

```bash
docker compose up -d
```

| Service | URL |
|---------|-----|
| MySQL | `localhost:3310` |
| Redis | `localhost:6379` |
| Elasticsearch | `http://localhost:9200` |
| Kibana | `http://localhost:5601` |
| RabbitMQ Management | `http://localhost:15672` |

### 3. Run API server

```bash
cd lms-api
go run ./cmd/...
```

Server runs at `http://localhost:8080`

### 4. Run Worker

```bash
cd lms-worker
go run ./cmd/...
```

### 5. Seed data

```bash
cd lms-api
make seed
# or directly:
mysql -h 127.0.0.1 -P 3310 -u root -p lms-backend < seed.sql
```

Creates: 1 admin, 3 students, 5 categories, 5 courses, 10 sections, 20 lessons, 5 quizzes, 20 questions.

---

## Default Accounts (after seed)

| Role | Email | Password |
|------|-------|----------|
| Admin | admin@cybersoft.edu.vn | Admin@123 |
| Student | an.nguyen@email.com | Student@123 |
| Student | binh.tran@email.com | Student@123 |
| Student | chi.le@email.com | Student@123 |

---

## API Documentation

Swagger UI: `http://localhost:8080/docs`

### Auth

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/register` | Register student |
| POST | `/api/auth/login` | Login |
| POST | `/api/auth/refresh-token` | Refresh token |
| GET | `/api/auth/get-info` | Get current user info |
| PATCH | `/api/auth/change-password` | Change password |

### Public Courses

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/categories` | List categories |
| GET | `/api/courses` | List published courses |
| GET | `/api/courses/search` | Search courses via Elasticsearch |
| GET | `/api/courses/:id` | Course detail |
| GET | `/api/courses/:id/preview-lessons` | Preview lessons |

### Student — Enrollment & Learning

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/enrollments` | Enroll in course |
| GET | `/api/my/enrollments` | My enrolled courses |
| GET | `/api/my/courses/:courseId` | Enrolled course detail |
| DELETE | `/api/my/enrollments/:id` | Cancel enrollment |
| GET | `/api/my/courses/:courseId/lessons` | Lesson list |
| GET | `/api/my/lessons/:lessonId` | View lesson |
| POST | `/api/my/lessons/:lessonId/complete` | Mark lesson complete |
| GET | `/api/my/courses/:courseId/progress` | Course progress |
| GET | `/api/my/quizzes/:quizId` | Get quiz |
| POST | `/api/my/quizzes/:quizId/submit` | Submit quiz |
| GET | `/api/my/quizzes/:quizId/attempts` | Quiz attempt history |
| GET | `/api/my/certificates` | My certificates |
| GET | `/api/my/certificates/:course_id` | Certificate by course |
| GET | `/api/my/notifications` | Notifications |
| PATCH | `/api/my/notifications/:id/read` | Mark as read |
| PATCH | `/api/my/notifications/read-all` | Mark all as read |

### Admin

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/admin/dashboard/overview` | System overview |
| GET | `/api/admin/dashboard/top-courses` | Top courses |
| GET | `/api/admin/users` | User management |
| PATCH | `/api/admin/users/:id/status` | Block/unblock user |
| PATCH | `/api/admin/users/:id/role` | Change role |
| POST | `/api/admin/categories` | Create category |
| GET/PUT/DELETE | `/api/admin/categories/:id` | Category CRUD |
| POST | `/api/admin/courses` | Create course |
| GET/PUT/DELETE | `/api/admin/courses/:id` | Course CRUD |
| PATCH | `/api/admin/courses/:id/status` | Change course status |
| POST | `/api/admin/courses/:id/reindex` | Reindex to Elasticsearch |
| POST/GET/PUT/DELETE | `/api/admin/courses/:courseId/sections` | Section management |
| POST/GET/PUT/DELETE | `/api/admin/sections/:sectionId/lessons` | Lesson management |
| POST/GET/PUT/DELETE | `/api/admin/lessons/:lessonId/quizzes` | Quiz management |
| POST/GET/PUT/DELETE | `/api/admin/quizzes/:quizId/questions` | Question management |
| GET | `/api/admin/enrollments` | Enrollment management |
| GET | `/api/admin/certificates` | Certificate management |

---

## Key Features

### Automatic Certificate Issuance

When a student completes **100% of lessons** and **passes all quizzes** in a course:

1. A certificate is automatically created with a unique UUID code
2. Enrollment status is set to `completed`
3. A notification is sent via RabbitMQ (`notification.certificate_issued`)

### Redis Cache

- Course list: `courses:list:{hash}` — TTL 60s
- Course detail: `courses:detail:{id}` — TTL 60s
- Refresh token: `refresh_token:{userId}:{tokenId}` — TTL 24h
- Cache is invalidated when admin modifies course content

### Elasticsearch Search

Index `courses` — full-text search on `title` and `description`, returns only `published` courses.

### RabbitMQ Queues

| Queue | Trigger |
|-------|---------|
| `notification.course_enrolled` | Student enrolls in course |
| `notification.certificate_issued` | Student completes course |
| `search.course_index` | Admin creates/updates course |

---

## Response Format

Success:

```json
{
  "data": {},
  "message": "success"
}
```

Error:

```json
{
  "message": "error message"
}
```

---

## Database Schema

12 entities: `users`, `categories`, `courses`, `sections`, `lessons`, `enrollments`, `lesson_progresses`, `quizzes`, `questions`, `quiz_attempts`, `certificates`, `notifications`.

Soft delete on: `users`, `categories`, `courses`, `sections`, `lessons`, `quizzes`, `questions`.

---

## Modules

| # | Feature |
|---|---------|
| 1 | Project Setup, Docker Compose |
| 2 | Auth, JWT, Middleware |
| 3 | Admin — Category & Course CRUD |
| 4 | Admin — Section management |
| 5 | Admin — Lesson management |
| 6 | Admin — Quiz management |
| 7 | Admin — Question management |
| 8 | Client — Public courses, categories |
| 9 | Admin — User management |
| 10 | Enrollment |
| 11 | Lesson Progress |
| 12 | Quiz Client |
| 13 | Admin — Enrollment management |
| 14 | Redis Cache & Refresh Token |
| 15 | Elasticsearch |
| 16 | RabbitMQ Workers & Notifications |
| 17 | Certificates |
| 18 | Admin Dashboard & Swagger |
