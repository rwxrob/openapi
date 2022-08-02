// Copyright 2022 rwxrob/openapi Authors
// SPDX-License-Identifier: Apache-2.0

package openapi

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/compfile"
	"github.com/rwxrob/help"
	"github.com/rwxrob/java"
)

func init() {
	Z.Dynamic["jar"] = func() string { return JAR }
	Z.Dynamic["cachedir"] = func() string { return java.CacheDir }
}

var Cmd = &Z.Cmd{
	Name:      `openapi`,
	Summary:   `encapsulated {{jar}}`,
	Version:   `v0.1.0`,
	Copyright: `Copyright 2021 rwxrob/openapi Authors`,
	License:   `Apache-2.0`,
	Source:    `git@github.com:rwxrob/openapi.git`,
	Issues:    `github.com/rwxrob/openapi/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, runCmd, genGinCmd, validateCmd,
	},

	Description: `
		The **{{.Name}}** command encapsulates and extends the commands
		available from the openapitools/openapi-generator-cli. Most will use
		the {{pre "run"}} command to pass arguments directly to the
		embedded/cached jar file. For convenience, additional command aliases
		have been added and all jar commands have been added as parameters
		so that bash tab completion will work. 

		Embedded JAR 
		
		The {{jar}} file is embedded into this command and package but does
		not occupy RAM. It is extracted and cached into {{cachedir}} at init
		time whenever this command (or any subcommand is executed). 
		
		Dependencies 
		
		This command obviously depends on {{pre "java"}} being installed
		before running it. The openapi-generator-cli requires version 1.8 or
		higher. A Dockerfile is available with the source of this command
		for those who wish to bundle a specific version of Java with this
		command. 
		
		`}

var runCmd = &Z.Cmd{
	Name:    `run`,
	Summary: `run openapi-generator-cli jar commands directly`,

	Params: []string{
		"list", "validate", "author", "batch", "config-help",
		"generate", "help", "meta", "validate", "version",
	},

	Description: `
		The **{{.Name}}** command passes the arguments directly to the encapsulated 
		openapi-generator-cli jar file.`,

	Call: func(_ *Z.Cmd, args ...string) error {
		return Run(args...)
	},
}

var genGinCmd = &Z.Cmd{
	Name:    `gengin`,
	Summary: `generate go-gin-server`,
	Usage:   `SPECFILE OUTDIR`,
	Comp:    compfile.New(),

	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The **{{.Name}}** command generates Go gin server code using the
		go-gin-server generator and cleans up the resulting code to be ready
		for Go 1.18+ development. The {{pre "SPECFILE"}} must be the path to
		an OpenAPI 3.0 specification YAML file. The {{pre "OUTDIR"}} must be
		the path to the directory that will be created to contain the
		generated Go code and other artifacts. Anything there will be
		overwritten.`,

	NumArgs: 2,

	Call: func(_ *Z.Cmd, args ...string) error {
		return GenGin(args[0], args[1])
	},
}

var validateCmd = &Z.Cmd{
	Name:    `validate`,
	Summary: `validate YAML spec file with recommendations`,
	Usage:   `SPECFILE`,
	Comp:    compfile.New(),

	Commands: []*Z.Cmd{help.Cmd},

	Description: `
		The **{{.Name}}** command validates a {{pre "SPECFILE"}} with the
		{{pre "--recommends"}} flag set.`,

	NumArgs: 1,

	Call: func(_ *Z.Cmd, args ...string) error {
		return Validate(args[0])
	},
}
