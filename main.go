package main

import (
	"fmt"
	"github.com/Azer0s/alexandria/launchctrl"
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
	cfg := launchctrl.GetConfig()
	launchctrl.ConfigureLog(cfg)
	launchctrl.ConfigureZones(cfg)

	if cfg.PrintTitle {
		fmt.Println(ASCII_ART)
		fmt.Printf("  version %s\n", VERSION)
	}

	launchctrl.Startup(cfg)
	launchctrl.RunAndWaitForExit()
	launchctrl.Teardown()
}
