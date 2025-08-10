package tree

type Trie struct {
	children map[string]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: make(map[string]*Trie),
		isEnd:    false,
	}
}

func (this *Trie) Insert(word string) {
	root := this
	for _, c := range word {
		if _, ok := root.children[string(c)]; !ok {
			root.children[string(c)] = &Trie{children: make(map[string]*Trie)}
		}
		root = root.children[string(c)]
	}
	root.isEnd = true
}

func (this *Trie) Search(word string) bool {
	root := this.SearchPrefix(word)
	return root != nil && root.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	root := this
	for _, ch := range prefix {
		if root.children[string(ch)] == nil {
			return nil
		}
		root = root.children[string(ch)]
	}
	return root
}
