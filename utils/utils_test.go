package utils

import (
	"testing"
	"fmt"
	"crypto/md5"
)

func TestGetPage(t *testing.T) {	
	data, _ := GetPage("https://minemeld/feeds/office365_IPv4s")
	fmt.Println(data)
	// Hash received data in order to find if web page content has changed
	fmt.Printf("\nmd5sum is %x\n", md5.Sum([]byte(data)))
	data, _ = ReadXML("https://support.content.office.net/en-us/static/O365IPAddresses.xml")
	fmt.Println(data)
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