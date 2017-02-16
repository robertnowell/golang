package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// func main() {
// 	if len(os.Args) == 1{
// 		fmt.Println("please include a url")
// 	}
// 	for _, url := range os.Args[1:] {
// 		if !strings.HasPrefix(url, "http://"){
// 			url = "http://" + url
// 		}
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n, err")
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)
// 		stat := resp.Status
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("STATUS:%s\n\nBODY:%s\n", stat, b)
// 	}
// }

func main() {
	if len(os.Args) == 1{
		fmt.Println("please include a url")
	}
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://"){
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n, err")
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
