package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Gateway311/engine/request"
	"Gateway311/engine/router"
	"Gateway311/engine/services"
	"Gateway311/engine/telemetry"

	"github.com/ant0ine/go-json-rest/rest"
	log "github.com/jeffizhungry/logrus"
)

var (
	configFile string

	// Debug switches on some debugging statements.
	Debug = false
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/v1/services.json", request.Services),
		rest.Post("/v1/requests.json", request.Create),
		rest.Get("/v1/requests.json", request.Search),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func init() {
	log.Setup(false, log.DebugLevel)

	flag.BoolVar(&Debug, "debug", false, "Activates debug logging.")
	flag.StringVar(&configFile, "config", "config.json", "Config file. This is a full or relative path.")
	flag.Parse()

	fmt.Printf("Debug: %v  Config: %v\n", Debug, configFile)

	if err := router.Init(configFile); err != nil {
		log.Fatal("Unable to start - data initilization failed.\n")
	}

	telemetry.Init(router.GetMonitorAddress())

	go signalHandler(make(chan os.Signal, 1))
	fmt.Println("Press Ctrl-C to shutdown...")

	time.Sleep(time.Second * 2)
	services.Refresh()
}

func signalHandler(c chan os.Signal) {
	signal.Notify(c, os.Interrupt)
	for s := <-c; ; s = <-c {
		switch s {
		case os.Interrupt:
			fmt.Println("Ctrl-C Received!")
			stop()
			os.Exit(0)
		case os.Kill:
			fmt.Println("SIGKILL Received!")
			stop()
			os.Exit(1)
		}
	}
}

func stop() error {
	return nil
}
