package main

import (
	"github.com/IrekRomaniuk/ochepist/utils"
	"fmt"
	"bufio"
	"strings"
	"os"
)



func main() {

data, err:= utils.GetPage("https://minemeld/feeds/office365_IPv4s")
if err != nil {
	os.Exit(1)
}

scanner := bufio.NewScanner(strings.NewReader(data))
      for scanner.Scan() {
				//fmt.Println(scanner.Text())
				line:=scanner.Text()
			  	switch  {
			  	case strings.ContainsAny(line,"-"):
				    s:= strings.Split(line,"-")
					if utils.ValidIP4(s[0]) && utils.ValidIP4(s[1]) {
						fmt.Println("r" + line, s[0], s[1])
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