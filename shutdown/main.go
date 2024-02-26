package shutdown

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var observerSingleton *observerSingle
var once sync.Once

type ObserverFunc func() error

type Observer interface {
	AddCommand(fn ObserverFunc)
}

type observerSingle struct {
	functions []ObserverFunc
}

func (o *observerSingle) AddCommand(fn ObserverFunc) {
	o.functions = append(o.functions, fn)
}

func (o *observerSingle) hookOnShutdown() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		log.Println("Shutting down gracefully...")
		success, failed := o.executeCommands()
		log.Println("Executed ", success+failed, " shutdown commands")
		log.Println("Succeeded: ", success)
		log.Println("Failed: ", failed)
		log.Println("...exiting.")
		os.Exit(1)
	}()
}

func (o *observerSingle) executeCommands() (int, int) {
	success := 0
	failed := 0
	for _, command := range o.functions {
		err := command()
		if err != nil {
			log.Println(err)
			failed++
		} else {
			success++
		}
	}

	return success, failed
}

func GetObserver() Observer {
	once.Do(func() {
		observerSingleton = &observerSingle{
			functions: make([]ObserverFunc, 0),
		}
		observerSingleton.hookOnShutdown()
	})

	return observerSingleton
}
