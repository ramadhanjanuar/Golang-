package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// pq is postgres driver for database/sql
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("test")
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=%s port=%d",
		"core", "password", "core", "127.0.0.1", "disable", 5433))
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB(db)

	router := http.NewServeMux()

	router.HandleFunc("/probes/readiness", func(res http.ResponseWriter, req *http.Request) {
		log.Println("masuk")
		if err := db.PingContext(req.Context()); err != nil {
			res.WriteHeader(503)
		}
	})

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", "8081"),
	}

	go func() {
		panic(srv.ListenAndServe())
	}()

	//Create Channel for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	//Recieve shutdown signals.
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}

func closeDB(db io.Closer) {
	if err := db.Close(); err != nil {
		log.Println(errors.New("err closing db connection"))
	} else {
		log.Println("db connection gracefully closed")
	}
}
