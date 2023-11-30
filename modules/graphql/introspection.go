package graphql

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"api-testing-framework/utils"

	"github.com/TwiN/go-color"
)

/*
AUTHOR: 		Dustin Silveri
DESCRIPTION: 	Sends an introspection query to the GraphQL endpoint, using the introspection.json template.  This can be modified if needed.
TYPE:			Detection
*/

func Introspection(client *http.Client, Endpoint string) {
	data, _ := utils.ReadFile("templates/graphql/introspectionquery.json")

	// The GraphQL queries to test
	//var jsonStr = []byte(`{"query":"\n    query IntrospectionQuery {\n      __schema {\n        \n        queryType { name }\n        mutationType { name }\n        subscriptionType { name }\n        types {\n          ...FullType\n        }\n        directives {\n          name\n          description\n          \n          locations\n          args {\n            ...InputValue\n          }\n        }\n      }\n    }\n\n    fragment FullType on __Type {\n      kind\n      name\n      description\n      \n      fields(includeDeprecated: true) {\n        name\n        description\n        args {\n          ...InputValue\n        }\n        type {\n          ...TypeRef\n        }\n        isDeprecated\n        deprecationReason\n      }\n      inputFields {\n        ...InputValue\n      }\n      interfaces {\n        ...TypeRef\n      }\n      enumValues(includeDeprecated: true) {\n        name\n        description\n        isDeprecated\n        deprecationReason\n      }\n      possibleTypes {\n        ...TypeRef\n      }\n    }\n\n    fragment InputValue on __InputValue {\n      name\n      description\n      type { ...TypeRef }\n      defaultValue\n      \n      \n    }\n\n    fragment TypeRef on __Type {\n      kind\n      name\n      ofType {\n        kind\n        name\n        ofType {\n          kind\n          name\n          ofType {\n            kind\n            name\n            ofType {\n              kind\n              name\n              ofType {\n                kind\n                name\n                ofType {\n                  kind\n                  name\n                  ofType {\n                    kind\n                    name\n                  }\n                }\n              }\n            }\n          }\n        }\n      }\n    }\n  ","variables":{},"operationName":"IntrospectionQuery"}`)

	req, err := http.NewRequest("POST", Endpoint, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	// Set custom header if needed.
	// req.Header.Set("X-Custom-Header", "blah")
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
	fmt.Println("[*] Sending the Introspection Query...")
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

	//fmt.Println("response Body:", string(body))

	b, _ := io.ReadAll(resp.Body)
	p, _ := utils.PrettyPrint(b)
	fmt.Printf("\n\n")
	fmt.Printf(color.White+"%s"+color.Reset, p)
	fmt.Println("") // New line for formatting.
}
