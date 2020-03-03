package main

import (
	"fmt"
	"github.com/Azer0s/alexandria/launchctrl"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

const VERSION = "0.1.0"
const ASCII_ART = `
        _                          _      _         _____  _   _  _____ 
       | |                        | |    (_)       |  __ \| \ | |/ ____|
   __ _| | _____  ____ _ _ __   __| |_ __ _  __ _  | |  | |  \| | (___  
  / _` + "\u0060" + ` | |/ _ \ \/ / _` + "\u0060" + ` | '_ \ / _` + "\u0060" + ` | '__| |/ _` + "\u0060" + ` | | |  | | . ` + "\u0060" + ` |\___ \
 | (_| | |  __/>  < (_| | | | | (_| | |  | | (_| | | |__| | |\  |____) |
  \__,_|_|\___/_/\_\__,_|_| |_|\__,_|_|  |_|\__,_| |_____/|_| \_|_____/`

func main() {
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.JSONFormatter{})

	fmt.Println(ASCII_ART)
	fmt.Printf("  version %s\n", VERSION)

	launchctrl.Startup()

	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	fmt.Println()
	fmt.Println("Bye bye")
	// TODO: Do teardown if any
}