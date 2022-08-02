package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gogr/graph"
	"gogr/graph/generated"
	entityModel "gogr/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8485"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	
	dsn := "root:S0o5YX7Nkc2FrZ6Gphc2RzZA@tcp(localhost:3306)/db_getgo?charset=utf8mb4&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	dbErr := db.AutoMigrate(&entityModel.NewTodo{}, &entityModel.Todo{}, &entityModel.User{})
	if dbErr != nil {
		return
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
