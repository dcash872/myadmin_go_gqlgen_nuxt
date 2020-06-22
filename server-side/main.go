package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dcash872/myadmin_go_gqlgen_nuxt/server-side/graph"
	"github.com/dcash872/myadmin_go_gqlgen_nuxt/server-side/graph/generated"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	_, error := gorm.Open(
		"mymysql",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			"localhost", 13306, "root", "myadmin", "password",
		),
	)
	if error != nil {
		log.Fatelln(error)
	}

	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	echo.GET("/", welcome())

	// GraphQL設定
	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	echo.POST("/query", func(cxt echo.Context) error {
		graphqlHandler.ServeHTTP(cxt.Response(), cxt.Request())
		return nil
	})

	echo.GET("/playground", func(cxt echo.Context) error {
		playgroundHandler.ServeHTTP(cxt.Response(), cxt.Request())
		return nil
	})

	// Echo起動
	error := echo.Start(":8080")
	if error != nil {
		log.Fatalln(error)
	}
}

func welcome() echo.HandlerFunc {
	return func(cxt echo.Context) error {
		return cxt.String(http.StatusOK, "Welcome!")
	}
}
