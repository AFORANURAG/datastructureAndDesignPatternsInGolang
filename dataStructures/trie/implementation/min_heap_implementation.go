package trie

// implement trie
type TrieNode struct {
	childrens [26]*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

// the character of word could from a to z
func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	// iterate over the word
	curr := this.root
	for _, c := range word {
		i := c - 'a'
		// its an array
		if curr.childrens[i] == nil {
			curr.childrens[i] = &TrieNode{}

		}
		curr = curr.childrens[i]
	}
	curr.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, c := range word {
		i := c - 'a'
		if curr.childrens[i] == nil {
			return false
		}
		curr = curr.childrens[i]
	}
	return curr.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, c := range prefix {
		i := c - 'a'
		if curr.childrens[i] == nil {
			return false
		}
		curr = curr.childrens[i]
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
