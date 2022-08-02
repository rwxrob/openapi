package openapi_test

import (
	"github.com/rwxrob/openapi"
)

func ExampleRun() {
	openapi.Run("version")
	//Output:
	// 6.0.1
}

/*

func ExampleValidate() {

	var err error
	orig := os.Stdout
	os.Stdout, err = os.Create("testdata/validate.out")
	if err != nil {
		log.Println(err)
	}

	openapi.Validate("testdata/openapi.yaml")

	os.Stdout = orig

	// Output:
	// blah
}


func ExampleGenGin_single_File() {

	var err error
	orig := os.Stdout
	os.Stdout, err = os.Create("testdata/gengin.out")
	if err != nil {
		log.Println(err)
	}
	defer func() { os.Stdout = orig }()

	openapi.GenGin("testdata/openapi.yaml", "testdata/gin")

	// Output:
	// blah
}
*/
