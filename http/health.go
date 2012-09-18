package main

import (
	"fmt"
	"net/http"
	"runtime"
    "strconv"
)

func check_health(x int, c chan int) (bool, error)  {
	runtime.Gosched()

	response, err := http.Get("https://www.google.com/?"+ strconv.Itoa(x))
	if err != nil {
		fmt.Println("Error: ", err)
		return false, err
	}
	if response.StatusCode != 200 {
		fmt.Println("StatusCode: ", response.Status)
		return false, nil
	}
	fmt.Println(x, " - ", response.Status)
	c <- 0
	return true, nil
}

func main() {
	channel := make(chan int)
	for i := 0; i < 20; i++ {
		go check_health(i, channel)
	}
	<- channel
	fmt.Println("Done")
}
