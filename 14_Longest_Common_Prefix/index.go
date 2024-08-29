// package main

// import "fmt"

// type Node struct {
// 	children map[string]*Node
// 	isEnd    bool
// }

// type Trie struct {
// 	root *Node
// }

// func (t *Trie) Insert(w string) {
// 	currentNode := t.root
// 	for i := 0; i < len(w); i++ {
// 		char := string(w[i])
// 		if _, exists := currentNode.children[char]; !exists {
// 			currentNode.children[char] = &Node{
// 				children: make(map[string]*Node),
// 				isEnd:    false,
// 			}
// 		}
// 		currentNode = currentNode.children[char]
// 	}
// 	currentNode.isEnd = true
// }

// func initialize_trie() *Trie {
// 	return &Trie{
// 		root: &Node{
// 			children: make(map[string]*Node),
// 		},
// 	}
// }

// func helper(strs []string) *Trie {
// 	t := initialize_trie()
// 	for _, v := range strs {
// 		t.Insert(v)
// 	}

// 	return t
// }

// func longestCommonPrefix(strs []string) string {
// 	t := helper(strs)
// 	prefix := ""
// 	currentNode := t.root

// 	for len(currentNode.children) == 1 && !currentNode.isEnd {
// 		for k := range currentNode.children {
// 			prefix += k
// 			currentNode = currentNode.children[k]
// 		}
// 	}

// 	return prefix
// }

// func main() {
// 	strs := []string{
// 		"flower",
// 		"flow",
// 		"flight",
// 		"flash",
// 	}

// 	fmt.Println(longestCommonPrefix(strs))
// }
