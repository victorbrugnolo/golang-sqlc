package main

import (
	"context"
	"database/sql"
	"golang-sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

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

	//err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "Backend description", Valid: true},
	//})

	//if err != nil {
	//	panic(err)
	//}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
		ID:          "9113f4a0-0225-48a3-9a04-636fc833c126",
	})

	if err != nil {
		panic(err)
	}

	//err = queries.DeleteCategory(ctx, "9ad617fb-258a-461b-bb74-2414a1ddd73e")
	//err = queries.DeleteCategory(ctx, "dd346673-aa52-43a1-ab0a-b7ca11423602")

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
