![GitHub All Releases](https://img.shields.io/github/downloads/raspi/Json2ArchPkgBuild/total?style=for-the-badge)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/raspi/Json2ArchPkgBuild?style=for-the-badge)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/raspi/Json2ArchPkgBuild?style=for-the-badge)

# Json2ArchPkgBuild
Generate Arch Linux PKGBUILD files from JSON.

Uses https://github.com/raspi/go-PKGBUILD as library. See `man PKGBUILD`.

```
json2archpkgbuild - convert JSON to Arch Linux PKGBUILD - v0.0.4 (2020-01-10T05:09:05+02:00)
(c) Pekka Järvinen 2020- - <URL: https://github.com/raspi/Json2ArchPkgBuild >
Parameters:
  -build
      build script file path (default: "")
  -example
      generate example JSON template (default: "false")
  -incr
      increase $pkgrel (default: "false")
  -install
      install script file path (default: "")
  -json
      output newly generated JSON to file (default: "")
  -name
      package name (default: "")
  -now
      use current time as reference $epoch (default: "false")
  -prepare
      prepare script file path (default: "")
  -sums
      Use checksum file as reference (default: "")
  -t
      Checksum file type (sha1, sha224, sha256, sha384, sha512, b2, md5) (default: "sha256")
  -test
      test script file path (default: "")
  -ver
      package version (default: "")
Examples:
  ./json2archpkgbuild <file.json>
  ./json2archpkgbuild -example
  ./json2archpkgbuild -install install.sh app.json
```
## Why?

So you don't need to learn PKGBUILD file syntax. JSON is also easier to update with different tools.

## Example

See [Makefile](Makefile) `ldistro-arch` and [release/linux/arch](release/linux/arch) for usage example.

Generate example JSON which you can use to for generating PKGBUILD package:

    % json2archpkgbuild -example > my-package.json
    $EDITOR my-package.json 

Generate PKGBUILD    

    % json2archpkgbuild my-package.json > PKGBUILD

Increase `$pkgrel` and get updated JSON file

    % json2archpkgbuild -incr -json updated-package.json my-package.json > PKGBUILD

Inject `install.sh` file to output

    % json2archpkgbuild -install install.sh my-package.json > PKGBUILD

Use checksums as source files (`pkg_url_prefix` is used for prefixing the files)

    % json2archpkgbuild -sums checksums.sha256 my-package.json > PKGBUILD


## Example JSON
 
 Use the `-example` parameter as this might be old.
 
 ```json
 {
  "_meta": {
    "ver": "v1.0.0"
  },
  "maintainer": "John Doe",
  "maintainer_email": "jd@example.org",
  "name": [
    "exampleapp"
  ],
  "version": "v1.0.0",
  "release": 1,
  "release_time": "1970-01-01T02:00:00+02:00",
  "short_description": "my example application",
  "licenses": [
    "Apache 2.0"
  ],
  "url": "https://github.com/examplerepo/exampleapp",
  "changelog_file": "",
  "groups": null,
  "dependencies": {
    "": {
      "packages": [
        "example-core"
      ],
      "build_packages": [
        "example-dev"
      ],
      "test_packages": [
        "example-test"
      ]
    },
    "x86_64": {
      "packages": [
        "example-core-x86"
      ]
    }
  },
  "optional_packages": {
    "": [
      {
        "package": "php",
        "reason": "because PHP is EPIC!"
      }
    ]
  },
  "provides": null,
  "options": [
    "!strip",
    "docs",
    "libtool",
    "staticlibs",
    "emptydirs",
    "!zipman",
    "!ccache",
    "!distcc",
    "!buildflags",
    "makeflags",
    "!debug"
  ],
  "install": "$pkgname.install",
  "files": {
    "aarch64": [
      {
        "url": "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-arm64.tar.gz",
        "checksums": {
          "sha256": "11d2b36d6b320dfee489d475635b53206b59288537554ea8bc24f97d06139d64"
        }
      }
    ],
    "arm": [
      {
        "url": "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-arm.tar.gz",
        "checksums": {
          "sha256": "5e79210655a9a71a7b77a3168194e9ead024a120182fa8560348a24dc87da159"
        }
      }
    ],
    "ppc64": [
      {
        "url": "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-ppc64.tar.gz",
        "checksums": {
          "sha256": "f744e32caf67a609aa435df9f8c519460b1856f7968c057e6ba61397cf79ec15"
        }
      }
    ],
    "ppc64le": [
      {
        "url": "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-ppc64le.tar.gz",
        "checksums": {
          "sha256": "6baef7ee046ceb4450e703a87f05fa5662708d4c3562c26abb427d34b4c82819"
        }
      }
    ],
    "x86_64": [
      {
        "url": "https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-amd64.tar.gz",
        "checksums": {
          "sha256": "de3edfb94d5d0ae3d027c6c743e27290fa0500da4777da57154f2acab52775bf"
        }
      }
    ]
  },
  "commands": {
    "prepare": [
      "echo foo \u003e\u003e main.c"
    ],
    "build": [
      "make"
    ],
    "test": [
      "make test"
    ],
    "install": [
      "cd \"$srcdir\"",
      "install -Dm644 \"LICENSE\" -t \"$pkgdir/usr/share/licenses/$pkgname\"",
      "install -Dm644 \"README.md\" -t \"$pkgdir/usr/share/doc/$pkgname\"",
      "install -Dm755 \"bin/$pkgname\" -t \"$pkgdir/usr/bin\""
    ]
  }
}
 ```
 
## Example PKGBUILD output:
 
 ```bash
# Maintainer: John Doe <jd@example.org>
# Generated at: 2020-01-10 00:42:46.792588521 +0200 EET m=+0.000536267 

pkgname=exampleapp
pkgver=v1.0.0
pkgrel=1
pkgdesc="my example application"
url="https://github.com/examplerepo/exampleapp"
license=('Apache 2.0')
arch=('aarch64' 'arm' 'ppc64' 'ppc64le' 'x86_64')
install=$pkgname.install
depends_x86_64=('example-core-x86')

depends=('example-core')

makedepends=('example-dev')

checkdepends=('example-test')
optdepends=('php: because PHP is EPIC!')
sha256sums_aarch64=('11d2b36d6b320dfee489d475635b53206b59288537554ea8bc24f97d06139d64')
sha256sums_arm=('5e79210655a9a71a7b77a3168194e9ead024a120182fa8560348a24dc87da159')
sha256sums_ppc64=('f744e32caf67a609aa435df9f8c519460b1856f7968c057e6ba61397cf79ec15')
sha256sums_ppc64le=('6baef7ee046ceb4450e703a87f05fa5662708d4c3562c26abb427d34b4c82819')
sha256sums_x86_64=('de3edfb94d5d0ae3d027c6c743e27290fa0500da4777da57154f2acab52775bf')
source_aarch64=("https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-arm64.tar.gz")
source_arm=("https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-arm.tar.gz")
source_ppc64=("https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-ppc64.tar.gz")
source_ppc64le=("https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-ppc64le.tar.gz")
source_x86_64=("https://github.com/examplerepo/exampleapp/releases/download/$pkgver/$pkgname-$pkgver-linux-amd64.tar.gz")

prepare() {
  echo foo >> main.c
}

build() {
  make
}

check() {
  make test
}

package() {
  cd "$srcdir"
  install -Dm644 "LICENSE" -t "$pkgdir/usr/share/licenses/$pkgname"
  install -Dm644 "README.md" -t "$pkgdir/usr/share/doc/$pkgname"
  install -Dm755 "bin/$pkgname" -t "$pkgdir/usr/bin"
}
 ```
