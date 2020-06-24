package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yeremiaaryo/go-pkg/database"
	"github.com/yeremiaaryo/platform/cmd/internal"
	"github.com/yeremiaaryo/platform/cmd/web"
)

func main() {
	dbConf := database.DBConfig{
		MasterDSN:     "root:test1234@tcp(localhost:3306)/platform?parseTime=true&loc=Local",
		SlaveDSN:      "root:test1234@tcp(localhost:3306)/platform?parseTime=true&loc=Local",
		MaxConn:       100,
		MaxIdleConn:   10,
		RetryInterval: 5,
	}
	db := database.New(dbConf, database.DriverMySQL)
	usecase := internal.GetUsecase(db)

	server := web.New(&web.Options{ListenAddress: ":3000", Usecase: usecase})
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
