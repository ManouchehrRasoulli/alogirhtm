package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Read(path string) (func() (string, error), error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	return func() (string, error) {
		scanner.Scan()
		line, scannerErr := scanner.Text(), scanner.Err()
		if scannerErr != nil || line == "" {
			if scannerErr == io.EOF {
				_ = file.Close()
			}

			return "", fmt.Errorf("end")
		}

		return line, nil
	}, nil
}
