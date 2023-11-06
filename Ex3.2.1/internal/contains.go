package contains

import (
	"io"
	"os"
	"strings"
)

func Contains(filename string, substr string) (bool, error) {
	file, err := os.Open(filename)

	if err != nil {
		return false, err
	}

	defer file.Close()

	data := make([]byte, 256)

	var n int

	for err == nil {
		n, err = file.Read(data)

		str := string(data[:n])

		if strings.Contains(str, substr) {
			return true, nil
		}
	}

	if err != io.EOF {
		return false, nil
	}

	return false, nil
}
