package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"fmt"
	"encoding/xml"
)

type Products struct {
	XMLName xml.Name `xml:"products"`
	Product []Products `xml:"product"`
}

type Product struct {
	XMLName xml.Name `xml:"product"`
	Name string `xml:"name,attr"`
	List []List `xml:"addresslist"`
}

type List struct {
	XMLName xml.Name `xml:"addresslist"`
	Type string `xml:"type,attr"`
}

// ReadXML unmarshals xml
func ReadXML(htmlData string) (Product, error) {
    var products Products
    if err := xml.Unmarshal([]byte(htmlData), &products); err != nil {
        return "", err
    }
    return products.Product, nil
}

// GetPage from url and return body as string
func GetPage(url string) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	return string(htmlData), nil
}
// ValidIP4 checks IP address for validity
func ValidIP4(ip string) bool {
	ip = strings.Trim(ip, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ip) {
		return true
	}
	return false
}
// MaskDot converts netamsk to dot notation
func MaskDot(m []byte) string {
	if len(m) != 4 {
		panic("ipv4Mask: len must be 4 bytes")
	}

	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}
// GetHash reads hash from file
func GetHash(file string) (string, error) {
	b, err := ioutil.ReadFile(file) 
    if err != nil {
        return "", err
	}
	return string(b), nil
}

// SetHash writes hash to file
func SetHash(hash, file string) error {
	err := ioutil.WriteFile(file,[]byte(hash), 0644)
	return err
}