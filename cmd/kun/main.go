package main

import (
	"flag"
	"log"
	"github.com/arpanetus/kun/pkg/util"
	"fmt"
)




const usage = `kun v0.0.1 - a command line tool for getting sunrise/sunset times
Usage:
	- init: initializes kun config file
	- issundown: get "true" or "false" value whether sun is down right now or not
	- risesetpair: get sunrise and sunset time pair like "sunrise sunset"
`

func printHelp() {
	fmt.Println(usage)
}


func main() {
	// init flag
	initialize := flag.Bool("init", false, "initialize kun config file")
	// get issundown flag
	issundown := flag.Bool("issundown", false, "issundown")
	// get risesetpair flag
	risesetpair := flag.Bool("risesetpair", false, "risesetpair")	
	// default or help flag
	help := flag.Bool("help", false, "help")


	flag.Parse()

	if (*issundown && *risesetpair) || (*issundown && *initialize) || (*risesetpair && *initialize) {
		fmt.Println("given options are mutually exclusive")
		printHelp()
		
		return
	}

	log.Printf("initialize: %t, issundown: %t, risesetpair: %t", *initialize, *issundown, *risesetpair)

	if !*issundown && !*risesetpair && !*initialize {
		fmt.Println("no option is selected")
		printHelp()

		return
	}

	if *help {
		printHelp()
		
		return
	}

	if *initialize {
		log.Println("initilizing kun config file")
		util.Config()

		return
	}

	if *issundown {
		log.Println("checking if sun is down")
		fmt.Println(util.IsSunDown())

		return
	}

	if *risesetpair {
		log.Println("getting sunrise and sunset time pair")
		fmt.Println(util.SunriseSunset())

		return
	}



		
}