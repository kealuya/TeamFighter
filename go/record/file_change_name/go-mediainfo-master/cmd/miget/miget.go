package main

import (
	"file_change_name/go-mediainfo-master"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatalf("usage: %s filename parameter", os.Args[0])
	}

	mi := mediainfo.New()
	if err := mi.Open(args[0]); err != nil {
		log.Fatal(err)
	}
	fmt.Println((mi.Get(mediainfo.StreamGeneral, 0, args[1])))
	mi.Close()
}
