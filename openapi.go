/*
Package openapi embeds and encapsulates the the current OpenAPITools
openapi-generator-cli JAR file (see JAR constant for current version).
This allows all of its commands and functionality to be used from any Go
program easily with a single import (provided Java has been installed on
the host system). The Run function passes all arguments directly to the
embedded (or cached) jar file itself. Note that the embedded JAR file
does not consume any RAM when any function from this package is called
since embeds are saved in the .rodata section of the ELF binary, for
example.)

For those who prefer to use java directly, the JAR file is extracted at
init time to the java.CacheDir (default: os.UserCacheDir/gojavacache)
and can be used like any other jar file from that location until
explicitly removed. The embedded jar file will not be extracted and
cached if a file of the same name is already there.

*/
package openapi

import (
	"embed"
	"path/filepath"

	"github.com/rwxrob/fs/file"
	"github.com/rwxrob/java"
)

//go:embed embed
var files embed.FS

const JAR = "openapi-generator-cli-6.0.1.jar"

var jarFile = filepath.Join(java.CacheDir, JAR)

func init() {
	if file.Exists(jarFile) {
		return
	}
	java.Extract(files, "embed")
}

// Run passes all arguments directly to the JAR file. This function is
// the most important and is called by most other functions in this
// package (which is primarily just an encapsulation of that JAR file).
func Run(args ...string) error {
	cmd := []string{"-jar", JAR}
	cmd = append(cmd, args...)
	return java.Exec(cmd...)
}

// Validate is a shortcut for Run("validate", "-i", path, "--recommend")
func Validate(path string) error {
	return Run("validate", "-i", path, "--recommend")
}

// GenGin generates a Go Gin server using the go-gin-server generator
// taking a single path or URL as input and creating the directory
// specified (out) for output.
func GenGin(in, out string) error {
	return Run("generate", "-g", "go-gin-server", "-i", in, "-o", out)
}
