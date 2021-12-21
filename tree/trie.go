package tree

// 前缀树
// 208
type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	isLeaf   bool
	Val      rune
	children map[rune]*TrieNode
}

func NewTrie() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode, 0),
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, let := range word {
		if _, ok := cur.children[let]; !ok {
			cur.children[let] = &TrieNode{
				Val:      let,
				children: make(map[rune]*TrieNode, 0),
			}
		}
		cur = cur.children[let]
	}
	cur.isLeaf = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, let := range word {
		if _, ok := cur.children[let]; !ok {
			return false
		}
		cur = cur.children[let]
	}

	if !cur.isLeaf {
		return false
	}

	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, let := range prefix {
		if _, ok := cur.children[let]; !ok {
			return false
		}
		cur = cur.children[let]
	}

	return true
}

// 648
// 211
