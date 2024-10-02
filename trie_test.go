package redact_test

import (
	"fmt"
	"testing"

	"github.com/diegoholiveira/go-redact"
	"github.com/stretchr/testify/assert"
)

func TestTrieDataStructure(t *testing.T) {
	trie := redact.NewTrieTree(
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
		assert.Equal(t, scenario.Exists, trie.Verify(scenario.Expression), fmt.Sprintf("Failed expression: %s", scenario.Expression))
	}
}
