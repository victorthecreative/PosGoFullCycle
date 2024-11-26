package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/victorthecreative/PosGoFullCycle/GraphQL/graph"
	"github.com/victorthecreative/PosGoFullCycle/GraphQL/internal/database"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "/Users/victorcoelho/workspace_go/pos_go/graphql/cmd/server/data.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()
	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		CourseDB:   courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
