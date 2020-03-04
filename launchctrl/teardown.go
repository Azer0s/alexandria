package launchctrl

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RunAndWaitForExit() {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

}

func Teardown() {
	fmt.Println()
	fmt.Println("Bye bye")
	// TODO: Do teardown if any
}
