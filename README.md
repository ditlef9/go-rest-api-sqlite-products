
# Go API with SQLite database - Products

This is a Go-based REST API to manage products with authentication using JWT and a local SQLite database.
The code is a starter point for new GO developers that wants to setup an API with a local database.

---


## 🔥 Cool Features

* **Efficient Product Management:** Create, retrieve, update, and delete product records easily.
* **SQLite Integration:** Use a lightweight, file-based SQLite database for easy local development.
* **Google Cloud Ready:** Deploy effortlessly to Google Cloud Run to learn howto use GCP.
* **Gin-Gonic Framework:** Built on the blazing-fast Gin web framework, ensuring performance and simplicity.
* **Role-Based Access Control:** Supports authentication for both humans and services, enabling fine-grained access control to API endpoints.
* **Seamless Local Setup:** Get the API running locally in just a few steps on Windows.



---


## 🚀 Running Locally (Windows)

1. **Install Go:** https://golang.org/doc/install
2. **Install TDM-CCC:** Download 64+32-bit MinGW-w64 edition (https://jmeubank.github.io/tdm-gcc/download/) (Requires restart)
3. **Clone repository:** the repository and navigate to the project folder.

4. **Install dependencies:** Run `go mod tidy` to install dependencies.
5. **Run the server locally:** `go run main.go`.
6. **Access API:** The API will be available at `http://localhost:8080`.

---


## 🔑 Setting up a human user and service user

Open Postman, setup the following routes:

* POST http://localhost:8080/api/v1/users/signup<br>
Body > Raw:
```
{
    "email": "human@gmail.com",
    "password": "Vinterferie2024!"
}
```

* POST http://localhost:8080/api/v1/users/signup<br>
Body > Raw:
```
{
    "email": "service@gmail.com",
    "password": "Vinterferie2024!"
}
```

Set the new users as `approved` by opening sqlite.db in DB Browser for SQLite:

```sql
SELECT * FROM users;
UPDATE users SET approved=1;
```

Now login as human or service to get a authorization token:

* POST http://localhost:8080/api/v1/users/login<br>
  Body > Raw:
```
{
    "email": "human@gmail.com",
    "password": "Vinterferie2024!"
}
```

Next you can use the different routes with the authorization token:

* GET http://localhost:8080/api/v1/products<br>
  Headers:
  - Authorization TOKEN


---


## 🛠️ Deploy to Google Cloud Run

1. **Authenticate with GCP:** `gcloud auth login`
2. **Build and deploy:**
    ```bash
    gcloud builds submit --tag gcr.io/[PROJECT_ID]/product-api
    gcloud run deploy --image gcr.io/[PROJECT_ID]/product-api --platform managed
    ```

3. **Access API:** The API will be available via the generated GCP URL.

---

## 💻 Developer Notes

**Initializing the Project:**
```bash
go mod init ekeberg.com/go-api-sql-gcp-products
go get -u github.com/mattn/go-sqlite3
go get -u github.com/gin-gonic/gin
got mod tidy
```

**API Endpoints:**

User Authentication:

* POST /api/v1/users/signup: Register a new user (no authentication required).
* POST /api/v1/users/login: Login a user and receive a JWT token.

Product Management:

* GET /api/v1/products: List all products (authentication required as human or service).
* GET /api/v1/products/:id: Get a product by ID (authentication required as human or service).
* POST /api/v1/products: Add a new product (authentication required as human).
* PUT /api/v1/products/:id: Update an existing product (authentication required as human).
* DELETE /api/v1/products/:id: Delete a product (authentication required as human).

---

## 📖 License

This project is licensed under the
[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0).

```
Copyright 2024 github.com/ditlef9

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```