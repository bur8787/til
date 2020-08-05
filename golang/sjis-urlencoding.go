package main

import (
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"net/url"
)

func main() {
	// UTF-8
	v := url.Values{}
	v.Set("key", "あ")
	fmt.Printf("%s\n", v.Encode()) // key=%E3%81%82

	// SJIS
	sjisV, _, _ := transform.String(japanese.ShiftJIS.NewEncoder(), "あ")
	v2 := url.Values{}
	v2.Set("key", sjisV)
	fmt.Printf("%s\n", v2.Encode()) // key=%82%A0
}
