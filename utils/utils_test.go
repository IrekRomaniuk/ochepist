package utils

import (
	"testing"
	"fmt"
	"crypto/md5"
	"net"
    
)

func TestReadXML(t *testing.T) {
	data, _ := ReadXML(o365IP)
	fmt.Println(data[0].List[0].Address)	
}

func TestGetPage(t *testing.T) {	
	mm, _ := GetPage("https://minemeld/feeds/office365_IPv4s")
	//fmt.Println(mm)
	// Hash received data in order to find if web page content has changed
	fmt.Printf("\nmd5sum is %x\n", md5.Sum([]byte(mm)))
	xml, _ := GetPage("https://support.content.office.net/en-us/static/O365IPAddresses.xml")
	ms, _ := ReadXML(xml)
	fmt.Println(ms[0].List[1])
}

func TestGetHash(t *testing.T) {
	hash, _ := GetHash("../results/hash")
	fmt.Println(hash)
}

func TestMaskDot(t *testing.T) {
	_, ipv4Net, _ := net.ParseCIDR("10.0.0.0/8")
	fmt.Println(MaskDot(ipv4Net.Mask))
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