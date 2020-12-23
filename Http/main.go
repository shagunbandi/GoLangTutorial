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
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// Raw Response
	// fmt.Println(resp)

	// Using Read Interface
	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// Using io.Copy
	io.Copy(os.Stdout, resp.Body)

	// Using Custom Writer - LogWriter
	// lw := logWriter{}
	// io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (i int, err error) {
	fmt.Println(string(bs))
	fmt.Printf("Wrote %v bytes\n", len(bs))
	return len(bs), nil
}
