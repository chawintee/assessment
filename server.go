package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/chawintee/assessment/handle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func SetUpDB(url string) (*sql.DB, func()) {
	var err error
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createTb := `
		CREATE TABLE IF NOT EXISTS expenses (
			id SERIAL PRIMARY KEY,
			title TEXT,
			amount FLOAT,
			note TEXT,
			tags TEXT[]
		);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	teardown := func() {
		db.Close()
	}
	return db, teardown
}

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
	fmt.Println("start at database_url:", os.Getenv("DATABASE_URL"))

	db, teardown := SetUpDB(os.Getenv("DATABASE_URL"))
	defer teardown()

	e := echo.New()
	h := handle.NewApplication(db)
	e.Logger.SetLevel(log.INFO)
	e.Use(h.AuthMiddleware)
	e.POST("/expenses", h.CreateExpenses)
	e.GET("/expenses", h.GetExpenses)
	e.GET("/expenses/:id", h.GetExpense)
	e.PUT("/expenses/:id", h.EditExpense)
	// log.Fatal(e.Start(os.Getenv("PORT")))
	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
