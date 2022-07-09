package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/gops/agent"
	"github.com/ssup2/golang-tracing-example/pkg/net"
	"github.com/ssup2/golang-tracing-example/pkg/sync"
)

func main() {
	// Run HTTP server to expose profile endpoint
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Run gops agent
	go func() {
		agent.Listen(agent.Options{})
	}()

	// Run goroutines to load
	go sync.Mutex01()
	go sync.Mutex02()
	go sync.Mutex03()
	go net.Receive()
	go net.Send()

	// Block until receive a terminal signal
	log.Println("Waiting a terminal signal to shutdown gracefully")
	termSignal := make(chan os.Signal, 1)
	signal.Notify(termSignal, syscall.SIGTERM, syscall.SIGINT)
	<-termSignal
}
