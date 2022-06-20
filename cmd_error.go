package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/wilorios/go-error-handling/error_type"
)

func main() {
	err := testCustomErrorType()
	if err != nil {
		log.Println(err)
	}
	err = testUssingLog()
	if err != nil {
		log.Println(err)
	}

	err = testOpaqueError()
	if err != nil {
		log.Println(err)
	}
	// output
	//2022/04/09 18:19:31 Operation testCustomErrorType: Kind WARN Error testCustomErrorType fail blah blah
	// 2022/04/09 18:19:31 msg test error ussing log method testUssingLog level ERROR
	// 2022/04/09 18:19:31 test error
	// 2022/04/09 18:19:31 adding stacktrace: testOpaquerErrorStacktrace
}

// this has a drawbacks because we have coupling the code
// this shouldnt be part of your public API
// not recomended
func testCustomErrorType() error {
	return &error_type.CustomTypeError{
		Op:   "testCustomErrorType",
		Kind: "ERROR",
		Err:  errors.New("testCustomErrorType fail blah blah"),
	}
}

//Not recommended, if we need to do the stacktrace the same error appears in the log
func testUssingLog() error {
	log.Println(
		"level", "ERROR",
		"msg", "test error ussing log",
		"method", "testUssingLog",
	)
	return errors.New("test ussing log")
}

// you know that there are an error but you cant see inside it
// the caller only knows if the operation worked or it did not
//recommended
func testOpaqueError() error {
	err := testOpaquerErrorStacktrace()
	if err != nil {
		//this is doing a wrapping of the error
		return fmt.Errorf("adding stacktrace: %w", err)
	}
	return nil
}

func testOpaquerErrorStacktrace() error {
	return errors.New("ERROR testOpaquerErrorStacktrace blah blah")
}
