package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(resp)
		os.Exit(1)
	}

	// fmt.Println(*resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println("Response from google: ", string(bs))

	io.Copy(os.Stdout, resp.Body)

}
