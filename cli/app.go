package cli

import (
	"log"
	"net/url"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var app = kingpin.New("fantasygoldportal", "FGC DApp Server")

var fantasygoldRPC = app.Flag("fantasygold-rpc", "URL of fantasygold RPC service").Envar("FGC_RPC").Default("").String()

func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
}

func getFantasyGoldRPCURL() *url.URL {
	if *fantasygoldRPC == "" {
		log.Fatalln("Please set FGC_RPC to fantasygoldd's RPC URL")
	}

	url, err := url.Parse(*fantasygoldRPC)
	if err != nil {
		log.Fatalln("FGC_RPC URL:", *fantasygoldRPC)
	}

	if url.User == nil {
		log.Fatalln("FGC_RPC URL (must specify user & password):", *fantasygoldRPC)
	}

	return url
}
