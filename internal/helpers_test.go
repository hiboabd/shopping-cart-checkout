package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDivisibleByTrue(t *testing.T) {
	result := isDivisibleBy(2, 2)
	assert.Equal(t, true, result)
}

func TestIsDivisibleByFalse(t *testing.T) {
	result := isDivisibleBy(2, 3)
	assert.Equal(t, false, result)
}

func TestPenniesToPounds(t *testing.T) {
	result := penniesToPounds(100)
	assert.Equal(t, 1.00, result)
}
