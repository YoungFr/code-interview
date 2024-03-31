package template

// 字典树
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		ch := word[i]
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i]
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}
