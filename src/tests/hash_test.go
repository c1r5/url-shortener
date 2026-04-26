package tests

import (
	"testing"

	"github.com/c1r5/url-shortener/src/internal"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assertion := assert.New(t)
	encoded := internal.GenerateCode("https://chatgpt.com/c/69e42a83-6ca4-83e9-93a8-b09497a17e47")
	t.Logf("Encoded: %s", encoded)
	assertion.Len(encoded, 6)
}
