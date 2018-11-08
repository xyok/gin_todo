package main

import (
	"context"
	"flag"
	"fmt"
	"gin_todo/models"
	"gin_todo/routers"
	"gin_todo/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func Run() {
	models.Setup()

	router := routers.InitRouter()
	addr := ":" + strconv.Itoa(setting.ServerSetting.HttpPort)

	log.Println("listen ", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	models.CloseDB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func migrate() {
	models.Setup()
	models.Migrate()
	models.CloseDB()
}

func usage() {
	content := `command usage:

    run     	- run server
    migrate     - create tables
    help       	- print this help
`

	fmt.Println(content)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	action := flag.Arg(0)
	switch action {
	case "run":
		Run()
	case "migrate":
		migrate()
	default:
		flag.Usage()
	}

}
