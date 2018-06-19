package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

type customWriter struct {}

func main(){
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/

	cw := customWriter{}

	io.Copy(cw, resp.Body)

}

func (customWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes")
	return len(bs), nil
}