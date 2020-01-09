package main

// This is used for -example argument output

import "github.com/raspi/go-PKGBUILD"

func example() PKGBUILD.Template {
	sources := map[string][]PKGBUILD.Source{
		`x86_64`: {
			PKGBUILD.Source{
				URL: "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-amd64.tar.gz",
				Checksums: map[string]string{
					`sha256`: `de3edfb94d5d0ae3d027c6c743e27290fa0500da4777da57154f2acab52775bf`,
				},
			},
		},
		`ppc64le`: {
			PKGBUILD.Source{
				URL: "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-ppc64le.tar.gz",
				Checksums: map[string]string{
					`sha256`: "6baef7ee046ceb4450e703a87f05fa5662708d4c3562c26abb427d34b4c82819",
				},
			},
		},
		`aarch64`: {
			PKGBUILD.Source{
				URL: "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-arm64.tar.gz",
				Checksums: map[string]string{
					`sha256`: "11d2b36d6b320dfee489d475635b53206b59288537554ea8bc24f97d06139d64",
				},
			},
		},
		`arm`: {
			PKGBUILD.Source{
				URL: "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-arm.tar.gz",
				Checksums: map[string]string{
					`sha256`: "5e79210655a9a71a7b77a3168194e9ead024a120182fa8560348a24dc87da159",
				},
			},
		},
		`ppc64`: {
			PKGBUILD.Source{
				URL: "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-ppc64.tar.gz",
				Checksums: map[string]string{
					`sha256`: "f744e32caf67a609aa435df9f8c519460b1856f7968c057e6ba61397cf79ec15",
				},
			},
		},
	}

	cmds := PKGBUILD.Commands{
		Prepare: []string{
			`echo foo >> main.c`,
		},
		Build: []string{
			`make`,
		},
		Test: []string{
			`make test`,
		},
		Install: []string{
			`cd "$srcdir"`,
			`install -Dm644 "LICENSE" -t "$pkgdir/usr/share/licenses/$pkgname"`,
			`install -Dm644 "README.md" -t "$pkgdir/usr/share/doc/$pkgname"`,
			`install -Dm755 "bin/$pkgname" -t "$pkgdir/usr/bin"`,
		},
	}

	deps := map[string]PKGBUILD.Depends{
		``: {
			Packages:      []string{"example-core"},
			BuildPackages: []string{"example-dev"},
			TestPackages:  []string{"example-test"},
		},
		`x86_64`: {
			Packages: []string{"example-core-x86"},
		},
	}

	optional := map[string][]PKGBUILD.OptionalPackage{
		``: {
			PKGBUILD.OptionalPackage{
				Package: "php",
				Reason:  "because PHP is EPIC!",
			},
		},
	}

	options := []string{
		`!strip`,
		`docs`,
		`libtool`,
		`staticlibs`,
		`emptydirs`,
		`!zipman`,
		`!ccache`,
		`!distcc`,
		`!buildflags`,
		`makeflags`,
		`!debug`,
	}

	tpl := PKGBUILD.New(sources, cmds, deps, optional, options)

	tpl.Maintainer = `John Doe`
	tpl.MaintainerEmail = `jd@example.org`
	tpl.Name = []string{`exampleapp`}
	tpl.Version = `v1.0.0`
	tpl.Licenses = []string{"Apache 2.0"}
	tpl.ShortDescription = `my example application`
	tpl.URL = `https://github.com/examplerepo/exampleapp`
	tpl.Install = `$pkgname.install`

	return tpl
}
