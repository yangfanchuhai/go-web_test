package main

import (
	"context"
	_ "github.com/yangfanchuhai/go-web_test/internal/store"
	"github.com/yangfanchuhai/go-web_test/server"
	"github.com/yangfanchuhai/go-web_test/store/factory"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem")
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8080", s)

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("web server start failed:", err)
	}

	log.Println("web server start ok")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <- errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}
}