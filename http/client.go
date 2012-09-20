package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func check_health(x int, c chan string) (bool, error)  {
	response, err := http.Get("http://127.0.0.1:8888/health?"+ strconv.Itoa(x))
	if err != nil {
		fmt.Println("Error: ", err)
		return false, err
	}
	if response.StatusCode != 200 {
		fmt.Println("StatusCode: ", response.Status)
		return false, nil
	}
	c <- response.Status
	return true, nil
}

func main() {
	channel := make(chan string)
	for i := 0; i < 130; i++ {
		go check_health(i, channel)
	}
	<- channel
	fmt.Println("Done")
}
