package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

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

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), Setting.ReadTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}