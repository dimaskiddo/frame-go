package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dimaskiddo/frame-go/utils"
)

func main() {
	// Initialize Channel for OS Signal
	signalOS := make(chan os.Signal, 1)

	// Bootstrap
	utils.Bootstrap()

	// Load Routes
	LoadRoutes()

	// Initialize Server
	server := utils.NewServer(utils.RouterHandler)

	// Starting Server
	server.Start()
	defer server.Stop()

	// Catch OS Signal from Channel
	signal.Notify(signalOS, os.Interrupt, syscall.SIGTERM)

	// Return OS Signal as Exit Code
	<-signalOS

	// Give Information for Server Stop
	fmt.Println(" Stopping Server ")

	// Close Any Database Connections
	switch strings.ToLower(utils.Config.GetString("DB_DRIVER")) {
	case "mysql":
		defer utils.MySQL.Close()
	case "mongo":
		defer utils.MongoSession.Close()
	}
}
