package utils

import (
	"testing"
	"fmt"
	"io/ioutil"
	"time"
)

func TestGetPage(t *testing.T) {
	const layout = "Jan 2"
	suf := time.Now()
    res := suf.Format(layout)	
	data1, _ := GetPage("https://minemeld/feeds/office365_IPv4s")
	//fmt.Println(data1)
	ioutil.WriteFile("./data/mm-"+res+".txt", []byte(data1), 0644)
	data2, _ := ReadXML("https://support.content.office.net/en-us/static/O365IPAddresses.xml")
	//fmt.Println(data2)
	ioutil.WriteFile("./data/ms-"+res+".txt", []byte(data2), 0644)
}

func TestGetHash(t *testing.T) {
	hash, _ := GetHash("../results/hash")
	fmt.Println(hash)
}



const o365IP = `<products updated="9/1/2017">
<product name="o365">
<addresslist type="IPv4">
<address>13.64.196.27/32</address>
<address>13.64.198.19/32</address>
<address>13.64.198.97/32</address>
</addresslist>
</product>
<product name="OneNote">
<addresslist type="URL">
<address>*.onenote.com</address>
<address>*.msecnd.net</address>
<address>*.microsoft.com</address>
<address>*.office.net</address>
</addresslist>
</product>
</products>`