package main

import (
	"flag"
	"log"

	"github.com/ilgooz/service-http-server/httpserver"
	"github.com/mesg-foundation/core/client/service"
	"github.com/mesg-foundation/core/x/xsignal"
)

var (
	serverAddr = flag.String("serverAddr", ":2300", "address of server")
)

func main() {
	flag.Parse()

	mesgService, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	s, err := httpserver.New(mesgService, *serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	// start the service.
	go func() {
		log.Println("http server service has been started")
		log.Printf("http server listening at %s\n", s.ListeningAddr())

		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for interrupt and gracefully shutdown the service.
	<-xsignal.WaitForInterrupt()

	log.Println("shutting down...")

	if err := s.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("shutdown")
}
