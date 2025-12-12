# 🍽️ Restaurant Management

Backend for managing menus, foods, tables, orders, order items, invoices, and users — actively being built.

Status: 🚧 In Progress • 🛠️ Building features step by step

## 🧰 Tech Stack
- Go: `1.25.3` (from `go.mod`)
- Dependencies: standard library for now (router/DB to be integrated)

## 🗂️ Project Structure
- `cmd/restaurantManagement/` — entrypoint (`main.go`)
- `controllers/` — request handlers (food, menu, order, table, invoice, user, notes)
- `routes/` — route registration per domain
- `models/` — entity models (Food, Menu, Order, OrderItem, Table, Invoice, User)
- `middlewares/` — cross-cutting concerns (e.g., auth)
- `helpers/` — utilities (e.g., token helpers)
- `database/` — DB connection bootstrap
- `architechture/` — domain relationships and notes

## 🚀 Getting Started
1) Install Go (match the version in `go.mod` or newer)
2) Clone the repo
	 - `git clone https://github.com/suryanshvermaa/restaurantManagement.git`
	 - `cd restaurantManagement`
3) Run (current scaffolding serves as a base)
	 - `go run ./cmd/restaurantManagement`
4) Build a binary
	 - `go build -o bin/restaurant-management ./cmd/restaurantManagement`

## 🔧 Configuration
Set these environment variables as features land:
- `PORT` — HTTP port (e.g., `8080`)
- `DB_URI` — database connection string (Postgres/MySQL/MongoDB)
- `JWT_SECRET` — secret for signing tokens (if using JWT)

You can export them in your shell or use a `.env` loader.

## 🌐 API Overview (Planned & In-Progress)
- 👤 Users: signup, login, profile, list
- 🍽️ Tables: CRUD, assign seats
- 🧾 Menus: CRUD categories
- 🥘 Foods: CRUD items, link to menu
- 🧺 Orders: open/close per table
- 🧩 Order Items: add/update/remove items
- 💳 Invoices: generate and fetch
- 📝 Notes: operational/kitchen notes

Common patterns:
- `GET /api/v1/<resource>` — list
- `POST /api/v1/<resource>` — create
- `GET /api/v1/<resource>/:id` — get by id
- `PUT /api/v1/<resource>/:id` — update
- `DELETE /api/v1/<resource>/:id` — delete

## 🔐 Auth (Planned)
- JWT-based flow via `middlewares/` and `helpers/`
- Protect write operations and sensitive reads

## 🧭 Domain Architecture
See `architechture/README.md` for relationships:
- Menu → Food Items
- Table → Order → Order Items → Invoice

## 🛣️ Roadmap
- Wire HTTP router and register `routes/` from `cmd/restaurantManagement/main.go`
- Implement controllers using `models/`
- Implement `database/databaseConnection.go` (choose driver/ORM)
- Add JWT auth middleware and token utilities
- Add validation, error handling, and logging
- Add tests and CI

## 🧪 Try It Locally
Once the router and DB are wired, a typical start command looks like:
- `PORT=8080 DB_URI="..." JWT_SECRET="..." go run ./cmd/restaurantManagement`

## 📄 License
MIT — see `LICENSE`.

## 🤝 Contributing
Issues and PRs are welcome! Ideas for router integration, DB wiring, and example handlers are especially helpful.