package main

import (
	"fmt"
	"net"
	"regexp"
        "encoding/json"
        "io/ioutil"
)

func ParseJson(f string) {
  fText, err := ioutil.ReadFile(f)
  if err != nil {
    fmt.Println(err)
  }

  err = json.Unmarshal(fText, &C)
  if err != nil {
    fmt.Println(err)
  }
}

func checkIP(IP string) bool {
	if (net.ParseIP(IP) == nil) {
		return false
	} else {
		return true
	}
}

func CheckStr(str string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		fmt.Println(err)
	}
	newStr := reg.ReplaceAllString(str, "")

	return newStr
}

func checkIntSize(number int) bool {
	if number > 9999 {
		return false
	}

	return true
}

func CheckInput(target, method string, time, port int) bool {
	res := checkIP(target)
	if res != true {
		return false
	}

	res = checkIntSize(port)
	if res != true {
		return false
	}

	return true
}
