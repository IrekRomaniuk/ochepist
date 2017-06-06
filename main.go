package main

import (
	"github.com/IrekRomaniuk/ochepist/utils"
	"fmt"
	"bufio"
	"strings"
	"os"
	"text/template"
	"log"
)

func init() {
	
}


func main() {

data, err:= utils.GetPage("https://minemeld/feeds/office365_IPv4s")
if err != nil {
	os.Exit(1)
}

tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

dbedit := map[string]string{"name": "", "ipaddr": "", "ipaddr_first": "", "ipaddr_last": "", "netmask": ""}

scanner := bufio.NewScanner(strings.NewReader(data))
	
	
      for scanner.Scan() {
				//fmt.Println(scanner.Text())
				line:=strings.Trim(scanner.Text(),"")
			  	switch  {
			  	case strings.ContainsAny(line,"-"):
				    s:= strings.Split(line,"-")
					if utils.ValidIP4(s[0]) && utils.ValidIP4(s[1]) {
						//fmt.Println("r" + line, s[0], s[1])
						dbedit["name"] = "r" + line
						dbedit["ipaddr_first"] = s[0]
						dbedit["ipaddr_last"] = s[1]
						
						err := tpl.ExecuteTemplate(os.Stdout, "ranges.gotxt", dbedit)
						if err != nil {
						log.Fatalln(err)
						}
						//fmt.Println("\n",dbedit)
					}		
			 	case strings.ContainsAny(line,"/"):
				 	s:= strings.Split(line,"/")
					if utils.ValidIP4(s[0]) && utils.ValidIP4(s[1]) {
						fmt.Println("n" + s[0] + "-" + s[1], s[0], s[1])
					}		
				default: 
					if utils.ValidIP4(line) {
						fmt.Println("h" + line)
					}
						
			  }
      }

}