package main

import (
	"fmt"
	"net/http"
)



func main() {
	url := "http://album.zhenai.com/u/1614063457"
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("fetch err")
	}
	defer resp.Body.Close()
}
