package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/arpanetus/kun/pkg/util"
)

const usage = `kun v0.0.1 - a command line tool for getting sunrise/sunset times
Usage:
	- init: initializes kun config file
	- issundown: get "true" or "false" value whether sun is down right now or not
	- risesetpair: get sunrise and sunset time pair like "sunrise sunset"`

var (
	debugLogger   = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	defaultLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLogger     = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

var (
	// init flag
	initialize = flag.Bool("init", false, "initialize kun config file")
	// get issundown flag
	issundown = flag.Bool("down", false, "get whether sun is down")
	// get risesetpair flag
	risesetpair = flag.Bool("rs", false, "get sunrise and sunset time pair in HH:MM:SS format")
	// verbose
	verbose = flag.Bool("v", false, "verbose output")
)

func init() {
	flag.Parse()
	if !*verbose {
		debugLogger.SetOutput(ioutil.Discard)
	}
}

func main() {
	if (*issundown && *risesetpair) || (*issundown && *initialize) || (*risesetpair && *initialize) {
		errLogger.Fatal(flag.ErrHelp)

		return
	}

	debugLogger.Printf("initialize: %t, issundown: %t, risesetpair: %t", *initialize, *issundown, *risesetpair)

	if !*issundown && !*risesetpair && !*initialize {
		errLogger.Fatal(flag.ErrHelp)

		return
	}

	if *initialize {
		defaultLogger.Println("initilizing kun config file")

		_, err := util.Config(defaultLogger)
		if err != nil {
			errLogger.Fatalf("cannot initialize kun config file: %s", err)
		}

		return
	}

	c, err := util.Config(debugLogger)
	if err != nil {
		errLogger.Fatalf("cannot get kun config file: %s", err)
	}

	rs := util.NewRiseSet(debugLogger, c)

	if *issundown {
		debugLogger.Println("checking if sun is down")

		down, err := rs.IsSunDown()
		if err != nil {
			errLogger.Println(err)
		}

		fmt.Println(down)
		return
	}

	if *risesetpair {
		debugLogger.Println("getting sunrise and sunset time pair")
		sunrise, sunset, err := rs.SunriseSunset()
		if err != nil {
			errLogger.Fatal(err)
		}

		fmt.Println(sunrise, sunset)

		return
	}
}
