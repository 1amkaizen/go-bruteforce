package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func adminBF(domain string) {
	url := domain
	url = "http://" + domain + "/"

	file, err := os.Open(w)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	fmt.Println(ColorRed, "SCANNIN1G...")
	for sc.Scan() {
		sc.Text()

		var name string
		name = url + sc.Text()
		resp, err := http.Get(name)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode >= 200 && resp.StatusCode <= 399 {
			fmt.Println(ColorGreen, name, resp.StatusCode, ColorReset)
		}

	}
}
