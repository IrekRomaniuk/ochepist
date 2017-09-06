package utils

import (
	"testing"
	"fmt"
	"crypto/md5"
	"net"
    
)


//Testing GetPage
func TestGetPage(t *testing.T) {
	
	data, _ := GetPage("https://minemeld/feeds/office365_IPv4s")
	
	fmt.Println(data)
	// Hash received data in order to find if web page content has changed
	fmt.Printf("\nmd5sum is %x\n", md5.Sum([]byte(data)))
}

func TestGetHash(t *testing.T) {
	hash, _ := GetHash("../results/hash")
	fmt.Println(hash)
}

func TestMaskDot(t *testing.T) {
	_, ipv4Net, _ := net.ParseCIDR("10.0.0.0/8")
	fmt.Println(MaskDot(ipv4Net.Mask))
}