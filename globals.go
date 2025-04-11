package main

import (
	"flag"
	"time"
)

var Delay = 1 * time.Second

// Flags
var HideDelayFlag = flag.Bool("quick", false, "Skip delays")
var NoPressEnterFlag = flag.Bool("noenter", false, "Skip press enter command")
var UseMockFlag = flag.Bool("mock", false, "Use mock to fill players name")
