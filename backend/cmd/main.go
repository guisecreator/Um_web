package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/guisecreator/um_web/db"
	"github.com/guisecreator/um_web/graphql"
	middleware "github.com/guisecreator/um_web/pkg/middlewares"
	"github.com/guisecreator/um_web/pkg/sessions"
	"github.com/uptrace/bun"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	log.SetFlags(log.Lshortfile)

	database := db.InitBunDb()
	defer func(database *bun.DB) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(database)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(
		cors.New(
			cors.Options{
				AllowedOrigins:      []string{"http://localhost:3000"},
				AllowCredentials:    true,
				AllowPrivateNetwork: true,
				AllowedHeaders: []string{
					"Accept",
					"Content-Type",
					"Content-Length",
					"Accept-Encoding",
					"Authorization",
					"Set-Cookie",
				},
				AllowedMethods: []string{
					"GET",
					"POST",
					"PUT",
					"DELETE",
					"UPDATE",
					"PATCH",
				},
				Debug: true,
				ExposedHeaders: []string{
					"Set-Cookie",
				},
			},
		).Handler)
	router.Use(middleware.AuthMiddleware())

	resolver := &graphql.Resolver{
		Db:       database,
		Sessions: sessions.NewObjectOfSessions(),
	}

	graphConfig := graphql.Config{
		Resolvers: resolver,
	}

	graphConfig.Directives.Authenticated = resolver.Authentication
	srv := CreateServer(graphConfig)

	router.Handle("/", playground.
		Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	httpSrv := http.Server{Addr: ":" + port, Handler: router}

	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Print("Shutting down application...")

		if errShutdown := httpSrv.Shutdown(context.Background()); errShutdown != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", errShutdown)
		}
		close(idleConsClosed)
	}()

	if errListen := httpSrv.ListenAndServe(); errListen != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", errListen)
	}
	<-idleConsClosed

	log.Print("Shutting down application...")
}

func CreateServer(config graphql.Config) *handler.Server {
	srv := handler.NewDefaultServer(
		graphql.NewExecutableSchema(config))
	srv.AddTransport(
		&transport.Websocket{
			Upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					// Check against your desired domains here
					return r.Host == "localhost:3000"
				},
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		})
	return srv
}
