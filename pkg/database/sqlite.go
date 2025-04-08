package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func NewSQLiteConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "../data.db")
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к SQLite: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("(Auth) Не удалось проверить связь с SQLite: %w", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("(Auth) Ошибка создания драйвера миграции: %v\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3", driver,
	)
	if err != nil {
		log.Fatalf("ошибка создания миграции: %v\n", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("нет новых миграций для выполнения")
		} else {
			log.Fatalf("ошибка выполнения миграции: %v\n", err)
		}
	} else {
		log.Println("миграция выполнена успешно")
	}

	return db, nil
}
