// db/db.go
package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var DB *sql.DB

func InitDB() {
	log.Println("db.go::InitDB()::Init --------------------------------")

	// Connect to database
	var err error
	DB, err = sql.Open("sqlite3", "sqlitedb.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	// Set connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Run migrations
	createMigrationTable()
	runMigrations()
}

func createMigrationTable() {
	log.Println("db.go::createMigrationTable()::Creating migration table")
	var sql string = `
	CREATE TABLE IF NOT EXISTS migrations (
      migration_id INTEGER PRIMARY KEY AUTOINCREMENT,
      migration_module VARCHAR(255),
      migration_name VARCHAR(255) NOT NULL,
      migration_run_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
	`

	_, err := DB.Exec(sql)

	if err != nil {
		panic("db.go::createMigrationTable()::Error could not create migration table:" + err.Error())
	}

}

func runMigrations() {
	log.Println("db.go::runMigrations()::Init")
	files, err := ioutil.ReadDir("db/migrations")
	if err != nil {
		panic("db.go::runMigrations()::Error reading migrations directory: " + err.Error())
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			migrationName := file.Name()
			if !isMigrationApplied(migrationName) {
				err := applyMigration(migrationName)
				if err != nil {
					log.Printf("db.go::runMigrations()::Error applying migration %s: %v", migrationName, err)
					panic(err)
				}
				logMigration(migrationName)
			}
		}
	}
}

func isMigrationApplied(migrationName string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM migrations WHERE migration_name = ?)`
	err := DB.QueryRow(query, migrationName).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("db.go::isMigrationApplied()::Error checking migration %s: %v", migrationName, err)
	}
	return exists
}

func applyMigration(migrationName string) error {
	log.Printf("db.go::applyMigration()::Applying migration %s", migrationName)
	content, err := os.ReadFile(filepath.Join("db/migrations", migrationName))
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(content))
	return err
}

func logMigration(migrationName string) {
	sql := `INSERT INTO migrations (migration_module, migration_name) VALUES (?, ?)`
	_, err := DB.Exec(sql, "default", migrationName)
	if err != nil {
		log.Fatalf("db.go::logMigration()::Error logging migration %s: %v", migrationName, err)
	}
	log.Printf("db.go::logMigration()::Migration %s logged successfully", migrationName)
}
