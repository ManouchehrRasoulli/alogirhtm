package trie

type Trie struct {
	x       map[byte]*Trie
	anyWord bool
}

func Constructor() Trie {
	t := Trie{
		x:       make(map[byte]*Trie),
		anyWord: false,
	}

	return t
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.anyWord = true
		return
	}

	if t, ok := this.x[word[0]]; ok {
		t.Insert(word[1:])
	} else {
		tr := Constructor()
		this.x[word[0]] = &tr
		this.Insert(word)
	}
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.anyWord
	}

	if t, ok := this.x[word[0]]; ok {
		return t.Search(word[1:])
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	if t, ok := this.x[prefix[0]]; ok {
		return t.StartsWith(prefix[1:])
	}

	return false
}
