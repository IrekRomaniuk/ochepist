package utils

import (
	"text/template"
	"bytes"
	"log"
	"bufio"
	"strings"
)

//Write data with IP addresses in dbedit format
func Ip2dbedit(data, group, comment, templates string) bytes.Buffer {
	tpl, err := template.ParseGlob(templates)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(data))
	dbedit := map[string]string{"name": "", "ipaddr": "", "ipaddr_first": "", "ipaddr_last": "", "netmask": "", "group": group, "comment": comment} 
	var result bytes.Buffer
	result.WriteString("create network_object_group g-net_group " + group + "\n")
	for scanner.Scan() {
				line:=strings.Trim(scanner.Text(),"")
				var out bytes.Buffer
			  	switch  {
			  	case strings.ContainsAny(line,"-"):
				    s:= strings.Split(line,"-")
					if ValidIP4(s[0]) && ValidIP4(s[1]) {
						dbedit["name"] = "r" + line
						dbedit["ipaddr_first"] = s[0]
						dbedit["ipaddr_last"] = s[1]
						
						if err = tpl.ExecuteTemplate(&out, "ranges.gotxt", dbedit); err != nil {
						log.Fatalln(err)
						}
						result.WriteString(out.String())
					}		
			 	case strings.ContainsAny(line,"/"):
				 	s:= strings.Split(line,"/")
					if ValidIP4(s[0]) && ValidIP4(s[1]) {
						dbedit["name"] = "n" + s[0] + "-" + s[1]
						dbedit["ipaddr"] = s[0]
						dbedit["netmask"] = s[1]
						
						if err = tpl.ExecuteTemplate(&out, "networks.gotxt", dbedit); err != nil {
						log.Fatalln(err)
						}
						result.WriteString(out.String())
					}		
				default: 
					s := strings.Split(line,"")
					if ValidIP4(line) {
						dbedit["name"] = "r" + line
						dbedit["ipaddr"] = s[0]
						
						if err = tpl.ExecuteTemplate(&out, "hosts.gotxt", dbedit); err != nil {
						log.Fatalln(err)
						}
						result.WriteString(out.String())
					}
						
			  }
      }
	  return result
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}