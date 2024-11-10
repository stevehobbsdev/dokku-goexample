package main

import (
	"fmt"
	"os"
	"strconv"

	dokku "github.com/dokku/dokku/plugins/common"
	flag "github.com/spf13/pflag"
)

const (
	helpContent = `
    hello <app>, says 'hello <app>'
    hello:world, says 'hello world!'`
)

var appName = flag.String("app", "", "The app name")

func main() {
	// Run the root command
	flag.Usage = usage
	flag.Parse()
	cmd := flag.Arg(0)

	switch cmd {
	case "hello:help":
		usage()
	case "help":
		fmt.Println(helpContent)
	case "hello":
		// Support app name as first argument as well as a flag
		if *appName == "" && len(flag.Args()) > 1 {
			*appName = flag.Args()[1]
		}

		// Verify it
		err := dokku.VerifyAppName(*appName)

		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Hello, %s!\n", *appName)

	default:
		dokkuNotImplementedExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("Error parsing DOKKU_NOT_IMPLEMENTED_EXIT")
			dokkuNotImplementedExitCode = 10
		}
		os.Exit(dokkuNotImplementedExitCode)
	}
}

func usage() {
	fmt.Println("Usage: dokku hello[:world] [options] [arguments]")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("\t--%s: %s (default: %s)\n", f.Name, f.Usage, f.DefValue)
	})
}
