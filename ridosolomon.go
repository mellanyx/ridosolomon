package ridosolomon

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func Hello(name string) {
	fmt.Sprintf("Hello %v, i created by Alexandr Starikov!", name)
}

func fileToArByte(filename string) ([]byte, error) {
	// return []byte, nil if status OK
	// return nil, err if status FAIL
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	return b, nil
}
