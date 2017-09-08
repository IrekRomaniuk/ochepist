package main

import (
	"github.com/IrekRomaniuk/ochepist/utils"
	"fmt"
	"os"
	"log"
	"flag"
	"crypto/md5"
)

var (
	//URL to IP list
	URL       = flag.String("url", "https://minemeld/feeds/office365_IPv4s", "URL to pull IP addresses from")
	//GROUP is object group where IP list is included
	GROUP     = flag.String("g", "g-ochepist", "Object group to include retrieved IP addresses")
	//PATH to results
	PATH    = flag.String("p", "./results/", "path to created dbedit file")
	//TEMPLATES to dbedit templates
	TEMPLATES    = flag.String("t", "templates/*", "path to dbedit tempalte files")
	//COMMENT added to created objects
	COMMENT     = flag.String("c", "Created by ochepist with dbedit", "comment added to objects")
	version   = flag.Bool("v", false, "Prints current version")
	// Version : Program version
	Version   = "No Version Provided" 
	// BuildTime : Program build time
	BuildTime = ""
)

const (
	// HASH is the name of the file where state is kept
	HASH = "hash" 
	// OXML with O365IPAddresses.xml
	OXML = "https://support.content.office.net/en-us/static/O365IPAddresses.xml"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2017 @IrekRomaniuk. All rights reversed.\n")
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

	var (
		hash int // exit value
		err error
		data string	// data downloaded from web 
	) 
	// Crate hash file if not exists to track changes on web
	if _, err = os.Stat(*PATH + HASH); os.IsNotExist(err) {
		h, err := os.Create(*PATH + HASH)
		check(err)
		defer h.Close()
	}	
	// download ip address data either from web directly or extract from xml
	if *URL == "O365IPAddresses.xml" {
		data, err = utils.ReadXML(OXML)		
	} else {
		data, err = utils.GetPage(*URL)
	}
	check(err)

	oldhash, _ := utils.GetHash(*PATH + HASH)
	newhash := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	//fmt.Printf("old: %s new: %s\n", oldhash, newhash)
	if (oldhash != newhash) {
		if err = utils.SetHash(newhash, *PATH + HASH); err != nil {
				hash = 2	// //HASH has changed but not written to file
			} else {
				hash = 1 //HASH has chaned
			}
			// parses data to dbedit format	
			output := utils.IP2dbedit(data, *GROUP, *COMMENT, *TEMPLATES)
			f, err := os.Create(*PATH + *GROUP + "-dbedit.txt")
			check(err)
			defer f.Close()
			// writes data in dbedit format to file
			_, err = f.Write(output.Bytes())
			check(err)
		
	} else  {
		hash = 0 // HASH NOT CHANGED
	}
	os.Exit(hash)
}

func check(e error) {
    if e != nil {
        log.Fatalln(e)
    }
}

func hand(path string) (*os.File, error) {
	h, err := os.Create(path)
	return h, err
}