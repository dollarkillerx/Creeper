package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dollarkillerx/creeper/sdk/creeper_sdk"
)

var creeperAddr = flag.String("creeperAddr", "", "Creeper Addr")
var creeperToken = flag.String("creeperToken", "", "Creeper Token")
var creeperIndex = flag.String("creeperIndex", "", "Creeper Index")
var logPath = flag.String("logPath", "", "Log Path")

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		creeperAddr = refString(os.Getenv("CreeperAddr"))
		creeperToken = refString(os.Getenv("CreeperToken"))
		creeperIndex = refString(os.Getenv("CreeperIndex"))
		logPath = refString(os.Getenv("LogPath"))
		if *logPath == "" {
			fmt.Println("Creeper Rlog")
			flag.PrintDefaults()
			return
		}
	}

	sdk := creeper_sdk.New(*creeperAddr, *creeperToken)
	_, err := sdk.Index()
	if err != nil {
		log.Fatalln("creeper configuration error: ", err)
	}

	if *logPath == "" {
		log.Fatalln("log path is null")
	}

	fmt.Println("Log: ", *logPath)

	open, err := os.Open(*logPath)
	if err != nil {
		log.Fatalln("Read LogPath error: ", err)
	}
	defer open.Close()

	open.Seek(0, 2)

	reader := bufio.NewReader(open)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			time.Sleep(time.Second)
			continue
		}

		err = sdk.Log(*creeperIndex, string(bytes))
		if err != nil {
			log.Println(err)
		}
	}
}

func refString(r string) *string {
	return &r
}
