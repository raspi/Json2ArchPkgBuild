package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/raspi/go-PKGBUILD"
	"io/ioutil"
	"os"
	"time"
)

var VERSION = `v0.0.0`
var BUILD = `dev`
var BUILDDATE = `0000-00-00T00:00:00+00:00`

const AUTHOR = `Pekka Järvinen`
const HOMEPAGE = `https://github.com/raspi/Json2ArchPkgBuild`

func main() {
	generateArg := flag.Bool(`example`, false, `generate example JSON template`)
	nowEpochArg := flag.Bool(`now`, false, `use current time as reference $epoch`)
	increaseReleaseArg := flag.Bool(`incr`, false, `increase $pkgrel`)

	nameArg := flag.String(`name`, ``, `package name`)
	versionArg := flag.String(`ver`, ``, `package version`)
	jsonArg := flag.String(`json`, ``, `output newly generated JSON to file`)

	cmdInstallArg := flag.String(`install`, ``, `install script file path`)
	cmdPrepareArg := flag.String(`prepare`, ``, `prepare script file path`)
	cmdBuildArg := flag.String(`build`, ``, `build script file path`)
	cmdTestArg := flag.String(`test`, ``, `test script file path`)

	checksumFileArg := flag.String(`sums`, ``, `Use checksum file as reference`)
	checksumTypeArg := flag.String(`t`, `sha256`, `Checksum file type (sha1, sha224, sha256, sha384, sha512, b2, md5)`)

	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stdout, `json2archpkgbuild - convert JSON to Arch Linux PKGBUILD - %s (%s)`+"\n", VERSION, BUILDDATE)
		_, _ = fmt.Fprintf(os.Stdout, `(c) %s 2020- - <URL: %s >`+"\n", AUTHOR, HOMEPAGE)
		_, _ = fmt.Fprintln(os.Stdout, `Parameters:`)

		flag.VisitAll(func(f *flag.Flag) {
			_, _ = fmt.Fprintf(os.Stdout, "  -%s\n      %s (default: %q)\n", f.Name, f.Usage, f.DefValue)
		})

		_, _ = fmt.Fprintln(os.Stdout, `Examples:`)
		_, _ = fmt.Fprintf(os.Stdout, `  %s <file.json>`+"\n", os.Args[0])
		_, _ = fmt.Fprintf(os.Stdout, `  %s -example`+"\n", os.Args[0])
		_, _ = fmt.Fprintf(os.Stdout, `  %s -install install.sh app.json`+"\n", os.Args[0])
	}

	flag.Parse()

	if *generateArg {
		// Generate example
		exampletpl := example()

		b, err := json.MarshalIndent(&exampletpl, ``, `  `)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}

		_, _ = fmt.Fprint(os.Stdout, string(b))

		os.Exit(0)
	}

	if flag.NArg() == 0 {
		_, _ = fmt.Fprintln(os.Stdout, `See -h for help`)
		os.Exit(0)
	}

	fname := flag.Arg(0)

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
		os.Exit(1)
	}

	tpl, err := PKGBUILD.FromJson(b)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
		os.Exit(1)
	}

	if *nameArg != `` {
		if len(tpl.Name) == 0 {
			tpl.Name = append(tpl.Name, *nameArg)
		} else {
			tpl.Name[0] = *nameArg
		}
	}

	if *versionArg != `` {
		tpl.Version = *versionArg
	}

	if *cmdInstallArg != `` {
		tpl.Commands.Install, err = PKGBUILD.GetLinesFromFile(*cmdInstallArg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}
	}

	if *cmdPrepareArg != `` {
		tpl.Commands.Prepare, err = PKGBUILD.GetLinesFromFile(*cmdPrepareArg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}
	}

	if *cmdBuildArg != `` {
		tpl.Commands.Build, err = PKGBUILD.GetLinesFromFile(*cmdBuildArg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}
	}

	if *cmdTestArg != `` {
		tpl.Commands.Test, err = PKGBUILD.GetLinesFromFile(*cmdTestArg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}
	}

	if *nowEpochArg {
		tpl.ReleaseTime = time.Now()
	}

	if *increaseReleaseArg {
		tpl.Release++
	}

	if *checksumFileArg != `` {
		ctype := PKGBUILD.Sha256

		switch *checksumTypeArg {
		case `sha1`:
			ctype = PKGBUILD.Sha1
		case `sha224`:
			ctype = PKGBUILD.Sha224
		case `sha256`:
			ctype = PKGBUILD.Sha256
		case `sha384`:
			ctype = PKGBUILD.Sha384
		case `sha512`:
			ctype = PKGBUILD.Sha512
		case `b2`:
			ctype = PKGBUILD.B2
		case `md5`:
			ctype = PKGBUILD.Md5
		default:
			_, _ = fmt.Fprintf(os.Stderr, `unknown type: %q`, *checksumTypeArg)
			os.Exit(1)
		}

		tpl.Files = PKGBUILD.GetChecksumsFromFile(ctype, *checksumFileArg, tpl.DefaultChecksumFilesFunc)
	}

	errs := tpl.Validate()
	if errs != nil {
		_, _ = fmt.Fprintf(os.Stderr, `error:`+"\n")
		for _, e := range errs {
			_, _ = fmt.Fprintf(os.Stderr, `  - %v`+"\n", e)
		}

		os.Exit(1)
	}

	if *jsonArg != `` {
		f, err := os.Create(*jsonArg)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}
		defer f.Close()

		jb, err := json.MarshalIndent(&tpl, ``, `  `)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, `error: %v`, err)
			os.Exit(1)
		}

		f.Write(jb)
	}

	_, _ = fmt.Fprint(os.Stdout, tpl)
}
