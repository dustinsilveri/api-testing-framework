package main

import (
	"api-testing-framework/help"
	"api-testing-framework/modules"
	"api-testing-framework/modules/graphql"
	"api-testing-framework/utils"
	"bufio"
	"net/http"
	"net/url"
	"os/exec"

	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/TwiN/go-color"
)

var Endpoint string
var Proxy = false
var Modules []string
var Templates []string
var inputArray []string

func init() {
	// Create Modules & Templates Arrays for use at runtime.
	Modules = utils.InitArrays(Modules, "modules")
	Templates = utils.InitArrays(Templates, "templates")

	// Opening Banner
	fmt.Println(color.Green + ` ________________________________    `)
	fmt.Println(`< ` + color.Red + `Starting API Testing Framework` + color.Green + ` >   `)
	fmt.Println(` --------------------------------    `)
	fmt.Println(`        \   ^__^                     `)
	fmt.Println(`         \  (oo)\_______             `)
	fmt.Println(`            (__)\       )\/\         `)
	fmt.Println(`                ||----w |            `)
	fmt.Println(`                ||     ||            ` + color.Reset)

	fmt.Printf("\n%v Modules, %v Templates Imported\n", len(Modules), len(Templates))
}

func main() {

	flag.StringVar(&Endpoint, "E", "", "The Endoint to Test") // Tainted Input
	flag.BoolVar(&Proxy, "P", false, "Send all requests through a proxy ie, burp/zap")
	flag.Parse()

	// Usage help with no arguments
	if len(Endpoint) == 0 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Starting Point-->
	framework(Endpoint)
}

func framework(Endpoint string) {

	for {
		result, _ := prompt(Endpoint)
		inputArray = strings.Split(result, " ")

		switch inputArray[0] {

		// help Menu
		case "help":
			help.MainHelp()

		// use options
		case "use":

			switch inputArray[1] {
			case "modules": // all modules go here with arguments
				//result, _ = UsePrompt(Endpoint, "Modules")
				switch inputArray[2] {
				case "modules/detectmethods":
					modules.DetectMethods(httpClient(), Endpoint)

				case "modules/ratelimit":
					modules.RateLimit(Endpoint)

				case "modules/request":
					modules.GET(httpClient(), Endpoint)

				case "modules/graphql/fieldduplicationquery":
					graphql.FieldDuplicationQuery(httpClient(), Endpoint)

				case "modules/graphql/introspection":
					graphql.Introspection(httpClient(), Endpoint)

				case "modules/graphql/resourceintensivequery":
					graphql.ResourceIntesiveQuery(httpClient(), Endpoint)

				case "modules/graphql/aliasedbasedquery":
					graphql.AliasedBasedQuery(httpClient(), Endpoint)

				case "modules/graphql/batchquery":
					graphql.BatchQuery(httpClient(), Endpoint)

				case "modules/graphql/deeprecursionquery":
					graphql.DeepRecursionQuery(httpClient(), Endpoint)

				case "modules/graphql/circularquery":
					graphql.CircularQuery(httpClient(), Endpoint)
				}
			}

		//set options
		case "set":
			switch inputArray[1] {
			case "endpoint":
				Endpoint = inputArray[2] // syntax 'set endpoint http://localhost:5000/user'
			}

		// show options
		case "show":
			switch inputArray[1] {

			case "modules":
				utils.ShowModules(Modules)

			case "templates":
				utils.ShowTemplates(Templates)

			}

		// edit a template file
		case "edit":
			filename := result[5:]

			// Trim off new line chars win/lin
			if runtime.GOOS == "windows" {
				cmd := exec.Command("notepad.exe", filename)
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
			}
			if runtime.GOOS == "linux" {
				utils.ExecEditor()
			}

		// quit or exit
		case "quit":
			fmt.Println("[!] Quitting")
			os.Exit(0)
		}
	}
}

// Standard Prompt
func prompt(Endpoint string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	// Format the Prompt
	fmt.Print(color.White, "ATF (")
	fmt.Print(color.Blue, Endpoint)
	fmt.Print(color.White + ") > " + color.Reset)

	text, err := reader.ReadString('\n')
	if err != nil {
		panic("Something Went Wrong... " + err.Error())
	}

	// Trim off new line chars win/lin
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	} else {
		text = strings.TrimRight(text, "\n")
	}

	return text, err
}

// Not currently used
func UsePrompt(Endpoint string, Module string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	// Format the Prompt
	fmt.Print(color.White, "ATF (")
	fmt.Print(color.Red, Endpoint)
	fmt.Print(color.White + ") (" + color.Reset)
	fmt.Print(color.Yellow + Module + color.Reset)
	fmt.Print(color.White + ") > " + color.Reset)

	text, err := reader.ReadString('\n')
	if err != nil {
		panic("Something Went Wrong... " + err.Error())
	}

	// Trim off new line chars win/lin
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	} else {
		text = strings.TrimRight(text, "\n")
	}

	return text, err
}

// Setup httpClient for use throughout each module
func httpClient() *http.Client {
	client := &http.Client{}
	if Proxy {
		proxyStr := "http://localhost:8080"
		proxyURL, _ := url.Parse(proxyStr)
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}

		client = &http.Client{
			Transport: transport,
		}
	}
	return client
}
