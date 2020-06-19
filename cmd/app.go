package main

import (
	"log"
	"os"
	"os/signal"
	"platform/cmd/web"
	"syscall"
)

func main() {
	server := web.New(&web.Options{ListenAddress: ":3000"})
	go server.Run()

	select {
	case _ = <-terminateSignal():
		log.Println("Exiting gracefully...")
	case err := <-server.ListenError():
		log.Println("Error starting web server, exiting gracefully:", err)
	}
}

func terminateSignal() chan os.Signal {
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	return term
}
