package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	yadisk "../"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s token\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		fmt.Println("\nError:\n",
			"You must provide an OAuth token.\n",
			"See https://tech.yandex.ru/disk/rest/\n",
			"and https://tech.yandex.ru/oauth/\n",
			"") // last line is for the linter
		os.Exit(2)
		return
	}

	oAuthToken := flag.Arg(0)
	fmt.Printf("Using token \"%s\"...\n", oAuthToken)

	api := yadisk.NewAPI(oAuthToken)

	reader := strings.NewReader("TEST")
	err := api.Upload(reader, "app:/test.txt", true)
	if err != nil {
		log.SetPrefix("error: ")
		log.Println(err)
	}

	fmt.Println("Success!")
}
