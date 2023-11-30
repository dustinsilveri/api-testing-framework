package graphql

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"api-testing-framework/utils"

	"github.com/TwiN/go-color"
)

/*
AUTHOR: 		Dustin Silveri
DESCRIPTION:
TYPE:			Detection
*/

func BatchQuery(client *http.Client, Endpoint string) {
	data, _ := utils.ReadFile("templates/graphql/batchquery.json")
	start := time.Now().UnixNano() / int64(time.Millisecond)

	// The GraphQL queries to test
	//var jsonStr = []byte(`{"query":"query{__schema{subscriptionType{fields{name}}}}","variables":{},"operationName":null}`)

	req, err := http.NewRequest("POST", Endpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	// Set custom header if needed.
	//req.Header.Set("X-Custom-Header", "myvalue")
	if err != nil {
		panic(err)
	}

	// send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// Format the output to the screen
	fmt.Println("")
	fmt.Println(color.Green + "[*] Sending the Batch Query..." + color.Reset)
	fmt.Println(color.Green + "Request: " + color.Reset)
	for name, headers := range req.Header {
		for _, hdr := range headers {
			println(color.Green + name + color.Reset + ": " + hdr)
		}
	}
	fmt.Println("")

	fmt.Println(color.Green + "Response: " + color.Reset)
	fmt.Println(color.Green+"Status:"+color.Reset, resp.Status)
	// Iterate over Response headers
	for name, headers := range resp.Header {
		for _, hdr := range headers {
			println(color.Green + name + color.Reset + ": " + hdr)
		}
	}

	b, _ := io.ReadAll(resp.Body)
	p, _ := utils.PrettyPrint(b)
	fmt.Printf("\n\n")
	fmt.Printf(color.White+"%s"+color.Reset, p)

	diff := time.Now().UnixNano()/int64(time.Millisecond) - start
	fmt.Printf(color.Green+"\n[!] Request sent in %v ms.\n"+color.Reset, diff)
	fmt.Println("") // New line for formatting.
}
