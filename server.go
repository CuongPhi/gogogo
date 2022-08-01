package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gogr/graph"
	"gogr/graph/generated"
	"gogr/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8485"

//type User struct {
//	gorm.Model
//	Name string
//	Age  uint
//}
//
//type Todo struct {
//	gorm.Model
//	Text string
//	Done bool
//}

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
	dbErr := db.AutoMigrate(&model.NewTodo{}, &model.Todo{}, &model.User{})
	if dbErr != nil {
		return
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
