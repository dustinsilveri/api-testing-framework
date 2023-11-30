package modules

import (
	"fmt"
	"net/http"

	"github.com/TwiN/go-color"
)

/*
AUTHOR: 		Dustin Silveri
DESCRIPTION: 	Sends a request for each HTTP Verb, ie. GET, POST, DELETE, etc.
TYPE:			Detection
*/

func DetectMethods(client *http.Client, Endpoint string) {
	// Run and print the responses for each method detection
	fmt.Println("")
	fmt.Println(color.Green + "[*] Starting Methods Detection" + color.Reset)
	fmt.Println("")
	DetectGET(client, Endpoint)
	DetectHEAD(client, Endpoint)
	DetectPOST(client, Endpoint)
	DetectPUT(client, Endpoint)
	DetectDELETE(client, Endpoint)
	DetectOPTIONS(client, Endpoint)
	DetectPATCH(client, Endpoint)
	fmt.Println("")
}

func DetectGET(client *http.Client, Endpoint string) {
	req, err := http.NewRequest("GET", Endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(color.Green + resp.Request.Method + color.Reset + "\t" + resp.Status)
}

func DetectPOST(client *http.Client, Endpoint string) {
	req, err := http.NewRequest("POST", Endpoint, nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println(color.Green + resp.Request.Method + color.Reset + "\t" + resp.Status)
}

func DetectDELETE(client *http.Client, Endpoint string) {
	req, err := http.NewRequest("DELETE", Endpoint, nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println(color.Green + resp.Request.Method + color.Reset + "\t" + resp.Status)
}

func DetectPUT(client *http.Client, Endpoint string) {
	req, err := http.NewRequest("PUT", Endpoint, nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println(color.Green + resp.Request.Method + color.Reset + "\t" + resp.Status)
}

func DetectOPTIONS(client *http.Client, Endpoint string) {
	req, err := http.NewRequest("OPTIONS", Endpoint, nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Warn and change color if server responds with 200 OK
	if resp.StatusCode == 200 {
		fmt.Println(color.Red + resp.Request.Method + "\t" + resp.Status + color.Reset)
	} else {
		fmt.Println(color.Green + resp.Request.Method + color.Reset + "\t" + resp.Status)
	}

}

func DetectPATCH(client *http.Client, Endpoint string) {
	req, err := http.NewRequest("PATCH", Endpoint, nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println(color.Green + resp.Request.Method + color.Reset + "\t" + resp.Status)
}

func DetectHEAD(client *http.Client, Endpoint string) {
	req, err := http.NewRequest("HEAD", Endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println(color.Green + resp.Request.Method + color.Reset + "\t" + resp.Status)
}
