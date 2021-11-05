package ridosolomon

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
)

func Hello(name string) {
	fmt.Sprintf("Hello %v, i created by Alexandr Starikov!", name)
}

func FileToArByte(filename string) ([]byte, error) {
	// return []byte, nil if status OK
	// return nil, err if status FAIL
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	return b, nil
}

func ArByteToFile(ar []byte, ext string, permissions fs.FileMode) error {
	// return nil if status OK
	// return err if status FAIL
	err := ioutil.WriteFile(fmt.Sprintf("file_from_byte.%s", ext), ar, permissions)
	if err != nil {
		return errors.Unwrap(err)
	}

	return nil
}
