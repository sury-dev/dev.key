# 🔐 Dev.key

**A developer-first password manager**  
Securely store, manage, and retrieve development credentials like passwords, API keys, and environment variables — built by developers, for developers.

---

## 🚀 Tech Stack

| Layer       | Technology              |
|-------------|--------------------------|
| Backend     | Go (GoFiber framework)   |
| Database    | MongoDB Atlas            |
| Frontend    | Angular  |
| Auth        | JWT-based    |
| Infra       | (Planned: Docker)        |

---

## 📂 Project Structure (Backend)

```
/cmd/server/        - Main application entry point
/config/            - Environment & configuration loaders
/internal/
    /user/          - User logic (registration, login, etc.)
    /vault/         - Vault logic (key-value storage)
    /middleware/    - Auth, logging, security
/pkg/               - Reusable helpers (JWT, crypto, response)
/database/          - MongoDB connection logic
/routes/            - API route declarations
.env                - Environment config (not committed)
Makefile            - Dev utility commands
README.md           - This file
```

---

## 🧪 API Endpoints (WIP)

| Method | Route      | Description              |
|--------|------------|--------------------------|
| GET    | `/health`  | Check if server is alive |

> ✅ More routes for vault & auth coming soon.

---

## 🛠️ Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/sury-dev/dev.key.git
cd dev.key
```

### 2. Setup Environment

Create a `.env` file in the root directory and configure your secrets:

```dotenv
PORT=8080
MONGO_URI=mongodb+srv://<user>:<password>@cluster.mongodb.net/devkey
JWT_SECRET=yoursupersecret
APP_ENV=preview
```

> ⚠️ Do **not** commit `.env` files to Git. Keep them secret and environment-specific.

### 3. Run the Server

```bash
make start
```

> This will start your GoFiber server. By default, it runs at: `http://localhost:8080`

### 4. Test Health Endpoint

```bash
curl http://localhost:8080/health
```

---

## 🔐 Security

We take developer security seriously. Planned features:

- 🔒 JWT-based authentication
- 🔐 AES-256-GCM encryption for secrets
- 🧠 RBAC for workspace/team-based key control
- 📈 Audit logs (access & mutation tracking)

---

## 🧪 Testing (Coming Soon)

Tests will be added for:

- 🧪 Unit tests (handlers, services)
- 🔁 Integration tests (Mongo, routes)
- ✅ CI integration (GitHub Actions)

---

## 📦 Deployment

Coming soon. Will support:

- Dockerfile & Docker Compose
- MongoDB Atlas config
- Production-ready `.env.example`

---

## 📄 License

[MIT](./LICENSE) – Free to use and modify with attribution.

---

## 🤝 Contributions

Pull requests are welcome. For major changes, please open an issue first  
to discuss what you would like to change or add.

---

## 📬 Contact

Created by Suryansh Pandey ( sury-dev ) — feedback and suggestions welcome!
