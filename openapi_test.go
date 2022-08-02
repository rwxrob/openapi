package openapi_test

import (
	"github.com/rwxrob/openapi"
)

func ExampleRun() {
	openapi.Run("version")
	//Output:
	// 6.0.1
}

func ExampleValidate() {
	openapi.Validate("testdata/openapi.yaml")
	// Output:
	// blah
}
