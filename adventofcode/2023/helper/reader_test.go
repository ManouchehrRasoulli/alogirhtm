package helper

import "testing"

func TestReadAll(t *testing.T) {
	data, err := ReadAll("test.file")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
}
