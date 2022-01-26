package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
)

func main() {
	urlPtr := flag.String("url", "", "A single url to decode")

	flag.Parse()

	split_url_1 := strings.Split(*urlPtr, "outlook.com/?url=")
	split_url_2 := strings.Split(split_url_1[1], "&data")

	decoded_url, err := url.QueryUnescape(split_url_2[0])
	if err != nil {
		fmt.Println("URL Decoding Error: {}", err)
	}
	fmt.Println("Decoded URL: ", decoded_url)
}
