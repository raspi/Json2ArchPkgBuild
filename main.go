package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/raspi/go-PKGBUILD"
	"io/ioutil"
	"os"
)

var VERSION = `v0.0.0`
var BUILD = `dev`
var BUILDDATE = `0000-00-00T00:00:00+00:00`

const AUTHOR = `Pekka Järvinen`
const HOMEPAGE = `https://github.com/raspi/Json2ArchPkgBuild`

func main() {
	generateArg := flag.Bool(`example`, false, `generate example JSON template`)

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, `json2archpkgbuild %s (%s)`+"\n", VERSION, BUILDDATE)
		fmt.Fprintf(os.Stdout, `(c) %s 2020- - <URL: %s >`+"\n", AUTHOR, HOMEPAGE)
		fmt.Fprintln(os.Stdout, `Examples:`)
		fmt.Fprintf(os.Stdout, `  %s <file.json>   -- Print PKGBUILD from JSON file`+"\n", os.Args[0])
		fmt.Fprintf(os.Stdout, `  %s -example      -- Print out example JSON`+"\n", os.Args[0])
	}

	flag.Parse()

	if *generateArg {
		// Generate example
		exampletpl := example()

		b, err := json.MarshalIndent(&exampletpl, ``, `  `)
		if err != nil {
			fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}

		fmt.Fprint(os.Stdout, string(b))

		os.Exit(0)
	}

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stdout, `See -h for help`)
		os.Exit(0)
	}

	fname := flag.Arg(0)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, `error: %v`, err)
		os.Exit(1)
	}

	tpl, err := PKGBUILD.FromJson(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, `error: %v`, err)
		os.Exit(1)
	}

	errs := tpl.Validate()
	if errs != nil {
		fmt.Fprintf(os.Stderr, `error:`+"\n")
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, `  - %v`+"\n", e)
		}

		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, `%s`, tpl)
}
