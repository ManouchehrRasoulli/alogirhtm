package trie

import "testing"

func TestString(t *testing.T) {
	t.Log("some-string"[1:])

	trie := Constructor()

	trie.Insert("apple")
	t.Log(trie.Search("apple"))
	t.Log(trie.Search("app"))
	t.Log(trie.StartsWith("app"))
	trie.Insert("app")
	t.Log(trie.Search("app"))

	t.Log(trie)
}
