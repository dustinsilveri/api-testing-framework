package modules

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TwiN/go-color"
)

/*
AUTHOR: 		Dustin Silveri
DESCRIPTION: 	Tests an endpoint for rate limiting.  Note: Proxy traffic is disabled for this as it degrades performance.
TYPE:			Detection
*/

var TotalRequests int
var NewStatusCode int
var NewContentLength int64

func RateLimit(Endpoint string) {
	TotalRequests := 0
	fmt.Println(color.Green + "[*] Starting Rate Limit Detection" + color.Reset)
	fmt.Println(color.Green + "[*] These requests will not go through the proxy due to performance issues..." + color.Reset)

	fmt.Printf(color.Green + "[*] Sending initial request for baseline\n\n" + color.Reset)

	// Calculate initial time baseline
	start := time.Now().UnixNano() / int64(time.Millisecond)

	resp, err := http.Get(Endpoint)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	TotalRequests += 1

	diff := time.Now().UnixNano()/int64(time.Millisecond) - start

	// If all is well proceed further
	if resp.StatusCode != 405 {
		InitialStatusCode := resp.StatusCode
		InitialContentLength := resp.ContentLength
		InitialResponseTime := diff
		fmt.Printf(color.Green+"Status Code:"+color.Reset+"\t\t%v\n", InitialStatusCode)
		fmt.Printf(color.Green+"Content-Lenth:"+color.Reset+"\t\t%v\n", InitialContentLength)
		fmt.Printf(color.Green+"Response Time (ms):"+color.Reset+"\t%v\n\n", InitialResponseTime)

		// Send first set of requests
		Round := 10
		fmt.Printf(color.Green+"[*] Sending %v requests\n"+color.Reset, Round)
		for i := 1; i < Round; i++ {
			resp, err := http.Get(Endpoint)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			NewStatusCode = resp.StatusCode
			NewContentLength = resp.ContentLength
			TotalRequests += 1

			if (NewStatusCode != InitialStatusCode) || (InitialContentLength != NewContentLength) {
				break
			}
		}
		diff = time.Now().UnixNano()/int64(time.Millisecond) - start
		fmt.Printf("%v requests sent in %v ms.\n\n", TotalRequests, diff)

		if diff > (InitialResponseTime * int64(Round)) {
			fmt.Println(color.Green + "[!] Rate limiting looks like it is in place." + color.Reset)
			return
		}

		// Sending next set of requests
		Round = 51
		fmt.Printf(color.Green+"[*] Sending %v requests\n"+color.Reset, Round)
		for i := 1; i < Round; i++ {
			resp, err := http.Get(Endpoint)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			NewStatusCode = resp.StatusCode
			NewContentLength = resp.ContentLength
			TotalRequests += 1

			if (NewStatusCode != InitialStatusCode) || (InitialContentLength != NewContentLength) {
				break
			}

		}
		diff = time.Now().UnixNano()/int64(time.Millisecond) - start
		fmt.Printf("%v requests sent in %v ms.\n\n", TotalRequests, diff)

		if diff > (InitialResponseTime * int64(Round)) {
			fmt.Println(color.Green + "[!] Rate limiting looks like it is in place." + color.Reset)
			return
		}

		// Sending next set of requests
		Round = 201
		fmt.Printf(color.Green+"[*] Sending %v requests\n"+color.Reset, Round)
		for i := 1; i < Round; i++ {
			resp, err := http.Get(Endpoint)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			NewStatusCode = resp.StatusCode
			NewContentLength = resp.ContentLength
			TotalRequests += 1

			if (NewStatusCode != InitialStatusCode) || (InitialContentLength != NewContentLength) {
				break
			}

		}
		diff = time.Now().UnixNano()/int64(time.Millisecond) - start
		fmt.Printf("%v requests sent in %v ms.\n", TotalRequests, diff)

		if diff > (InitialResponseTime * int64(Round)) {
			fmt.Println(color.Green + "[!] Rate limiting looks like it is in place." + color.Reset)
			return
		}

		// Sending next set of requests
		Round = 741
		fmt.Printf(color.Green+"\n[*] Sending %v requests\n"+color.Reset, Round)
		for i := 1; i < Round; i++ {
			resp, err := http.Get(Endpoint)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			NewStatusCode = resp.StatusCode
			NewContentLength = resp.ContentLength
			TotalRequests += 1

			if (NewStatusCode != InitialStatusCode) || (InitialContentLength != NewContentLength) {
				break
			}

		}
		diff = time.Now().UnixNano()/int64(time.Millisecond) - start
		fmt.Printf("%v requests sent in %v ms.\n", TotalRequests, diff)

		// Outcome Results
		if diff > (InitialResponseTime * int64(Round)) {
			fmt.Println(color.Green + "[!] Rate limiting looks like it is in place." + color.Reset)
			return
		} else {
			fmt.Printf(color.Red + "\n[!] Rate limiting does not look like it is in place.\n" + color.Reset)
		}

	}
}
