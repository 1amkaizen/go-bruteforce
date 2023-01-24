package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// the below function currently returns , either 200 or 404 response , will modify this in future
func adminPanel(uk string) { //defining admin panel bruteforce function
	url := uk
	url = "http://" + uk //appending https:// schema

	colorize(ColorCyan, "Note : Screenshots will be saved in the same directory  ")

	colorize(ColorRed, "Admin Panel BruteForce")
	file, err := os.Open("read.txt") // opening file containing paths
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var i string
	for scanner.Scan() {
		scanner.Text()
		i = scanner.Text()
		var name string
		name = url + scanner.Text()
		resp, err := http.Get(name)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode >= 200 && resp.StatusCode <= 399 {
			colorize(ColorGreen, name) // logging 200 - 399 responses
			var k string
			k = i
			r := strings.NewReplacer(`/`, "-", `/`, "-", ".", "-")
			colorize(ColorYellow, "\n [!] Taking screenshot : ")
			colorize(ColorCyan, " [!] Saving Screenshot as ==> : "+uk+r.Replace(k)+".png")

			////////////////////////////////////////////////////////////////////////////

			var test1 string
			test1 = "http://webshot.okfnlabs.org/api/generate?url=" + name
			colorize(ColorBlue, " [!] Screen shot status  : ")
			resp, err := http.Get(test1)
			// handle the error if there is one
			if err != nil {
				panic(err)
			}
			// do this now so it won't be forgotten
			defer resp.Body.Close()
			// reads html as a slice of bytes
			html, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			// show the nmap as a string %s
			// fmt.Printf("%s\n", html)
			var s string
			s = string(html)
			if s == "Unable to take a screenshot" {
				colorize(ColorRed, " [!] Error , Unable to take a screenshot\n")
			} else {
				response, err := http.Get(test1)
				if err != nil {
					log.Fatal(err)
				}
				defer response.Body.Close()

				// Create output file
				outFile, err := os.Create(uk + r.Replace(k) + ".png")
				if err != nil {
					log.Fatal(err)
				}
				defer outFile.Close()

				// Copy data from HTTP response to file
				_, err = io.Copy(outFile, response.Body)
				if err != nil {
					log.Fatal(err)
				}
				colorize(ColorGreen, " [!] Screenshot Saved successfully\n")
			}

		} else {
			colorize(ColorRed, name) // logging responses other than 200 - 399

		}
		// Print the HTTP Status Code and Status Name
		fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
