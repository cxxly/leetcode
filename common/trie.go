//remove duplications of phone number

package main

type TrieNode struct {
	isNum  bool
	childs [9]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

type Trie struct {
	root TrieNode
}

func NewTrie() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(num string) {
	p := t.root
	for i := 0; i < len(num); i++ {
		index := num[i] - '0'
		if p.childs[index] == nil {
			p.childs[index] = NewTrieNode()
		}
		p = p.childs[index]
	}
	p.isNum = true
}

func (t *Trie) Search(num string) {

}

func (t *Trie) StartWith(prefix string) {

}

func (t *Trie) Walk(root TrieNode, tmp string, res *[]string) {
	if root.isNum == true {
		res := append(res, tmp)
	} else {
		for i := 0; i < len(p.childs); i++ {
			Walk(root.childs[i], append(tmp, string(root.childs)), res)
		}
	}
}
