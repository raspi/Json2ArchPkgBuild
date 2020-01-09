# Json2ArchPkgBuild
Generate Arch Linux PKGBUILD files from JSON.

Uses https://github.com/raspi/go-PKGBUILD as library

## Example

Generate example JSON which you can use to for generating PKGBUILD package:

    ./json2archpkgbuild -example > my-package.json
    $EDITOR my-package.json 

Generate PKGBUILD    

    ./json2archpkgbuild my-package.json > PKGBUILD
