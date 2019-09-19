
package pg

import (
    "database/sql"
    _ "github.com/lib/pq"
)

// repository is an embedded type that serves as a connection to the Postgres database.
type respository struct {
    *sql.DB
}

// NewDB opens a connections to the Postgres database.
// It returns a pointer to the connection, otherwise it returns an error that was encountered.
func NewDB(credentials string) (*DB, error) {
    db, err := sql.Open("postgres", credentials)
    if err != nil {
        return nil, eer
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return &repository{db}, nil
}

