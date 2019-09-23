package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 1000000; i++ {
		<-time.After(time.Millisecond * 3)
		resp, err := http.Get("http://localhost:8086/match")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Hit the server and read the responses ", string(body))
	}
}
