package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error when fetching data: ", err)
		os.Exit(1)
	}
	//fmt.Println("Response fetch data: ", resp)

	bs := make([]byte, 999999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}
