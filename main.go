package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func init() {
	flag.Usage = func() {
		h := []string{
			"-d  for domain",
			"-w for wordlists\n",
		}
		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
}

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorCyan         = "\u001b[36m"
	ColorReset        = "\u001b[0m"
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

var d string
var w string

func getHeader(url string) {
	url = "http://" + url
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	colorize(ColorRed, "\n\t\tHEADER INFO")
	fmt.Println("+---------------------------------------+")
	for k, v := range resp.Header {

		fmt.Print(ColorYellow, k)
		fmt.Print(ColorRed, " : ")
		fmt.Println(ColorGreen, v)
	}
}

func main() {
	if len(os.Args) == 1 {
		flag.Usage()
		return
	}
	flag.StringVar(&d, "d", "", "")
	flag.StringVar(&w, "w", "", "")
	flag.Parse()
	adminBF(d)
}
