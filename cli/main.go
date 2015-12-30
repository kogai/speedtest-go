package main

import (
	speedtest "github.com/kogai/speedtest-go"
	"gopkg.in/alecthomas/kingpin.v2"
)

func SetTimeout() {
	if *timeoutOpt != 0 {
		timeout = *timeoutOpt
	}
}

var (
	showList   = kingpin.Flag("list", "Show available speedtest.net servers").Short('l').Bool()
	serverIds  = kingpin.Flag("server", "Select server id to speedtest").Short('s').Ints()
	timeoutOpt = kingpin.Flag("timeout", "Define timeout seconds. Default: 10 sec").Short('t').Int()
	timeout    = 10
)

func main() {
	kingpin.Version("1.0.0")
	kingpin.Parse()

	SetTimeout()
	speedtester := speedtest.New()
	speedtester.FetchServers()

	speedtester.ShowUser()

	if *showList {
		speedtester.ShowList()
		return
	}
	speedtester.ShowResult(*serverIds)
}
