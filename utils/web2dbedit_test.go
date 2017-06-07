package utils

import (
"testing"
"fmt"
)


//func Ip2dbedit(data *string, group string) bytes.Buffer {
func TestIp2dbedit(t *testing.T) {
	output := Ip2dbedit(data, "g-ochepist-test","added by ochecpist with dbedit - test", "../templates/*")
	fmt.Println(output.String())
}

const data = `
1.1.1.1-1.1.1.254
2.2.2.2/255.255.255.0
3.3.3.3

Not an IP address after empty line
`