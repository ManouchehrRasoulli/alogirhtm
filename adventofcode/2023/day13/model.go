package day13

import (
	"bytes"
	"strings"
)

type Value rune

const (
	sand Value = '.'
	rock Value = '#'
)

type Field [][]Value

func (f Field) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("\n")
	for _, row := range f {
		buffer.WriteString(string(row))
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (f Field) StringRow(row int) (string, bool) {
	if row < 0 || row >= len(f) {
		return "", false
	}

	return string(f[row]), true
}

func (f Field) StringColumn(column int) (string, bool) {
	if column < 0 || column >= len(f[0]) {
		return "", false
	}

	var colString bytes.Buffer
	for _, row := range f {
		colString.WriteRune(rune(row[column]))
	}

	return colString.String(), true
}

func ScanFields(input string) ([]Field, error) {
	var (
		fields       []Field
		currentField Field
	)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			fields = append(fields, currentField)
			currentField = Field{}
			continue
		}

		currentField = append(currentField, []Value(line))
	}

	fields = append(fields, currentField)

	return fields, nil
}
