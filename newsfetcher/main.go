package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Newsfetcher: No news! Is good news ;-)")
/*	
	import
	"io/ioutil"

	if r.URL.Path == "/" {
		var q = r.URL.Query()
		var source = q.Get("s")
		var resp *http.Response = nil
		var err error
		if (source == "rte") {
			resp, err = http.Get("http://rte.ie/")
		} else if (source == "it") {
			resp, err = http.Get("http://irishtimes.com/")
		} 
		if (err != nil) {
			fmt.Fprintln(w, "Newsfetcher[ERROR]:" + err.Error())
		} else {
			if (resp != nil) {
				if resp.StatusCode == http.StatusOK {
    				bodyBytes, err2 := ioutil.ReadAll(resp.Body)
					if (err2 != nil) {
						fmt.Fprintln(w, "Newsfetcher[ERROR]:" + err2.Error())
					} else {
						fmt.Fprintln (w, string(bodyBytes))
					}
				} else {
					fmt.Fprintln(w, "Newsfetcher[ERROR]: HTTP returned status " + string(resp.StatusCode))
				}
			} else {
				fmt.Fprintln (w, "Newsfetcher: Unknown news source :-(.")
			}
		}
	} else {
		fmt.Fprintln(w, "Newsfetcher: Hmmm, not sure how to respond to that request.")
	}
*/}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
