package redact

type node struct {
	current rune
	childs  map[rune]*node
}

const (
	rootNode        rune = '+'
	endOfExpression rune = '|'
)

// NewTrieTree creates a new node for a Trie Tree data structure.
//
// The Trie data structure [1] is used to efficiently store and
// search for single and composed words like "Company" and "Company X"
//
// [1] https://en.wikipedia.org/wiki/Trie
func NewTrieTree(expressions ...string) *node {
	n := node{
		current: rootNode,
		childs:  make(map[rune]*node),
	}

	for _, expression := range expressions {
		n.Add(expression)
	}

	return &n
}

func newNode(current rune) *node {
	n := node{
		current: current,
		childs:  make(map[rune]*node),
	}

	return &n
}

func (tree *node) Add(expression string) {
	if len(expression) == 0 {
		return
	}

	if tree.current != rootNode {
		return
	}

	node := tree
	for i, c := range expression {
		if _, found := node.childs[c]; !found {
			node.childs[c] = newNode(c)
		}
		node = node.childs[c]
		if len(expression) == i+1 {
			node.childs[endOfExpression] = newNode(endOfExpression)
		}
	}
}

func (tree *node) Verify(expression string) bool {
	node := tree
	found := false
	for _, c := range expression {
		if node, found = node.childs[c]; !found {
			return false
		}
	}
	if _, found := node.childs[endOfExpression]; found {
		return true
	}
	return false
}
