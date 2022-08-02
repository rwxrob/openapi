package openapi

import (
	"embed"

	"github.com/rwxrob/fs/dir"
	"github.com/rwxrob/java"
)

//go:embed embed
var files embed.FS

func init() {
	if dir.Exists(java.CacheDir) {
		//log.Print("cached")
		return
	}
	java.Extract(files, "embed")
}

func Run(args ...string) error {
	cmd := []string{"-jar", "openapi-generator-cli-6.0.1.jar"}
	cmd = append(cmd, args...)
	return java.Exec(cmd...)
}

func Validate(file string) error {
	return Run("validate", "-i", file, "--recommend")
}
