
# Go API with SQLite database - Products

This is a Go-based REST API to manage products with authentication using JWT and a local SQLite database.
The code is a starter point for new GO developers that wants to setup an API with a local database.

## 🔥 Cool Features

* Efficient Product Management: Create, retrieve, update, and delete product records easily.
* SQLite Integration: Use a lightweight, file-based SQLite database for easy local development.
* Google Cloud Ready: Deploy effortlessly to Google Cloud Run to learn howto use GCP.
* Gin-Gonic Framework: Built on the blazing-fast Gin web framework, ensuring performance and simplicity.
* Seamless Local Setup: Get the API running locally in just a few steps on Windows.


## 🚀 Running Locally (Windows)

1. **Install Go:** https://golang.org/doc/install
2. **Install TDM-CCC:** Download 64+32-bit MinGW-w64 edition (https://jmeubank.github.io/tdm-gcc/download/) (Requires restart)
3. **Clone repository:** the repository and navigate to the project folder.

4. **Install dependencies:** Run `go mod tidy` to install dependencies.
5. **Run the server locally:** `go run main.go`.
6. **Access API:** The API will be available at `http://localhost:8080`.

## 🛠️ Deploy to Google Cloud Run

1. **Authenticate with GCP:** `gcloud auth login`
2. **Build and deploy:**
    ```bash
    gcloud builds submit --tag gcr.io/[PROJECT_ID]/product-api
    gcloud run deploy --image gcr.io/[PROJECT_ID]/product-api --platform managed
    ```

3. **Access API:** The API will be available via the generated GCP URL.


## 📜 Create the SQLite Database Schema

Create table and insert sample data:
```sql
CREATE TABLE products (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   name VARCHAR(200),
   description TEXT,
   ean VARCHAR(13),
   price_out REAL
);

INSERT INTO products (name, description, ean, price_out) VALUES ('Epler', 'Friske, saftige røde epler.', '1234567890123', 29.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Bananer', 'Gule, modne bananer.', '2345678901234', 19.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Melk', 'Helmelk 1 liter, pasteurisert.', '3456789012345', 14.50);
INSERT INTO products (name, description, ean, price_out) VALUES ('Brød', 'Grovbrød med solsikkefrø.', '4567890123456', 29.00);
INSERT INTO products (name, description, ean, price_out) VALUES ('Egg', 'Økologiske egg, 12 stk.', '5678901234567', 39.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Ost', 'Norvegia, halvfast hvitost 500g.', '6789012345678', 79.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Smør', 'Tine usaltet smør 250g.', '7890123456789', 34.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Poteter', 'Norske mandelpoteter, 1 kg.', '8901234567890', 25.00);
INSERT INTO products (name, description, ean, price_out) VALUES ('Laks', 'Fersk laks, fileter 400g.', '9012345678901', 89.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Kyllingfilet', 'Kyllingbrystfilet, 500g.', '1123456789012', 64.50);
INSERT INTO products (name, description, ean, price_out) VALUES ('Spaghetti', 'Fullkorn spaghetti, 500g.', '2234567890123', 19.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Tomatsaus', 'Hjemmelaget tomatsaus på glass.', '3345678901234', 24.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Appelsiner', 'Saftige appelsiner fra Spania.', '4456789012345', 32.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Yoghurt', 'Naturell yoghurt, 1 liter.', '5567890123456', 19.90);
INSERT INTO products (name, description, ean, price_out) VALUES ('Havregryn', 'Grovkornet havregryn, 1 kg.', '6678901234567', 14.90);
```

## 💻 Developer Notes

**Initializing the Project:**
```bash
go mod init ekeberg.com/go-api-sql-gcp-products
go get -u github.com/mattn/go-sqlite3
go get -u github.com/gin-gonic/gin
```

**API Endpoints:**

* GET /api/v1/product: List all products
* GET /api/v1/product/:id: Get a product by ID
* POST /api/v1/product: Add a new product 
* PUT /api/v1/product/:id: Update an existing product
* DELETE /api/v1/product/:id: Delete a product

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