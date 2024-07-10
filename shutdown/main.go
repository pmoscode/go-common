// Package shutdown can be used to execute stuff, when the app is going to exit (exit code x or SigTerm)
package shutdown

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// One global instance, which holds all functions to execute on exit
var observerSingleton *observerSingle

// Ensures, that observerSingleton is a real singleton
var once sync.Once

// ObserverFunc Function signature for the function to execute
type ObserverFunc func() error

// Observer Interface of the singleton. Adds a function.
type Observer interface {
	AddCommand(fn ObserverFunc)
}

// observerSingle Struct, which holds the functions to execute.
type observerSingle struct {
	functions []ObserverFunc
}

// AddCommand Adds a function, which will be executed on exit.
// Must match ObserverFunc
func (o *observerSingle) AddCommand(fn ObserverFunc) {
	o.functions = append(o.functions, fn)
}

// hookOnSigTerm Internal function, which hooks on the SigTerm (15).
// Will be automatically executed on first call to GetObserver.
func (o *observerSingle) hookOnSigTerm() {
	channel := make(chan os.Signal)
	//lint:ignore SA1017 Kill code 15 should always lead to the final execution of provided functions
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		finalize("Graceful shutdown hooks")
		close(channel)
		os.Exit(1)
	}()
}

// executeCommands main function, which executes the provided functions.
// Prints the error, if the function return any.
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

// GetObserver Gets the singleton instance.
// Creates one on first call.
func GetObserver() Observer {
	once.Do(func() {
		observerSingleton = &observerSingle{
			functions: make([]ObserverFunc, 0),
		}
		observerSingleton.hookOnSigTerm()
	})

	return observerSingleton
}

// Exit Use this function to do a clean exit. Only here the functions will be executed.
// Not on os.Exit
func Exit(exitCode int) {
	finalize("Shutdown hooks")
	os.Exit(exitCode)
}

// ExitOnPanic This function can be called, when panic is caught. Call then this to exit gracefully.
func ExitOnPanic() {
	if r := recover(); r != nil {
		finalize("Panic shutdown hooks")
		fmt.Println("Printing panic cause: \n", r)
	}
}

// finalize Prints the final log. Counts the failed and succeeded function executions
func finalize(title string) {
	if observerSingleton != nil {
		fmt.Println("########## ", title)
		success, failed := observerSingleton.executeCommands()
		fmt.Println("######## Executed ", success+failed, " shutdown commands")
		fmt.Println("######## Succeeded: ", success)
		fmt.Println("######## Failed: ", failed)
		fmt.Println("##########")
	}
}
