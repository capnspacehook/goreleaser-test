package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"rsc.io/quote"
)

const projectName = "Go Project Template"

var printVersion bool

func usage() {
	fmt.Fprintf(os.Stderr, `
<Project description>

	<binary name> [flags]
	
<Project details/usage>

%s accepts the following flags:

`[1:], projectName)
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, `

For more information, see https://github.com/<user>/<repo>.
`[1:])
}

func init() {
	flag.Usage = usage
	flag.BoolVar(&printVersion, "version", false, "print version and build information and exit")
}

func main() {
	os.Exit(mainRetCode())
}

func mainRetCode() int {
	flag.Parse()

	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Println("build information not found")
		return 1
	}

	if printVersion {
		printVersionInfo(info)
		return 0
	}

	fmt.Println(quote.Glass())

	return 0
}
