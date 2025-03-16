package helper

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func Read(path string) (func() (string, error), error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	return func() (string, error) {
		if !scanner.Scan() {
			_ = file.Close()

			return "", io.EOF
		}

		var line string
		line, err = scanner.Text(), scanner.Err()
		if err != nil {
			_ = file.Close()

			return "", err
		}

		return line, nil
	}, nil
}

func ReadAll(path string) (string, error) {
	reader, err := Read(path)
	if err != nil {
		return "", err
	}

	lines := make([]string, 0)
	for {
		line, err := reader()
		if err != nil {
			break
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n"), nil
}
