package redact_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/diegoholiveira/go-redact"
)

func TestTrieDataStructure(t *testing.T) {
	trie := redact.NewTrie(
		"Capelinha",
		"Minas Gerais",
		"Brasil",
	)

	scenarios := []struct {
		Expression string
		Exists     bool
	}{
		{
			Expression: "Capelinha",
			Exists:     true,
		},
		{
			Expression: "Minas Gerais",
			Exists:     true,
		},
		{
			Expression: "Republica Federativa do Brasil",
			Exists:     false,
		},
	}

	for _, scenario := range scenarios {
		assert.Equal(t, scenario.Exists, trie.Search(scenario.Expression), fmt.Sprintf("Failed expression: %s", scenario.Expression))
	}
}
