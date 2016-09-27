// trie_test.go
package main

import (
	"fmt"
	"testing"
)

func TestDuplicationsNum(t *testing.T) {
	nums := []string{"15105201705", "17888817055", "13811016886", "13811016887", "17888817055", "15001250513"}
	trie := NewTrie()
	for num := range nums {
		trie.Insert(num)
	}

}
