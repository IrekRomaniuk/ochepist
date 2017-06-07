package main

import (
	"github.com/IrekRomaniuk/ochepist/utils"
	"fmt"
	"os"
	"log"
	"flag"
)

var (
	//URL to IP list
	URL       = flag.String("url", "https://minemeld/feeds/office365_IPv4s", "URL to pull IP addresses from")
	//GROUP is object group where IP list is included
	GROUP     = flag.String("g", "g-ochepist-temp", "Object group to include retrieved IP addresses")
	//PATH to dbedit
	PATH     = flag.String("p", "./results/", "path to dbedit file")
	version   = flag.Bool("v", false, "Prints current version")
)
var (
	// Version : Program version
	Version   = "No Version Provided" 
	// BuildTime : Program build time
	BuildTime = ""
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All rights resgit brerved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}	
}


func main() {

data, err:= utils.GetPage(*URL)
if err != nil {
	log.Fatalln(err)
}
	
output := utils.Ip2dbedit(&data, *GROUP)

f, err := os.Create(*PATH + *GROUP + "-dbedit.txt")
check(err)
defer f.Close()
b, err := f.Write(output.Bytes())
check(err)
fmt.Printf("wrote %d bytes\n", b)
}



func check(e error) {
    if e != nil {
        panic(e)
    }
}