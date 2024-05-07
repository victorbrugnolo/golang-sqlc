package main

import (
	"context"
	"database/sql"
	"fmt"
	"golang-sqlc/internal/db"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(queries *db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)

	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}

		return err
	}

	return tx.Commit()
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}

	defer func(dbConn *sql.DB) {
		_ = dbConn.Close()
	}(dbConn)

	queries := db.New(dbConn)
}
