package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "net/http/pprof"

	"github.com/golang-migrate/migrate"
	_ "github.com/lib/pq"

	"github.com/pkg/errors"
	httpBroker "github.com/romanyx/erply/internal/broker/http"
	"github.com/romanyx/erply/internal/storage/postgres"
	"github.com/romanyx/erply/internal/storage/postgres/schema"
)

func main() {
	var (
		addr      = flag.String("addr", ":8080", "address of http server")
		dbURL     = flag.String("db_url", "", "postgres database URL")
		token     = flag.String("token", "token_very_secret", "valid auth token")
		debugAddr = flag.String("debug", ":1234", "debug server addr")
	)
	flag.Parse()

	// Setup db connection.
	db, err := sql.Open("postgres", *dbURL)
	if err != nil {
		log.Fatalf("failed to connect db: %v\n", err)
	}
	defer db.Close()

	// Migrate schema.
	if err := schema.Migrate(db); err != nil {
		if errors.Cause(err) != migrate.ErrNoChange {
			log.Fatalf("failed to migrate schema: %v", err)
		}
	}

	// Setup handlers.
	srv := setupServer(*addr, *token, db)

	// Make a channel for errors.
	errChan := make(chan error)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errChan <- errors.Wrap(err, "failed to serve grpc")
		}
	}()

	// Debug server.
	debugServer := setupDebugServer(*debugAddr)
	go func() {
		if err := debugServer.ListenAndServe(); err != nil {
			errChan <- errors.Wrap(err, "debug server")
		}
	}()
	defer debugServer.Close()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Fatalf("critical error: %v\n", err)
	case <-osSignals:
		log.Println("stopping by signal")
		if err := srv.Close(); err != nil {
			log.Fatalf("failed to stop: %v", err)
		}
	}
}

func setupServer(addr, token string, db *sql.DB) *http.Server {
	repo := postgres.NewRepository(db)
	return httpBroker.NewServer(addr, token, repo)
}

func setupDebugServer(addr string) *http.Server {
	s := http.Server{
		Addr:    addr,
		Handler: http.DefaultServeMux,
	}

	return &s
}
