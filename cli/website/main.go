package main

import (
	"flag"
	"log"

	"github.com/ilgooz/service-website/website"
	"github.com/mesg-foundation/core/x/xsignal"
	mesg "github.com/mesg-foundation/go-service"
)

var (
	serverAddr = flag.String("addr", ":2300", "Server's listening address")
)

func main() {
	flag.Parse()

	service, err := mesg.New()
	if err != nil {
		log.Fatal(err)
	}

	w, err := website.New(service, *serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	// start the website service.
	go func() {
		log.Println("website service has been started")
		log.Printf("http server listening at %s\n", w.ListeningAddr)

		if err := w.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for interrupt and gracefully shutdown the website service.
	<-xsignal.WaitForInterrupt()

	log.Println("shutting down...")

	if err := w.Close(); err != nil {
		log.Fatal(err)
	}

	log.Println("shutdown")
}
