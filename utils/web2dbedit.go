package utils

import (
	"text/template"
	"bytes"
	"log"
	"bufio"
	"strings"
)

var tpl *template.Template
var dbedit map[string]string
var result bytes.Buffer

//Write data with IP addresses in dbedit format
func Ip2dbedit(data, group, comment, templates string) bytes.Buffer {
	var err error
	tpl, err = template.ParseGlob(templates)
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(data))
	dbedit = map[string]string{"name": "", "ipaddr": "", "ipaddr_first": "", "ipaddr_last": "", 
		"netmask": "", "group": group, "comment": comment} 
	result.WriteString("create network_object_group " + group + "\n")
	for scanner.Scan() {
				line:=strings.Trim(scanner.Text(),"")				
			  	switch  {
			  	case strings.ContainsAny(line,"-"):
				    s:= strings.Split(line,"-")
					if ValidIP4(s[0]) && ValidIP4(s[1]) {
						write2dbedit("ranges.gotxt", "r", "-", s[0],s[1])
					}					
			 	case strings.ContainsAny(line,"/"):
				 	s:= strings.Split(line,"/")
					if ValidIP4(s[0]) && ValidIP4(s[1]) {
						write2dbedit("networks.gotxt", "n", "-", s[0],s[1])
					}											
				default: 
					s := strings.Split(line,"\n")
					if ValidIP4(s[0]) { 
						write2dbedit("hosts.gotxt", "h", "", s[0])
					}											
			  }
      }
	  return result
}

func write2dbedit(template, prefix, separator string, ip ...string) { 
	var out bytes.Buffer
	var ip2nd string												
	
	if len(ip) > 1 { 
		ip2nd = ip[1] 
	} else {
			ip2nd = ""
	}							
	dbedit["name"] = strings.Trim(prefix + ip[0] + separator + ip2nd, "")
	dbedit["ipaddr_first"] , dbedit["ipaddr"] = ip[0], ip[0]
	dbedit["ipaddr_last"] , dbedit["netmask"] = ip2nd, ip2nd											
						
	if err := tpl.ExecuteTemplate(&out, template, dbedit); err != nil {
		log.Fatalln(err)
	}
	result.WriteString(out.String())		
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}