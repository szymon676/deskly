package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/szymon676/desk-managment/types"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{
		db: db,
	}
}

func NewPostgresDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS desks (
			id SERIAL PRIMARY KEY,
			Location     TEXT UNIQUE,  
			Availability BOOLEAN 
			);
    `)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}

func (ps *PostgresStore) GetAvailableDesks() ([]types.Desk, error) {
	rows, err := ps.db.Query("select * from desks where availability = True;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var desks []types.Desk
	for rows.Next() {
		desk, err := scanDesk(rows)

		if err != nil {
			return nil, err
		}

		desks = append(desks, *desk)
	}

	return desks, nil
}

func (ps *PostgresStore) CreateDesk(desk types.BindDesk) error {
	if err := types.NewDeskFromRequest(desk); err != nil {
		return err
	}

	query := "insert into desks (location, availability) values($1, $2);"
	_, err := ps.db.Exec(query, desk.Location, "true")
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostgresStore) GetAllDesks() ([]types.Desk, error) {
	rows, err := ps.db.Query("select * from desks;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var desks []types.Desk
	for rows.Next() {
		desk, err := scanDesk(rows)

		if err != nil {
			return nil, err
		}

		desks = append(desks, *desk)
	}

	return desks, nil
}

func scanDesk(rows *sql.Rows) (*types.Desk, error) {
	desk := new(types.Desk)
	err := rows.Scan(
		&desk.ID,
		&desk.Location,
		&desk.Availability,
	)

	return desk, err
}
