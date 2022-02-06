package main

import (
	"WebFramework/project"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	flag.Parse()
	parameters := flag.Args()

	switch parameters[0] {
	case "runserver":
		server := http.Server{
			Addr: project.GetHostsAddress(),
		}

		fmt.Println(project.GetHostsAddress())

		server.ListenAndServe()

	default:
		fmt.Println("go run main.go " + parameters[0] + ": unknown command")
	}
}
