package shutdown

import (
	"fmt"
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

func (o *observerSingle) hookOnSigTerm() {
	channel := make(chan os.Signal)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		finalize("Graceful shutdown hooks")
		os.Exit(1)
	}()
}

func (o *observerSingle) executeCommands() (int, int) {
	success := 0
	failed := 0
	for _, command := range o.functions {
		err := command()
		if err != nil {
			fmt.Println(err)
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
		observerSingleton.hookOnSigTerm()
	})

	return observerSingleton
}

func Exit(exitCode int) {
	finalize("Shutdown hooks")
	os.Exit(exitCode)
}

func ExitOnPanic() {
	if r := recover(); r != nil {
		finalize("Panic shutdown hooks")
		fmt.Println("Printing panic cause: \n", r)
	}
}

func finalize(title string) {
	fmt.Println("########## ", title)
	success, failed := observerSingleton.executeCommands()
	writeSummaryToConsole(success, failed)
	fmt.Println("##########")
}

func writeSummaryToConsole(success, failed int) {
	fmt.Println("######## Executed ", success+failed, " shutdown commands")
	fmt.Println("######## Succeeded: ", success)
	fmt.Println("######## Failed: ", failed)
}
