package file

import (
	"io/ioutil"
	"os"
)

func Read(filename string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func Write(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
func Clear(filename string) error {
	return os.Truncate(filename, 0)
}
