package help

import "fmt"

func MainHelp() {
	fmt.Println("\nCore Commands")
	fmt.Println("=============")
	fmt.Println("")

	fmt.Println("Commamnd\t\t\tDescription")
	fmt.Println("--------\t\t\t-----------")
	fmt.Println("?/help\t\t\t\tHelp menu")
	fmt.Println("quit\t\t\t\tExit")
	fmt.Println("set endpoint <URL>\t\tChange the current endpoint")
	fmt.Println("show templates\t\t\tShow all templates")
	fmt.Println("show modules\t\t\tShow all moodules available")
	//fmt.Println("SET USERAGENT <USERAGENT>\tChange the useragent in the requests. ie 'api-testing-framework'")
	//fmt.Println("SHOW ENDPOINT\t\t\tShow the current endpoint.")
	//fmt.Println("GET\t\t\t\tSend a GET request to the current endpoint.")
	fmt.Println("")

	fmt.Println("Attacks\t\t\t\tDescription")
	fmt.Println("--------\t\t\t-----------")
	fmt.Println("use ratelimit\t\t\tDetect if ratelimiting is in place on the endpoint")
	fmt.Println("use methods\t\t\tDetect all allowed methods")
	fmt.Println("use introspection\t\tDetect and run an introspection query on a GraphQL endpoint")
	fmt.Println("use request\t\t\tSends a standard GET request and returns the response and headers")
	fmt.Println("")

}
