package utils

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)
//Get Page from url and return body as string
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

func ValidIP4(ip string) bool {
	ip = strings.Trim(ip, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ip) {
		return true
	}
	return false
}