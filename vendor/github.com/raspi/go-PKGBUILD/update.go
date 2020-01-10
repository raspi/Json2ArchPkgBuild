package PKGBUILD

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"regexp"
)

type checksumType uint8

const (
	Sha1 checksumType = iota
	Sha224
	Sha256
	Sha384
	Sha512
	B2
	Md5
)

func (ct checksumType) String() string {
	switch ct {
	case Sha1:
		return `sha1`
	case Sha224:
		return `sha224`
	case Sha256:
		return `sha256`
	case Sha384:
		return `sha384`
	case Sha512:
		return `sha512`
	case B2:
		return `b2`
	case Md5:
		return `md5`
	default:
		return `?unknown?`
	}
}

// Update checksums to file(s)
// File must be in format
// <checksum> <file path>
// String ReplaceFromChecksumFilename is replaced with architecture name from checksum filename's architecture
func GetChecksumsFromFile(chtype checksumType, path string, fn func(fpath string) (url string, arch string, alias string)) (f Files) {
	f = make(Files)
	lines, err := GetLinesFromFile(path)

	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		checksumAndFilename := regexp.MustCompile(`^([^\s]+)\s+([^\s]+)$`)
		matches := checksumAndFilename.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		checksum := matches[1]
		fname := matches[2]

		url, arch, alias := fn(fname)

		if url == `` {
			continue
		}

		newSource := Source{
			URL: url,
			Checksums: map[string]string{
				chtype.String(): checksum,
			},
		}

		if alias != `` {
			newSource.Alias = alias
		}

		f[arch] = append(f[arch], newSource)
	}

	return f
}

// Read a file and split with new line separator
func GetLinesFromFile(path string) (lines []string, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(bytes.NewReader(b))

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, nil
}
