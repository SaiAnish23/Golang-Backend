package sqlite

import (
	"database/sql"

	"github.com/SaiAnish23/Golang-Backend/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (id INTEGER PRIMARY KEY AUTOINCREMENT,
	 name TEXT,
	 email TEXT,
	  age INTEGER)`)

	if err != nil {
		return nil, err
	}

	return &sqlite{
		Db: db,
	}, nil

}
