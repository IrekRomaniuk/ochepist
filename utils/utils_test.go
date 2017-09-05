package utils

import (
	"testing"
	"fmt"
	"crypto/md5"
    
)


//Testing GetPage
func TestGetPage(t *testing.T) {
	
	data, _ := GetPage("https://minemeld/feeds/office365_IPv4s")
	
	fmt.Println(data)
	// Hash received data in order to find if web page content has changed
	fmt.Printf("\nmd5sum is %x\n", md5.Sum([]byte(data)))
}