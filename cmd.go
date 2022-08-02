// Copyright 2022 openapi (bonzai branch) Authors
// SPDX-License-Identifier: Apache-2.0

package openapi

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:      `openapi`,
	Summary:   `encapsulated commands for OpenAPITools generator`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 openapi (bonzai branch) Authors`,
	License:   `Apache-2.0`,
	Source:    `git@github.com:rwxrob/openapi.git`,
	Issues:    `github.com/rwxrog/openapi/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, runCmd,
	},

	Description: `
		The **{{.Name}}** command encapsulates and extends the commands
		available from the openapitools/openapi-generator-cli project
		(written in Java).  `,
}

var runCmd = &Z.Cmd{
	Name:    `run`,
	Summary: `run openapi-generator-cli command directly`,

	Commands: []*Z.Cmd{},

	Params: []string{"list", "validate"},

	Description: `
		The **{{.Name}}** command passes the arguments directly to the encapsulated 
		openapi-generator-cli jar file.`,

	Call: func(_ *Z.Cmd, args ...string) error {
		return Run(args...)
	},
}
