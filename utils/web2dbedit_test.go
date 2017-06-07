package utils

import (
"testing"
"fmt"
"os"
)


//Testing 
func TestIp2dbedit(t *testing.T) {
	output := Ip2dbedit(data, "g-ochepist-test", "added by ochepist with dbedit - test", "../templates/*")
	fmt.Println(output.String())
	f, _ := os.Create("../results/" + "g-ochepist-test" + "-dbedit.txt")
	defer f.Close()
	b, _ := f.Write(output.Bytes())
	fmt.Printf("wrote %d bytes\n", b)
}

const data = `
1.1.1.1-1.1.1.254
2.2.2.0/255.255.255.0
3.3.3.3

Not an IP address after empty line
`