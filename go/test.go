package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	body, err := ioutil.ReadFile("/home/vagrant/mxcc/html/login.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
