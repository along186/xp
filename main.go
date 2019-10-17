package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"xp/pkg/Setting"

	"xp/bootstrap"
)

func main()  {

	router := bootstrap.Run()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", Setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    Setting.ReadTimeout,
		WriteTimeout:   Setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), Setting.ReadTimeout)
		defer cancel()
		if err := s.Shutdown(ctx); nil != err {
			log.Fatalf("server shutdown failed, err: %v\n", err)
		}
		log.Println("server gracefully shutdown")
		close(processed)
	}()

	// serve
	err := s.ListenAndServe()
	if http.ErrServerClosed != err {
		log.Fatalf("server not gracefully shutdown, err :%v\n", err)
	}

	// waiting for goroutine above processed
	<-processed
}