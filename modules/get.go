package modules

import (
	"fmt"
	"io"
	"net/http"

	"github.com/TwiN/go-color"
)

/*
AUTHOR: 		Dustin Silveri
DESCRIPTION: 	Sends a Standard GET request to the target.
TYPE:			Detection
*/

func GET(client *http.Client, Endpoint string) {

	req, err := http.NewRequest("GET", Endpoint, nil)
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Iterate over headers
	fmt.Println("")
	fmt.Println(color.Green + "[*] Sending GET request to endpoint" + color.Reset)

	fmt.Println(color.Green + "Response Headers:" + color.Reset)
	for name, headers := range resp.Header {
		for _, hdr := range headers {
			println(color.Green + name + ": " + color.Reset + hdr)
		}
	}
	fmt.Printf("\n\n")

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	} else {
		fmt.Printf(color.Green+"Status Code: "+color.Reset+"%d\n\n", resp.StatusCode)
	}
}
