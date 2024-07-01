package apiserver

import (
	"database/sql"
	"github.com/adametsderschopfer/http-rest-api_go-edu/internal/store/sqlstore"
	"net/http"
)

// Bootstrap ...
func Bootstrap(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	server := newServer(store)

	return http.ListenAndServe(config.BindAddress, server)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
